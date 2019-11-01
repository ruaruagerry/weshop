package goods

import (
	"encoding/json"
	"sort"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
)

type listReq struct {
	CategoryID string `json:"categoryid"`
}

type listRsp struct {
	GoodsList []*rconst.Good `json:"goodslist"`
}

func listHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "goods.listHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &listReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("listHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	// playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("SMEMBERS", rconst.SetGoodsCategoryPrefix+req.CategoryID)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	goodids, _ := redis.Strings(redisMDArray[0], nil)
	log.Info("goodids:", goodids)

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

	// do something
	goodslist := []*rconst.Good{}
	for i := range goodids {
		goodbytes, _ := redis.Bytes(redisMDArray[i], nil)

		tmp := &rconst.Good{}
		err := json.Unmarshal(goodbytes, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s goodbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		goodslist = append(goodslist, tmp)
	}

	sort.Stable(goodid(goodslist))

	// rsp
	rsp := &listRsp{
		GoodsList: goodslist,
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

	log.Info("listHandle rsp, rsp:", string(data))

	return
}
