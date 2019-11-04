package shop

import (
	"encoding/json"
	"sort"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type addressListRsp struct {
	AddressList []*rconst.ShopAddress `json:"addresslist"`
}

func addressListHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.addressListHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGETALL", rconst.HashShopStoreAddressPrefix+playerid)
	// conn.Send("HGET", rconst.HashShopUsedAddress, playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	addressmap, _ := redis.StringMap(redisMDArray[0], nil)
	// usedaddressid, _ := redis.String(redisMDArray[1], nil)

	// do something
	addresslist := []*rconst.ShopAddress{}
	for _, v := range addressmap {
		tmp := &rconst.ShopAddress{}
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("地址unmarshal解析失败")
			log.Errorf("code:%d msg:%s shopaddress unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		addresslist = append(addresslist, tmp)
	}
	sort.Stable(sortaddressid(addresslist))

	// rsp
	rsp := &addressListRsp{
		AddressList: addresslist,
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

	log.Info("addressListHandle rsp, rsp:", string(data))

	return
}
