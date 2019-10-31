package catalog

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
	CategoryList []*rconst.Category `json:"categorylist"`
}

func currentHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "catalog.currentHandle")

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
	conn.Send("HGETALL", rconst.HashCatalogCategoryPrefix+idstr)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	categorybytes, _ := redis.StringMap(redisMDArray[0], nil)

	categorylist := []*rconst.Category{}
	for _, v := range categorybytes {
		tmp := &rconst.Category{}
		err := json.Unmarshal([]byte(v), tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("导航页签unmarshal解析失败")
			log.Errorf("code:%d msg:%s navbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		categorylist = append(categorylist, tmp)
	}

	sort.Stable(categoryid(categorylist))

	// rsp
	rsp := &currentRsp{
		CategoryList: categorylist,
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

	log.Info("currentHandle rsp, rsp:", string(data))

	return
}
