package catagory

import (
	"encoding/json"
	"fmt"
	"sort"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/gogo/protobuf/proto"
)

type currentReq struct {
	ID int `json:"id"`
}

type currentRsp struct {
	CatagoryList []*rconst.Catagory `json:"catagorylist"`
}

func currentHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "catagory.currentHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &currentReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s proto Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("currentHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	idstr := fmt.Sprintf("%d", req.ID)

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGETALL", rconst.HashCatagoryInfoPrefix+idstr)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:统一获取缓存操作失败 redisMDArray Values err, err:%s", httpRsp.GetResult(), err.Error())
		return
	}

	catagorybytes, _ := redis.ByteSlices(redisMDArray[0], nil)

	catagorylist := []*rconst.Catagory{}
	for _, v := range catagorybytes {
		tmp := &rconst.Catagory{}
		err := json.Unmarshal(v, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("导航页签unmarshal解析失败")
			log.Errorf("code:%d msg:%s navbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		catagorylist = append(catagorylist, tmp)
	}

	sort.Stable(catagoryid(catagorylist))

	// rsp
	rsp := &currentRsp{
		CatagoryList: catagorylist,
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

	log.Info("currentHandle rsp, rsp:", string(data))

	return
}
