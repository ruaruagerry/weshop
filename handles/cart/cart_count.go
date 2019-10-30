package cart

import (
	"encoding/json"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type countRsp struct {
	CartTotal int32 `json:"carttotal"`
}

func countHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.countHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("countHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("ZCARD", rconst.ZSetCartInfoPrefix+playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:统一获取缓存操作失败 redisMDArray Values err, err:%s", httpRsp.GetResult(), err.Error())
		return
	}

	carttotal, _ := redis.Int(redisMDArray[0], nil)

	// rsp
	rsp := &countRsp{
		CartTotal: int32(carttotal),
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("返回信息marshal解析失败")
		log.Errorf("code:%d msg:返回信息marshal解析失败 proto marshal err, err:%s", httpRsp.GetResult(), err.Error())
		return
	}
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	httpRsp.Data = data

	log.Info("countHandle rsp, rsp:", string(data))

	return
}
