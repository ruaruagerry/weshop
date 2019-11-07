package shop

import (
	"encoding/json"
	"strings"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type orderDetailReq struct {
	OrderID string `json:"orderid"`
}

type detailaddress struct {
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
}

type detailgood struct {
	GoodID        string `json:"goodid"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Num           int32  `json:"num"`
	Specification string `json:"specification"`
	Price         int64  `json:"price"`
}

type detailorder struct {
	OrderID      string `json:"orderid"`
	Time         int64  `json:"time"`
	Status       string `json:"status"`
	GoodsPrice   int64  `json:"goodsprice"`
	FreightPrice int64  `json:"freightprice"`
	Price        int64  `json:"price"`
}

type orderDetailRsp struct {
	Order      *detailorder   `json:"order"`
	OrderGoods []*detailgood  `json:"ordergoods"`
	Address    *detailaddress `json:"address"`
}

func orderDetailHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.orderDetailHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &orderDetailReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("orderDetailHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashShopOrderPrefix+playerid, req.OrderID)
	conn.Send("HGET", rconst.HashShopCheckoutConfig, rconst.FieldShopFreightPrice)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	orderbyte, _ := redis.Bytes(redisMDArray[0], nil)
	freightprice, _ := redis.Int64(redisMDArray[1], nil)

	// do something
	shoporder := &rconst.ShopOrder{}
	err = json.Unmarshal(orderbyte, shoporder)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("订单unmarshal解析失败")
		log.Errorf("code:%d msg:%s order unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	goodidmap := map[string]bool{}
	goodids := []string{}
	goodnummap := map[string]int32{}
	goodspecificationmap := map[string]string{}
	for _, v := range shoporder.Carts {
		specslice := []string{}
		for _, v1 := range v.Specification {
			specslice = append(specslice, v1.Value)
		}

		goodidmap[v.GoodID] = true
		goodnummap[v.GoodID] += v.Num
		goodspecificationmap[v.GoodID] = strings.Join(specslice, ";")
	}

	for k := range goodidmap {
		goodids = append(goodids, k)
	}

	conn.Send("MULTI")
	for _, v := range goodids {
		conn.Send("HGET", rconst.HashGoodsInfo, v)
	}
	conn.Send("HGET", rconst.HashShopStoreAddressPrefix+playerid, shoporder.AddressID)
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

	addressbyte, _ := redis.Bytes(redisMDArray[len(goodids)], nil)
	shopaddress := &rconst.ShopAddress{}
	err = json.Unmarshal(addressbyte, shopaddress)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("地址unmarshal解析失败")
		log.Errorf("code:%d msg:%s address unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// 拼接返回
	goodsprice := int64(0)
	for k, v := range goodnummap {
		goodsprice += goodmap[k].Price * int64(v)
	}

	rspdetailorder := &detailorder{
		OrderID:      shoporder.OrderID,
		Time:         shoporder.Time,
		Status:       statusToString(shoporder.Status),
		GoodsPrice:   goodsprice,
		FreightPrice: freightprice,
		Price:        goodsprice + freightprice,
	}

	rspdetailgoods := []*detailgood{}
	for _, v := range goodmap {
		tmp := &detailgood{
			GoodID:        v.GoodID,
			URL:           v.URL,
			Name:          v.Name,
			Num:           goodnummap[v.GoodID],
			Specification: goodspecificationmap[v.GoodID],
			Price:         v.Price,
		}

		rspdetailgoods = append(rspdetailgoods, tmp)
	}

	rspaddress := &detailaddress{
		Name:    shopaddress.Name,
		Mobile:  shopaddress.Mobile,
		Address: shopaddress.Address,
	}

	// rsp
	rsp := &orderDetailRsp{
		Order:      rspdetailorder,
		OrderGoods: rspdetailgoods,
		Address:    rspaddress,
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

	log.Info("orderDetailHandle rsp, rsp:", string(data))

	return
}
