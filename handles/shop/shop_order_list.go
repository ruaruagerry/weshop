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

type order struct {
	OrderID  string         `json:"orderid"`
	Status   int32          `json:"status"`
	GoodList []*rconst.Good `json:"goodlist"`
}

type orderListRsp struct {
	OrderList []*order `json:"orderlist"`
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
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	ordermap, _ := redis.StringMap(redisMDArray[0], nil)

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
	cartindexs := []int32{}
	for _, v := range orders {
		for _, v1 := range v.CartIndexs {
			cartindexs = append(cartindexs, v1)
		}
	}

	conn.Send("MULTI")
	for _, v := range cartindexs {
		conn.Send("ZRANGE", rconst.ZSetCartInfoPrefix+playerid, v, v)
	}
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cartbyte, _ := redis.ByteSlices(redisMDArray[0], nil)

	cartmap := map[int32]*rconst.Cart{}
	goodids := []string{}
	for i, v := range cartbyte {
		tmp := &rconst.Cart{}
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("购物车商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s cart unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		cartmap[int32(i)] = tmp
		goodids = append(goodids, tmp.GoodID)
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
	for _, v := range goodids {
		tmp := &rconst.Good{}
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s good unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		goodmap[v] = tmp
	}

	orderlist := []*order{}
	for _, v := range orders {
		tmpgoodlist := []*rconst.Good{}
		for _, v1 := range v.CartIndexs {
			tmpgoodlist = append(tmpgoodlist, goodmap[cartmap[v1].GoodID])
		}

		tmp := &order{
			OrderID:  v.OrderID,
			Status:   v.Status,
			GoodList: tmpgoodlist,
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
