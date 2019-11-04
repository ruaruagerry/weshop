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

type checkoutRsp struct {
	Address      *rconst.ShopAddress `json:"address"`
	FreightPrice int64               `json:"freightprice"`
}

func checkoutHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.checkoutHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("checkoutHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashShopCheckoutConfig, rconst.FieldShopFreightPrice)
	conn.Send("HGET", rconst.HashShopUsedAddress, playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// do something
	freightprice, _ := redis.Int64(redisMDArray[0], nil)
	addressid, _ := redis.String(redisMDArray[1], nil)

	address := &rconst.ShopAddress{}
	if addressid != "" {
		conn.Send("MULTI")
		conn.Send("HGET", rconst.HashShopStoreAddressPrefix+playerid, addressid)
		redisMDArray, err := redis.Values(conn.Do("EXEC"))
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("统一获取缓存操作失败")
			log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		addressbyte, _ := redis.Bytes(redisMDArray[0], nil)
		if len(addressbyte) != 0 {
			err := json.Unmarshal(addressbyte, address)
			if err != nil {
				httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
				httpRsp.Msg = proto.String("地址Unmarshal解析失败")
				log.Errorf("code:%d msg:%s address Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
				return
			}
		}
	}

	// rsp
	rsp := &checkoutRsp{
		Address:      address,
		FreightPrice: freightprice,
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

	log.Info("checkoutHandle rsp, rsp:", string(data))

	return
}
