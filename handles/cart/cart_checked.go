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

type checkedReq struct {
	Index   int32 `json:"index"`
	Checked bool  `json:"checked"`
}

func checkedHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.checkedHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &checkedReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("checkedHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("ZRANGE", rconst.ZSetCartInfoPrefix+playerid, req.Index, req.Index, "withscores")
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cartinfo, err := redis.ByteSlices(redisMDArray[0], nil)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("未找到对应购物商品")
		log.Errorf("code:%d msg:%s not find cart info, index:%d", httpRsp.GetResult(), httpRsp.GetMsg(), req.Index)
		return
	}

	// do something
	cart := &rconst.Cart{}
	err = json.Unmarshal(cartinfo[0], cart)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("购物商品unmarshal解析失败")
		log.Errorf("code:%d msg:%s cartbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	carttime := int64(0)
	err = json.Unmarshal(cartinfo[1], &carttime)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("购物商品时间unmarshal解析失败")
		log.Errorf("code:%d msg:%s carttime unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cart.Checked = req.Checked

	// redis multi set
	conn.Send("MULTI")
	conn.Send("ZREMRANGEBYRANK", rconst.ZSetCartInfoPrefix+playerid, req.Index, req.Index)
	cartbytes, err := json.Marshal(cart)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("购物商品marshal解析失败")
		log.Errorf("code:%d msg:%s cart marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	conn.Send("ZADD", rconst.ZSetCartInfoPrefix+playerid, carttime, cartbytes)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("checkedHandle rsp, result:", httpRsp.GetResult())

	return
}
