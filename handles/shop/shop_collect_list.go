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

type collectListRsp struct {
	CollectList []*rconst.Good `json:"collectlist"`
}

func collectListHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.collectListHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("collectListHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("SMEMBERS", rconst.SetShopCollectPrefix+playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	goodids, _ := redis.Strings(redisMDArray[0], nil)

	// do something
	sort.Strings(goodids)

	// get good
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

	collectlist := []*rconst.Good{}
	for i := range goodids {
		goodbyte, _ := redis.Bytes(redisMDArray[i], nil)

		tmp := &rconst.Good{}
		err := json.Unmarshal(goodbyte, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s good unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		collectlist = append(collectlist, tmp)
	}

	// rsp
	rsp := &collectListRsp{
		CollectList: collectlist,
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

	log.Info("collectListHandle rsp, rsp:", string(data))

	return
}
