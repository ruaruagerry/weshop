package tool

import (
	"encoding/json"
	"fmt"
	"weshop/gamecfg"
	"weshop/gconst"
	"weshop/pb"
	"weshop/rconst"
	"weshop/server"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
)

const (
	secert = "MM4QMybtZ9cG3G9ZALsjmH4EwBmvw2bJ"
)

func writeconfigHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "tool.writeconfigHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	log.Info("writeconfigHandle enter")

	conn := c.RedisConn
	// playerid := c.UserID

	reqsecert := c.Query.Get("secert")
	if reqsecert != secert {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParam))
		httpRsp.Msg = proto.String("Fuckman!!!!secert error")
		log.Errorf("code:%d msg:%s secert err, secert:%s reqsecert:%s", httpRsp.GetResult(), httpRsp.GetMsg(), secert, reqsecert)
		return
	}

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HKEYS", rconst.HashCatalogNav)
	conn.Send("SMEMBERS", rconst.SetUsers)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	navids, _ := redis.Strings(redisMDArray[0], nil)
	allusers, _ := redis.Strings(redisMDArray[1], nil)

	conn.Send("MULTI")
	for _, v := range navids {
		conn.Send("HKEYS", rconst.HashCatalogCategoryPrefix+v)
	}
	redisMDArray, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	/* 清除 */
	conn.Send("MULTI")
	for i, v := range navids {
		categoryids, _ := redis.Strings(redisMDArray[i], nil)
		for _, v1 := range categoryids {
			// 商品所属小类
			conn.Send("DEL", rconst.SetGoodsCategoryPrefix+v1)
		}
		// 小类
		conn.Send("DEL", rconst.HashCatalogCategoryPrefix+v)
	}
	// 大类
	conn.Send("DEL", rconst.HashCatalogNav)
	// 商品
	conn.Send("DEL", rconst.HashGoodsInfo)
	// 商品展示
	conn.Send("DEL", rconst.HashGoodsGallery)
	// 商品问答
	conn.Send("DEL", rconst.HashGoodsIssue)
	// 商品商标
	conn.Send("DEL", rconst.HashGoodsBrand)
	// 商品规格
	conn.Send("DEL", rconst.HashGoodsSpecification)
	// 系统配置
	conn.Send("DEL", rconst.HashShopCheckoutConfig)
	// 玩家数据
	for _, v := range allusers {
		// 订单
		conn.Send("DEL", rconst.HashShopOrderPrefix+v)
		// 收藏
		conn.Send("DEL", rconst.SetShopCollectPrefix+v)
		// 购物车
		conn.Send("DEL", rconst.ZSetCartInfoPrefix+v)
	}
	// 玩家索引
	conn.Send("DEL", rconst.HashCatalogNavIndex)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	/* 写入 */
	conn.Send("MULTI")
	// 大类
	for _, v := range gamecfg.Catalog_navArray {
		tmpnav := &rconst.Nav{
			ID:        int32(v.ID),
			Name:      v.Name,
			BannerURL: v.Bannerurl,
		}

		navbyte, err := json.Marshal(tmpnav)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("大类marshal解析失败")
			log.Errorf("code:%d msg:%s nav marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		conn.Send("HSET", rconst.HashCatalogNav, tmpnav.ID, navbyte)
	}
	// 大类中的小类
	for _, v := range gamecfg.Catalog_categoryArray {
		tmpcatalog := &rconst.Category{
			ID:         int32(v.ID),
			CategoryID: v.Categoryid,
			Name:       v.Name,
			Describe:   v.Describe,
			URL:        v.Url,
		}

		catalogbyte, err := json.Marshal(tmpcatalog)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("小类marshal解析失败")
			log.Errorf("code:%d msg:%s category marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}

		conn.Send("HSET", rconst.HashCatalogCategoryPrefix+fmt.Sprintf("%d", tmpcatalog.ID), tmpcatalog.CategoryID, catalogbyte)
	}
	// 小类中的商品
	for _, v := range gamecfg.GoodsArray {
		tmpgood := &rconst.Good{
			GoodID:   v.Goodid,
			URL:      v.Url,
			Name:     v.Name,
			Describe: v.Describe,
			Price:    int64(v.Price),
		}

		// 商品所属小类
		conn.Send("SADD", rconst.SetGoodsCategoryPrefix+v.Categoryid, v.Goodid)

		// 商品信息
		goodbyte, err := json.Marshal(tmpgood)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商品marshal解析失败")
			log.Errorf("code:%d msg:%s good marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsInfo, v.Goodid, goodbyte)

		// 展览信息
		gallerybyte, err := json.Marshal(v.GalleryArr)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("展览图marshal解析失败")
			log.Errorf("code:%d msg:%s gallery marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsGallery, v.Goodid, gallerybyte)

		// 商品属性
		tmpattribute := []*rconst.GoodAttribute{}
		for i := range v.AttributenameArr {
			tmp := &rconst.GoodAttribute{
				Name:  v.AttributenameArr[i],
				Value: v.AttributevalueArr[i],
			}

			tmpattribute = append(tmpattribute, tmp)
		}
		attributebyte, err := json.Marshal(tmpattribute)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("属性marshal解析失败")
			log.Errorf("code:%d msg:%s attribute marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsAttribute, v.Goodid, attributebyte)

		// 商品问题
		tmpissue := []*rconst.GoodIssue{}
		for i := range v.QuestionArr {
			tmp := &rconst.GoodIssue{
				Question: v.QuestionArr[i],
				Answer:   v.AnswerArr[i],
			}

			tmpissue = append(tmpissue, tmp)
		}
		issuebyte, err := json.Marshal(tmpissue)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("问题marshal解析失败")
			log.Errorf("code:%d msg:%s issue marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsIssue, v.Goodid, issuebyte)

		// 商品商标
		tmpbrand := &rconst.GoodBrand{
			BrandID:   v.Brandid,
			BrandName: v.Brandname,
		}
		brandbyte, err := json.Marshal(tmpbrand)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("商标marshal解析失败")
			log.Errorf("code:%d msg:%s brand marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsBrand, v.Goodid, brandbyte)

		// 商品规格
		tmpspecification := []*rconst.GoodSpecification{}
		for _, v1 := range v.SpecificationidArr {
			specfg := gamecfg.Goods_specificationMap[v1]
			if specfg == nil {
				httpRsp.Result = proto.Int32(int32(gconst.ErrTableConfig))
				httpRsp.Msg = proto.String("未找到规格配置")
				log.Errorf("code:%d msg:%s not find specification cfgm, cfgid:%s", httpRsp.GetResult(), httpRsp.GetMsg(), v1)
				return
			}

			tmp := &rconst.GoodSpecification{
				SpecificationID: v1,
				Name:            specfg.Name,
			}

			tmpdetail := []*rconst.GoodSpecificationDetail{}
			for i2 := range specfg.DetailidArr {
				ttmp := &rconst.GoodSpecificationDetail{
					DetailID: specfg.DetailidArr[i2],
					Value:    specfg.DetailvalueArr[i2],
				}

				tmpdetail = append(tmpdetail, ttmp)
			}

			tmp.ValueList = tmpdetail

			tmpspecification = append(tmpspecification, tmp)
		}
		specificationbyte, err := json.Marshal(tmpspecification)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
			httpRsp.Msg = proto.String("规格marshal解析失败")
			log.Errorf("code:%d msg:%s specification marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
		conn.Send("HSET", rconst.HashGoodsSpecification, v.Goodid, specificationbyte)
	}
	// 其它配置
	conn.Send("HSET", rconst.HashShopCheckoutConfig, rconst.FieldShopFreightPrice, gamecfg.GlobalArray[0].Freightprice)
	conn.Send("HSET", rconst.HashShopCheckoutConfig, rconst.FieldShopWXPayURL, gamecfg.GlobalArray[0].Wxpayurl)
	conn.Send("HSET", rconst.HashShopCheckoutConfig, rconst.FieldShopZFBPayURL, gamecfg.GlobalArray[0].Zfbpayurl)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// // rsp
	// rsp := &pb.HelloRsp{}
	// data, err := json.Marshal(rsp)
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
	// 	httpRsp.Msg = proto.String("返回信息marshal解析失败")
	// 	log.Errorf("code:%d msg:%s json marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
	// 	return
	// }
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	// httpRsp.Data = data

	// log.Info("helloHandle rsp, rsp:", string(data))

	return
}
