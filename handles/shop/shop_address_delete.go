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

type addressDeleteReq struct {
	AddressID string `json:"addressid"`
}

func addressDeleteHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.addressDeleteHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &addressDeleteReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("addressDeleteHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashShopStoreAddressPrefix+playerid, req.AddressID)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	addressbyte, _ := redis.Bytes(redisMDArray[0], nil)

	// do something
	shopaddress := &rconst.ShopAddress{}
	err = json.Unmarshal(addressbyte, shopaddress)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("地址unmarshal解析失败")
		log.Errorf("code:%d msg:%s address unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// redis multi set
	conn.Send("MULTI")
	conn.Send("HDEL", rconst.HashShopStoreAddressPrefix+playerid, shopaddress.AddressID)
	if shopaddress.IsDefault {
		conn.Send("HDEL", rconst.HashShopUsedAddress, playerid)
	}
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("addressDeleteHandle rsp, result:", httpRsp.GetResult())

	return
}
