package cart

import (
	"encoding/json"
	"time"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/garyburd/redigo/redis"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/proto"
)

type addSpecification struct {
	SpecificationID string `json:"specificationId"`
	DetailID        string `json:"valueDetailid"`
	Value           string `json:"valueText"`
}

type addReq struct {
	GoodID        string              `json:"goodid"`
	Num           int32               `json:"num"`
	Specification []*addSpecification `json:"specification"`
}

type addRsp struct {
	CartTotal int32 `json:"carttotal"`
}

// todo 加上订单最大数的限制
func addHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "cart.addHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &addReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("addHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("ZCARD", rconst.ZSetCartInfoPrefix+playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	carttotal, _ := redis.Int(redisMDArray[0], nil)

	// do something
	cartid, err := uuid.NewV4()
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrCreateUUID))
		httpRsp.Msg = proto.String("生成uuid失败")
		log.Errorf("code:%d msg:%s uuid newv4 err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	cart := &rconst.Cart{
		CartID:  cartid.String(),
		GoodID:  req.GoodID,
		Num:     req.Num,
		Checked: false,
	}

	cartspecification := []*rconst.CartSpecification{}
	for _, v := range req.Specification {
		tmp := &rconst.CartSpecification{
			SpecificationID: v.SpecificationID,
			DetailID:        v.DetailID,
			Value:           v.Value,
		}

		cartspecification = append(cartspecification, tmp)
	}
	cart.Specification = cartspecification

	carttotal++

	// redis multi set
	conn.Send("MULTI")
	cartbyte, err := json.Marshal(cart)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("订单货物marshal解析失败")
		log.Errorf("code:%d msg:%s cart Marshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	conn.Send("ZADD", rconst.ZSetCartInfoPrefix+playerid, time.Now().Unix(), cartbyte)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	rsp := &addRsp{
		CartTotal: int32(carttotal),
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

	log.Info("helloHandle rsp, result:", httpRsp.GetResult())

	return
}
