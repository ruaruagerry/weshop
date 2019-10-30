package cart

import (
	"weshop/gconst"
	"weshop/pb"
	"weshop/server"

	"github.com/golang/protobuf/proto"
)

func indexHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.indexHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("#######################")

	// // req
	// req := &pb.HelloReq{}
	// if err := proto.Unmarshal(c.Body, req); err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
	// 	httpRsp.Msg = proto.String("请求信息解析失败")
	// 	log.Errorf("code:%d msg:请求信息解析失败 proto Unmarshal err:%s", httpRsp.GetResult(), err.Error())
	// 	return
	// }

	// log.Info("helloHandle enter, req:", req.String())

	// conn := c.RedisConn
	// playerid := c.UserID

	// // redis multi get
	// conn.Send("MULTI")
	// redisMDArray, err := redis.Values(conn.Do("EXEC"))
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
	// 	httpRsp.Msg = proto.String("统一获取缓存操作失败")
	// 	log.Errorf("code:%d msg:统一获取缓存操作失败 redisMDArray Values err, err:%s", httpRsp.GetResult(), err.Error())
	// 	return
	// }

	// // do something

	// // redis multi set
	// conn.Send("MULTI")
	// _, err = conn.Do("EXEC")
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
	// 	httpRsp.Msg = proto.String("统一存储缓存操作失败")
	// 	log.Errorf("code:%d msg:统一存储缓存操作失败 exec err, err:%s", httpRsp.GetResult(), err.Error())
	// 	return
	// }

	// // rsp
	// rsp := &pb.HelloRsp{}
	// data, err := proto.Marshal(rsp)
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
	// 	httpRsp.Msg = proto.String("返回信息marshal解析失败")
	// 	log.Errorf("code:%d msg:返回信息marshal解析失败 proto marshal err, err:%s", httpRsp.GetResult(), err.Error())
	// 	return
	// }
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	// httpRsp.Data = data

	// log.Info("helloHandle rsp, rsp:", rsp.String())

	return
}
