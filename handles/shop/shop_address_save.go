package shop

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

type addressSaveReq struct {
	AddressID string `json:"addressid"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isdefault"`
}

// todo 加上地址最大数的限制
func addressSaveHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "shop.addressSaveHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("req:", string(c.Body))

	// req
	req := &addressSaveReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	log.Info("addressSaveHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashShopUsedAddress, playerid)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	usedaddressid, _ := redis.String(redisMDArray[0], nil)

	// do something
	addressid := req.AddressID
	if req.AddressID == "" {
		uuidaddressid, err := uuid.NewV4()
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrCreateUUID))
			httpRsp.Msg = proto.String("生成uuid失败")
			log.Errorf("code:%d msg:%s uuid newv4 err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		addressid = uuidaddressid.String()
	}

	// 修改默认地址属性
	usedaddress := &rconst.ShopAddress{}
	saveusedaddress := false
	if req.IsDefault && usedaddressid != "" && req.AddressID != usedaddressid {
		// redis multi get
		conn.Send("MULTI")
		conn.Send("HGET", rconst.HashShopStoreAddressPrefix+playerid, usedaddressid)
		// conn.Send("HGET", rconst.HashShopUsedAddress, playerid)
		redisMDArray, err = redis.Values(conn.Do("EXEC"))
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("统一获取已使用地址缓存操作失败")
			log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		usedaddressbyte, _ := redis.Bytes(redisMDArray[0], nil)
		if len(usedaddressbyte) != 0 {
			err := json.Unmarshal(usedaddressbyte, usedaddress)
			if err != nil {
				httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
				httpRsp.Msg = proto.String("默认地址marshal解析失败")
				log.Errorf("code:%d msg:%s default address json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
				return
			}

			usedaddress.IsDefault = false
			saveusedaddress = true
		}
	}

	address := &rconst.ShopAddress{
		AddressID: addressid,
		Name:      req.Name,
		Mobile:    req.Mobile,
		Address:   req.Address,
		IsDefault: req.IsDefault,
	}
	if req.AddressID == "" {
		address.Time = time.Now().Unix()
	}

	// redis multi set
	conn.Send("MULTI")
	addressbyte, err := json.Marshal(address)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("地址marshal解析失败")
		log.Errorf("code:%d msg:%s address json marshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	conn.Send("HSET", rconst.HashShopStoreAddressPrefix+playerid, address.AddressID, addressbyte)
	if saveusedaddress {
		usedaddressbyte, err := json.Marshal(usedaddress)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("默认地址marshal解析失败")
			log.Errorf("code:%d msg:%s default address json Marshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashShopStoreAddressPrefix+playerid, usedaddress.AddressID, usedaddressbyte)
	}
	if req.IsDefault {
		conn.Send("HSET", rconst.HashShopUsedAddress, playerid, address.AddressID)
	}
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("addressSaveHandle rsp, result:", httpRsp.GetResult())

	return
}
