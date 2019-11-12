package gamecfg

// Catalog_category TODO: game config struct
type Catalog_category struct {
    CfgID string `csv:"cfgId"` // ID
    ID int `csv:"id"` // 大类分组ID
    Categoryid string `csv:"categoryid"` // 小类分组ID
    Describe string `csv:"describe"` // 描述
    Name string `csv:"name"` // 名称
    Url string `csv:"url"` // 图片
}

// Catalog_nav TODO: game config struct
type Catalog_nav struct {
    CfgID string `csv:"cfgId"` // ID
    ID int `csv:"id"` // 大类分组ID
    Name string `csv:"name"` // 名称
    Bannerurl string `csv:"bannerurl"` // 浮框图片
}

// Global TODO: game config struct
type Global struct {
    CfgID string `csv:"cfgId"` // ID
    Freightprice int `csv:"freightprice"` // 运费
    Wxpayurl string `csv:"wxpayurl"` // 微信支付码
    Zfbpayurl string `csv:"zfbpayurl"` // 支付宝支付码
}

// Goods TODO: game config struct
type Goods struct {
    CfgID string `csv:"cfgId"` // ID
    Categoryid string `csv:"categoryid"` // 小类分组ID
    Goodid string `csv:"goodid"` // 商品ID
    Brandid string `csv:"brandid"` // 品牌ID
    Brandname string `csv:"brandname"` // 品牌名称
    Price int `csv:"price"` // 价格（单位元）
    Name string `csv:"name"` // 商品名称
    Describe string `csv:"describe"` // 商品描述
    Url string `csv:"url"` // 商品图片
    GalleryArr []string `csv:"gallery"` // 展示栏1
    AttributenameArr []string `csv:"attributename"` // 属性1
    AttributevalueArr []string `csv:"attributevalue"` // 属性值1
    QuestionArr []string `csv:"question"` // 问题1
    AnswerArr []string `csv:"answer"` // 答案1
    SpecificationidArr []string `csv:"specificationid"` // 规格1id
}

// Goods_specification TODO: game config struct
type Goods_specification struct {
    CfgID string `csv:"cfgId"` // 规格ID
    Name string `csv:"name"` // 规格名称
    DetailidArr []string `csv:"detailid"` // 规格详情id1
    DetailvalueArr []string `csv:"detailvalue"` // 规格详情值1
}

// Test TODO: game config struct
type Test struct {
    CfgID string `csv:"cfgId"` // 题目ID
    GroudID string `csv:"groudID"` // 组
    Type int `csv:"type"` // 答题类型(1单选题；2判断题；3多选题)
    NextCfgID string `csv:"nextCfgId"` // 下个问题ID
    Name string `csv:"name"` // 问题
    Text string `csv:"text"` // 答案解析
    AnswerArr []string `csv:"answer"` // 选择答案1
    TrueIndexArr []int `csv:"trueIndex"` // 正确答案1
}

