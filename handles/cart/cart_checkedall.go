package cart

import (
	"encoding/json"
	"strconv"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type checkedallReq struct {
	CheckedAll bool `json:"checkedall"`
}

func checkedallHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.checkedallHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &checkedallReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("checkedallHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("ZRANGE", rconst.ZSetCartInfoPrefix+playerid, 0, -1, "withscores")
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cartinfo, _ := redis.Strings(redisMDArray[0], nil)

	// do something
	cartslice := []*rconst.Cart{}
	carttimeslice := []int64{}
	for i, v := range cartinfo {
		if i%2 == 0 {
			cart := &rconst.Cart{}
			err = json.Unmarshal([]byte(v), cart)
			if err != nil {
				httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
				httpRsp.Msg = proto.String("购物商品unmarshal解析失败")
				log.Errorf("code:%d msg:%s cartbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
				return
			}

			cart.Checked = req.CheckedAll

			cartslice = append(cartslice, cart)
		} else {
			tmptime, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
				httpRsp.Msg = proto.String("时间ParseInt解析失败")
				log.Errorf("code:%d msg:%s time strconv err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
				return
			}

			carttimeslice = append(carttimeslice, tmptime)
		}
	}

	// redis multi set
	conn.Send("MULTI")
	conn.Send("DEL", rconst.ZSetCartInfoPrefix+playerid)
	for i, v := range cartslice {
		tmpbyte, err := json.Marshal(v)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("购物商品marshal解析失败")
			log.Errorf("code:%d msg:%s cart marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		conn.Send("ZADD", rconst.ZSetCartInfoPrefix, carttimeslice[i], tmpbyte)
	}
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("checkedallHandle rsp, result:", httpRsp.GetResult())

	return
}
