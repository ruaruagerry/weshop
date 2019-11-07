package shop

import (
	"encoding/json"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type listordergood struct {
	GoodID string `json:"goodid"`
	URL    string `json:"url"`
	Name   string `json:"name"`
	Num    int32  `json:"num"`
}

type listorder struct {
	OrderID  string           `json:"orderid"`
	Status   int32            `json:"status"`
	GoodList []*listordergood `json:"goodlist"`
	Price    int64            `json:"price"`
}

type orderListRsp struct {
	OrderList []*listorder `json:"orderlist"`
}

func orderListHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.orderListHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("orderListHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGETALL", rconst.HashShopOrderPrefix+playerid)
	conn.Send("HGET", rconst.HashShopCheckoutConfig, rconst.FieldShopFreightPrice)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	ordermap, _ := redis.StringMap(redisMDArray[0], nil)
	freightprice, _ := redis.Int64(redisMDArray[1], nil)

	// do something
	orders := []*rconst.ShopOrder{}
	for _, v := range ordermap {
		tmp := &rconst.ShopOrder{}
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("订单unmarshal解析失败")
			log.Errorf("code:%d msg:%s order unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		orders = append(orders, tmp)
	}

	// find cart
	goodidmap := map[string]bool{}
	goodids := []string{}
	for _, v := range orders {
		for _, v1 := range v.Carts {
			goodidmap[v1.GoodID] = true
		}
	}

	for k := range goodidmap {
		goodids = append(goodids, k)
	}

	conn.Send("MULTI")
	for _, v := range goodids {
		conn.Send("HGET", rconst.HashGoodsInfo, v)
	}
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	goodmap := map[string]*rconst.Good{}
	for i, v := range goodids {
		goodbyte, _ := redis.Bytes(redisMDArray[i], nil)

		tmp := &rconst.Good{}
		err := json.Unmarshal(goodbyte, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s good unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		goodmap[v] = tmp
	}

	orderlist := []*listorder{}
	for _, v := range orders {
		tmpgoodlist := []*listordergood{}
		price := freightprice
		for _, v1 := range v.Carts {
			tmpordergood := &listordergood{
				GoodID: goodmap[v1.GoodID].GoodID,
				URL:    goodmap[v1.GoodID].URL,
				Name:   goodmap[v1.GoodID].Name,
				Num:    v1.Num,
			}
			tmpgoodlist = append(tmpgoodlist, tmpordergood)

			price += goodmap[v1.GoodID].Price * int64(tmpordergood.Num)
		}

		tmp := &listorder{
			OrderID:  v.OrderID,
			Status:   v.Status,
			GoodList: tmpgoodlist,
			Price:    price,
		}

		orderlist = append(orderlist, tmp)
	}

	// rsp
	rsp := &orderListRsp{
		OrderList: orderlist,
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("返回信息marshal解析失败")
		log.Errorf("code:%d msg:%s json marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	httpRsp.Data = data

	log.Info("orderListHandle rsp, rsp:", string(data))

	return
}
