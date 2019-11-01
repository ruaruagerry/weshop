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

type detailReq struct {
	GoodID string `json:"goodid"`
}

type detailRsp struct {
	Gallery       []string                    `json:"gallery"`
	Attribute     []*rconst.GoodAttribute     `json:"attribute"`
	Issue         []*rconst.GoodIssue         `json:"issue"`
	Brand         *rconst.GoodBrand           `json:"brand"`
	Specification []*rconst.GoodSpecification `json:"specification"`
}

func detailHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "goods.detailHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &detailReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	log.Info("detailHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	// playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashGoodsGallery, req.GoodID)
	conn.Send("HGET", rconst.HashGoodsAttribute, req.GoodID)
	conn.Send("HGET", rconst.HashGoodsIssue, req.GoodID)
	conn.Send("HGET", rconst.HashGoodsBrand, req.GoodID)
	conn.Send("HGET", rconst.HashGoodsSpecification, req.GoodID)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	gallerybyte, _ := redis.Bytes(redisMDArray[0], nil)
	attributebyte, _ := redis.Bytes(redisMDArray[1], nil)
	issuebyte, _ := redis.Bytes(redisMDArray[2], nil)
	brandbyte, _ := redis.Bytes(redisMDArray[3], nil)
	specificationbyte, _ := redis.Bytes(redisMDArray[4], nil)

	// // do something
	gallery := []string{}
	if len(gallerybyte) != 0 {
		err = json.Unmarshal(gallerybyte, &gallery)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("展览图unmarshal解析失败")
			log.Errorf("code:%d msg:%s gallerybyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	attribute := []*rconst.GoodAttribute{}
	if len(attribute) != 0 {
		err = json.Unmarshal(attributebyte, &attribute)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("商品参数unmarshal解析失败")
			log.Errorf("code:%d msg:%s attributebyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	issue := []*rconst.GoodIssue{}
	if len(issuebyte) != 0 {
		err = json.Unmarshal(issuebyte, &issue)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("商品常见问题unmarshal解析失败")
			log.Errorf("code:%d msg:%s issuebyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	brand := &rconst.GoodBrand{}
	if len(brandbyte) != 0 {
		err = json.Unmarshal(brandbyte, &brand)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("商品牌子unmarshal解析失败")
			log.Errorf("code:%d msg:%s brandbyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	specification := []*rconst.GoodSpecification{}
	if len(specificationbyte) != 0 {
		err = json.Unmarshal(specificationbyte, &specification)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("商品规格unmarshal解析失败")
			log.Errorf("code:%d msg:%s specificationbyte Unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	// rsp
	rsp := &detailRsp{
		Gallery:       gallery,
		Attribute:     attribute,
		Issue:         issue,
		Brand:         brand,
		Specification: specification,
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

	log.Info("detailHandle rsp, rsp:", string(data))

	return
}
