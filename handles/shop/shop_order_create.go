package shop

import (
	"encoding/json"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/proto"
)

type orderCreateReq struct {
	AddressID string `json:"addressid"`
}

type orderCreateRsp struct {
	OrderID   string `json:"orderid"`
	WXPayURL  string `json:"wxpayurl"`
	ZFBPayURL string `json:"zfbpayurl"`
}

func orderCreateHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.orderCreateHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &orderCreateReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	log.Info("orderCreateHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HEXISTS", rconst.HashShopStoreAddressPrefix+playerid, req.AddressID)
	conn.Send("ZRANGE", rconst.ZSetCartInfoPrefix+playerid, 0, -1)
	conn.Send("HGET", rconst.HashShopCheckoutConfig, rconst.FieldShopWXPayURL)
	conn.Send("HGET", rconst.HashShopCheckoutConfig, rconst.FieldShopZFBPayURL)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	existaddress, _ := redis.Bool(redisMDArray[0], nil)
	cartbytes, _ := redis.ByteSlices(redisMDArray[1], nil)
	wxpayurl, _ := redis.String(redisMDArray[2], nil)
	zfbpayurl, _ := redis.String(redisMDArray[3], nil)

	// do something
	if !existaddress {
		httpRsp.Result = proto.Int32(int32(gconst.ErrShopNotFindAddress))
		httpRsp.Msg = proto.String("未找到收货地址")
		log.Errorf("code:%d msg:%s not find shop address, addressid:%s", httpRsp.GetResult(), httpRsp.GetMsg(), req.AddressID)
		return
	}

	cartindexs := []int32{}
	for i, v := range cartbytes {
		tmp := &rconst.Cart{}
		err := json.Unmarshal(v, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("购物车商品Unmarshal解析失败")
			log.Errorf("code:%d msg:%s cart unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		if tmp.Checked {
			cartindexs = append(cartindexs, int32(i))
		}
	}

	uuidorderid, err := uuid.NewV4()
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrCreateUUID))
		httpRsp.Msg = proto.String("生成uuid失败")
		log.Errorf("code:%d msg:%s uuid newv4 err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	shoporder := &rconst.ShopOrder{
		OrderID:    uuidorderid.String(),
		AddressID:  req.AddressID,
		CartIndexs: cartindexs,
		Status:     rconst.TypeOrderStatusNotCharged,
	}

	// redis multi set
	conn.Send("MULTI")
	shoporderbyte, err := json.Marshal(shoporder)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("订单marshal解析失败")
		log.Errorf("code:%d msg:%s order json marshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	conn.Send("HSET", rconst.HashShopOrderPrefix+playerid, shoporder.OrderID, shoporderbyte)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	rsp := &orderCreateRsp{
		OrderID:   shoporder.OrderID,
		WXPayURL:  wxpayurl,
		ZFBPayURL: zfbpayurl,
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

	log.Info("orderCreateHandle rsp, rsp:", string(data))

	return
}
