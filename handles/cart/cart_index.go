package cart

import (
	"encoding/json"
	"strings"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type cartgood struct {
	CartID        string `json:"cartid"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Num           int32  `json:"num"`
	Price         int64  `json:"price"`
	Specification string `json:"specification"`
}

type indexRsp struct {
	Cartlist  []*cartgood `json:"cartlist"`
	CartTotal int32       `json:"carttotal"`
}

func indexHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.indexHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("indexHandle enter")

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("ZRANGE", rconst.ZSetCartInfoPrefix+playerid, 0, -1)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cartsbyteslice, _ := redis.ByteSlices(redisMDArray[0], nil)

	// do something
	carts := []*rconst.Cart{}
	for _, v := range cartsbyteslice {
		tmp := &rconst.Cart{}
		err := json.Unmarshal(v, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("购买商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s cartbyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		carts = append(carts, tmp)
	}

	conn.Send("MULTI")
	for _, v := range carts {
		conn.Send("HGET", rconst.HashGoodsInfo, v.GoodID)
	}
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cartlist := []*cartgood{}
	for i, v := range carts {
		goodbytes, _ := redis.Bytes(redisMDArray[i], nil)

		tmp := &rconst.Good{}
		err := json.Unmarshal(goodbytes, tmp)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品unmarshal解析失败")
			log.Errorf("code:%d msg:%s goodbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		specslice := []string{}
		for _, v1 := range v.Specification {
			specslice = append(specslice, v1.Value)
		}

		tmprsp := &cartgood{
			CartID:        v.CartID,
			URL:           tmp.URL,
			Name:          tmp.Name,
			Num:           v.Num,
			Price:         tmp.Price,
			Specification: strings.Join(specslice, ";"),
		}

		cartlist = append(cartlist, tmprsp)
	}

	// rsp
	rsp := &indexRsp{
		Cartlist:  cartlist,
		CartTotal: int32(len(cartlist)),
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

	log.Info("indexHandle rsp, rsp:", string(data))

	return
}
