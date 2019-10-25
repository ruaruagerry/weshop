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
	"github.com/golang/protobuf/proto"
)

type indexRsp struct {
	NavList      []*rconst.Nav      `json:"navlist"`
	CatagoryList []*rconst.Catagory `json:"catagorylist"`
}

func indexHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "catagory.indexHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("indexHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGETALL", rconst.HashCatalogNav)
	conn.Send("HGET", rconst.HashCatalogNavIndex, playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// 导航页签
	navbytes, _ := redis.ByteSlices(redisMDArray[0], nil)
	// 玩家当前导航索引
	navindex, _ := redis.Int(redisMDArray[1], nil)

	// do something
	navlist := []*rconst.Nav{}
	for _, v := range navbytes {
		tmp := &rconst.Nav{}
		err := json.Unmarshal(v, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("导航页签unmarshal解析失败")
			log.Errorf("code:%d msg:%s navbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		navlist = append(navlist, tmp)
	}

	sort.Stable(navid(navlist))

	// 获取当前导航类型
	navid := fmt.Sprintf("%d", navlist[navindex].ID)

	// 获取当前导航物品
	conn.Send("MULTI")
	conn.Send("HGETALL", rconst.HashCatalogInfoPrefix+navid)
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
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
	rsp := &indexRsp{
		NavList:      navlist,
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

	log.Info("indexHandle rsp, rsp:", string(data))

	return
}
