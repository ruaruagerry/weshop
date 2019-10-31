package goods

import (
	"encoding/json"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
)

type countRsp struct {
	GoodsCount int32 `json:"goodscount"`
}

func countHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "goods.countHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("countHandle enter")

	conn := c.RedisConn

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HKEYS", rconst.HashCatalogNav)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	navids, _ := redis.Strings(redisMDArray[0], nil)

	// do something
	conn.Send("MULTI")
	for _, v := range navids {
		conn.Send("HLEN", rconst.HashCatalogCategoryPrefix+v)
	}
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	allcount := 0
	for i := range navids {
		curcount, _ := redis.Int(redisMDArray[i], nil)
		allcount += curcount
	}

	// rsp
	rsp := &countRsp{
		GoodsCount: int32(allcount),
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("返回信息marshal解析失败")
		log.Errorf("code:%d msg:%s proto marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	httpRsp.Data = data

	log.Info("countHandle rsp, rsp:", string(data))

	return
}
