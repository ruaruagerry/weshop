package gamecfg

// AppConfig TODO: game config struct
type AppConfig struct {
    CfgID string `csv:"cfgId"` // 活动ID
    Type int `csv:"type"` // 开关
    StartTime string `csv:"startTime"` // 活动开始时间
    EndTime string `csv:"endTime"` // 活动结束时间
    Describe string `csv:"describe"` // 备注
}

// ApplePayProductIDCfg TODO: game config struct
type ApplePayProductIDCfg struct {
    CfgID string `csv:"cfgId"` // 商品ID
    ProductID string `csv:"productId"` // Product ID
    Description string `csv:"description"` // 描述
    Price int `csv:"price"` // 价格
    AppID int `csv:"appId"` // Apple ID
}

// CompensationAwardCfg TODO: game config struct
type CompensationAwardCfg struct {
    CfgID string `csv:"cfgId"` // ID
    PropName string `csv:"propName"` // 道具名称
    PropType int `csv:"propType"` // 礼包类型
}

// CountAwardParamsOfContinueRechargeCfg TODO: game config struct
type CountAwardParamsOfContinueRechargeCfg struct {
    ArrCountAwardParams1 string `csv:"arr_CountAwardParams_1"` // 奖励1类型
    ArrCountAwardParams2 string `csv:"arr_CountAwardParams_2"` // 奖励1ID
    ArrCountAwardParams3 UNumber `csv:"arr_CountAwardParams_3"` // 奖励1数量
}
// ContinueRechargeCfg TODO: game config struct
type ContinueRechargeCfg struct {
    CfgID string `csv:"cfgId"` // ID 
    ConRechargeType int `csv:"conRechargeType"` // 连充类型
    ConRechargeDays int `csv:"conRechargeDays"` // 连充天数
    CountAwardParamsArr []*CountAwardParamsOfContinueRechargeCfg `csv:"CountAwardParams"` // 奖励1类型
}

// CountAwardParamsOfCumulateConsumeCfg TODO: game config struct
type CountAwardParamsOfCumulateConsumeCfg struct {
    ArrCountAwardParams1 string `csv:"arr_CountAwardParams_1"` // 奖励1类型
    ArrCountAwardParams2 string `csv:"arr_CountAwardParams_2"` // 奖励1ID
    ArrCountAwardParams3 UNumber `csv:"arr_CountAwardParams_3"` // 奖励1数量
}
// CumulateConsumeCfg TODO: game config struct
type CumulateConsumeCfg struct {
    CfgID string `csv:"cfgId"` // ID 
    CumConsumeValue int `csv:"cumConsumeValue"` // 累积消耗
    CountAwardParamsArr []*CountAwardParamsOfCumulateConsumeCfg `csv:"CountAwardParams"` // 奖励1类型
}

// CountAwardParamsOfCumulateRechargeCfg TODO: game config struct
type CountAwardParamsOfCumulateRechargeCfg struct {
    ArrCountAwardParams1 string `csv:"arr_CountAwardParams_1"` // 奖励1类型
    ArrCountAwardParams2 string `csv:"arr_CountAwardParams_2"` // 奖励1ID
    ArrCountAwardParams3 UNumber `csv:"arr_CountAwardParams_3"` // 奖励1数量
}
// CumulateRechargeCfg TODO: game config struct
type CumulateRechargeCfg struct {
    CfgID string `csv:"cfgId"` // ID 
    CumRechargeValue int `csv:"cumRechargeValue"` // 累积充值金额
    CountAwardParamsArr []*CountAwardParamsOfCumulateRechargeCfg `csv:"CountAwardParams"` // 奖励1类型
}

// CountAwardParamsOfCumulateRechargeForeverCfg TODO: game config struct
type CountAwardParamsOfCumulateRechargeForeverCfg struct {
    ArrCountAwardParams1 string `csv:"arr_CountAwardParams_1"` // 奖励1类型
    ArrCountAwardParams2 string `csv:"arr_CountAwardParams_2"` // 奖励1ID
    ArrCountAwardParams3 UNumber `csv:"arr_CountAwardParams_3"` // 奖励1数量
}
// CumulateRechargeForeverCfg TODO: game config struct
type CumulateRechargeForeverCfg struct {
    CfgID string `csv:"cfgId"` // ID 
    CumRechargeValue int `csv:"cumRechargeValue"` // 累积充值金额
    Describe string `csv:"describe"` // 充值文本
    IsRebornNoResetQuests bool `csv:"isRebornNoResetQuests"` // 是否激活重生任务不清0
    CountAwardParamsArr []*CountAwardParamsOfCumulateRechargeForeverCfg `csv:"CountAwardParams"` // 奖励1类型
}

// RandAwardParamOfLotteryRankRewardCfg TODO: game config struct
type RandAwardParamOfLotteryRankRewardCfg struct {
    ArrrandAwardParam1 string `csv:"arr_randAwardParam_1"` // 随机奖励1ID类型
    ArrrandAwardParam2 string `csv:"arr_randAwardParam_2"` // 随机奖励1ID
    ArrrandAwardParam3 int `csv:"arr_randAwardParam_3"` // 随机奖励1数量
    ArrrandAwardParam4 int `csv:"arr_randAwardParam_4"` // 随机奖励1单位
    ArrrandAwardParam5 int `csv:"arr_randAwardParam_5"` // 随机奖励1概率
}
// LotteryRankRewardCfg TODO: game config struct
type LotteryRankRewardCfg struct {
    CfgID string `csv:"cfgId"` // 排名ID
    MailID string `csv:"mailId"` // 奖励邮件ID
    AwardType string `csv:"awardType"` // 必选奖励1ID类型
    AwardID string `csv:"awardId"` // 必选奖励1ID
    AwardCount int `csv:"awardCount"` // 必选奖励1数量
    AwardUnit int `csv:"awardUnit"` // 必选奖励1单位
    RandAwardParamArr []*RandAwardParamOfLotteryRankRewardCfg `csv:"randAwardParam"` // 随机奖励1ID类型
}

// Achieve TODO: game config struct
type Achieve struct {
    CfgID string `csv:"cfgId"` // 功能ID
    GroupID string `csv:"groupID"` // 组ID
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    Icon string `csv:"icon"` // 成就图片
    TypeID string `csv:"typeID"` // 成就类型ID
    Param float32 `csv:"param"` // 参数1
    AwardID string `csv:"awardID"` // 奖励id
}

// AchieveBox TODO: game config struct
type AchieveBox struct {
    CfgID string `csv:"cfgId"` // ID
    PreviousAchieveCnt int `csv:"previousAchieveCnt"` // 上个奖励需要完成成就个数
    AchieveCnt int `csv:"achieveCnt"` // 需要完成成就个数
    Icon1 string `csv:"icon1"` // 奖励icon
    Icon2 string `csv:"icon2"` // 领奖后图片
    AwardID string `csv:"awardID"` // 奖励id
}

// AchieveGameLevel TODO: game config struct
type AchieveGameLevel struct {
    CfgID string `csv:"cfgId"` // 功能ID
    Name string `csv:"name"` // 功能名
    Description string `csv:"description"` // 功能描述
    Icon string `csv:"icon"` // 未解锁icon
    StartGameLevel int `csv:"startGameLevel"` // 开始进度条关卡
    GameLevel int `csv:"gameLevel"` // 解锁关卡
    AwardID string `csv:"awardID"` // 奖励ID
    IsShowWindows bool `csv:"isShowWindows"` // 是否有弹窗(默认true)
}

// AchieveVIP TODO: game config struct
type AchieveVIP struct {
    CfgID string `csv:"cfgId"` // 功能ID
    GroupID string `csv:"groupID"` // 组ID
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    TypeID string `csv:"typeID"` // 成就类型ID
    Param int `csv:"param"` // 参数1
    AwardID string `csv:"awardID"` // 奖励id
}

// ActivityLevelAward TODO: game config struct
type ActivityLevelAward struct {
    CfgID string `csv:"cfgId"` // ID
    Level int `csv:"level"` // 登录
    AwardID string `csv:"awardID"` // 奖励ID
}

// Adbox TODO: game config struct
type Adbox struct {
    CfgID string `csv:"cfgId"` // ID
    Spine string `csv:"spine"` // 模型
    Icon string `csv:"icon"` // 图标
    Probability float32 `csv:"probability"` // 随机概率
    HasAd bool `csv:"hasAd"` // 是否有广告，默认FALSE
    Duration int `csv:"duration"` // 相同宝箱刷新间隔时间秒，默认是0，表示没有时间限制
    DailyLimit int `csv:"dailyLimit"` // 该宝箱产出资源每日累积上限（每种宝箱只能产出一种类型的资源），默认是0，表示没有限制
    AwardID string `csv:"awardId"` // 资源ID
}

// AnnounceCondition TODO: game config struct
type AnnounceCondition struct {
    CfgID string `csv:"cfgId"` // ID
    Name string `csv:"name"` // 名字
    AnncType string `csv:"anncType"` // 类型 hero equipment
    Quality int `csv:"quality"` // 品质 高于品质才广播
    AwardWay string `csv:"awardWay"` // 途径 1.shopBox 商城箱子 2.gameAward 关卡奖励 3.instance  副本奖励 4.firstRecharge 首充
    Description string `csv:"description"` // 描述
}

// AwardOfAppDungeonActiviesAchieve TODO: game config struct
type AwardOfAppDungeonActiviesAchieve struct {
    Arraward1 string `csv:"arr_award_1"` // 奖励1类型
    Arraward2 string `csv:"arr_award_2"` // 奖励1ID
    Arraward3 UNumber `csv:"arr_award_3"` // 奖励1数量
}
// AppDungeonActiviesAchieve TODO: game config struct
type AppDungeonActiviesAchieve struct {
    CfgID string `csv:"cfgId"` // 排名id
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    TypeID string `csv:"typeID"` // 成就类型ID
    Param int `csv:"param"` // 参数1
    AwardArr []*AwardOfAppDungeonActiviesAchieve `csv:"award"` // 奖励1类型
}

// RankAwardOfAppDungeonActiviesRankAward TODO: game config struct
type RankAwardOfAppDungeonActiviesRankAward struct {
    ArrrankAward1 string `csv:"arr_rankAward_1"` // 奖励1类型
    ArrrankAward2 string `csv:"arr_rankAward_2"` // 奖励1ID
    ArrrankAward3 UNumber `csv:"arr_rankAward_3"` // 奖励1数量
}
// AppDungeonActiviesRankAward TODO: game config struct
type AppDungeonActiviesRankAward struct {
    CfgID string `csv:"cfgId"` // 排名id
    MailRewardID string `csv:"mailRewardID"` // 邮件ID
    OffRank int `csv:"offRank"` // 排名下限
    OnRank int `csv:"onRank"` // 排名上限
    RankAwardArr []*RankAwardOfAppDungeonActiviesRankAward `csv:"rankAward"` // 奖励1类型
}

// AwardOfAppEquipmentSuitActiviesAchieve TODO: game config struct
type AwardOfAppEquipmentSuitActiviesAchieve struct {
    Arraward1 string `csv:"arr_award_1"` // 奖励1类型
    Arraward2 string `csv:"arr_award_2"` // 奖励1ID
    Arraward3 UNumber `csv:"arr_award_3"` // 奖励1数量
}
// AppEquipmentSuitActiviesAchieve TODO: game config struct
type AppEquipmentSuitActiviesAchieve struct {
    CfgID string `csv:"cfgId"` // 排名id
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    TypeID string `csv:"typeID"` // 成就类型ID
    Param int `csv:"param"` // 参数1
    AwardArr []*AwardOfAppEquipmentSuitActiviesAchieve `csv:"award"` // 奖励1类型
}

// RankAwardOfAppEquipmentSuitActiviesRankAward TODO: game config struct
type RankAwardOfAppEquipmentSuitActiviesRankAward struct {
    ArrrankAward1 string `csv:"arr_rankAward_1"` // 奖励1类型
    ArrrankAward2 string `csv:"arr_rankAward_2"` // 奖励1ID
    ArrrankAward3 UNumber `csv:"arr_rankAward_3"` // 奖励1数量
}
// AppEquipmentSuitActiviesRankAward TODO: game config struct
type AppEquipmentSuitActiviesRankAward struct {
    CfgID string `csv:"cfgId"` // 排名id
    MailRewardID string `csv:"mailRewardID"` // 邮件ID
    OffRank int `csv:"offRank"` // 排名下限
    OnRank int `csv:"onRank"` // 排名上限
    RankAwardArr []*RankAwardOfAppEquipmentSuitActiviesRankAward `csv:"rankAward"` // 奖励1类型
}

// AwardOfAppNewServiceActiviesAchieve TODO: game config struct
type AwardOfAppNewServiceActiviesAchieve struct {
    Arraward1 string `csv:"arr_award_1"` // 奖励1类型
    Arraward2 string `csv:"arr_award_2"` // 奖励1ID
    Arraward3 UNumber `csv:"arr_award_3"` // 奖励1数量
}
// AppNewServiceActiviesAchieve TODO: game config struct
type AppNewServiceActiviesAchieve struct {
    CfgID string `csv:"cfgId"` // 排名id
    DayNum int `csv:"dayNum"` // 天数
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    TypeID string `csv:"typeID"` // 成就类型ID
    Param float32 `csv:"param"` // 参数1
    AwardArr []*AwardOfAppNewServiceActiviesAchieve `csv:"award"` // 奖励1类型
}

// AwardOfAppNewServiceActiviesAchieveShop TODO: game config struct
type AwardOfAppNewServiceActiviesAchieveShop struct {
    Arraward1 string `csv:"arr_award_1"` // 奖励1类型
    Arraward2 string `csv:"arr_award_2"` // 奖励1ID
    Arraward3 UNumber `csv:"arr_award_3"` // 奖励1数量
}
// AppNewServiceActiviesAchieveShop TODO: game config struct
type AppNewServiceActiviesAchieveShop struct {
    CfgID string `csv:"cfgId"` // 排名id
    Name string `csv:"name"` // 商品名
    DayNum int `csv:"dayNum"` // 天数
    ChildID string `csv:"childID"` // 子id
    Price float32 `csv:"price"` // 价格
    Discount float32 `csv:"discount"` // 打折
    BuyNum int `csv:"buyNum"` // 购买次数
    AwardArr []*AwardOfAppNewServiceActiviesAchieveShop `csv:"award"` // 奖励1类型
}

// AwardOfAppNewServiceActiviesDailyObj TODO: game config struct
type AwardOfAppNewServiceActiviesDailyObj struct {
    Arraward1 string `csv:"arr_award_1"` // 奖励1类型
    Arraward2 string `csv:"arr_award_2"` // 奖励1ID
    Arraward3 UNumber `csv:"arr_award_3"` // 奖励1数量
}
// AppNewServiceActiviesDailyObj TODO: game config struct
type AppNewServiceActiviesDailyObj struct {
    CfgID string `csv:"cfgId"` // 排名id
    DayNum int `csv:"dayNum"` // 天数
    ChildID string `csv:"childID"` // 子id
    Name string `csv:"name"` // 成就名称
    TypeID string `csv:"typeID"` // 成就类型ID
    Param float32 `csv:"param"` // 参数1
    AwardArr []*AwardOfAppNewServiceActiviesDailyObj `csv:"award"` // 奖励1类型
}

// RankAwardOfAppNewServiceActiviesRankAward TODO: game config struct
type RankAwardOfAppNewServiceActiviesRankAward struct {
    ArrrankAward1 string `csv:"arr_rankAward_1"` // 奖励1类型
    ArrrankAward2 string `csv:"arr_rankAward_2"` // 奖励1ID
    ArrrankAward3 UNumber `csv:"arr_rankAward_3"` // 奖励1数量
}
// AppNewServiceActiviesRankAward TODO: game config struct
type AppNewServiceActiviesRankAward struct {
    CfgID string `csv:"cfgId"` // 排名id
    MailRewardID string `csv:"mailRewardID"` // 邮件ID
    OffRank int `csv:"offRank"` // 排名下限
    OnRank int `csv:"onRank"` // 排名上限
    DayNum int `csv:"dayNum"` // 天数
    TypeID string `csv:"typeID"` // 成就类型ID
    RankAwardArr []*RankAwardOfAppNewServiceActiviesRankAward `csv:"rankAward"` // 奖励1类型
}

// AreanRobotLineupOfAreanRobot TODO: game config struct
type AreanRobotLineupOfAreanRobot struct {
    ArrareanRobotLineup1 string `csv:"arr_areanRobotLineup_1"` // 英雄1配置ID
    ArrareanRobotLineup2 int `csv:"arr_areanRobotLineup_2"` // 英雄1等级
    ArrareanRobotLineup3 int `csv:"arr_areanRobotLineup_3"` // 英雄1强化等级
    ArrareanRobotLineup4 int `csv:"arr_areanRobotLineup_4"` // 英雄1超越等级
    ArrareanRobotLineup5 int `csv:"arr_areanRobotLineup_5"` // 英雄1槽
}
// AreanRobot TODO: game config struct
type AreanRobot struct {
    CfgID string `csv:"cfgId"` // ID
    RobotType int `csv:"robotType"` // 类型
    RobotName string `csv:"robotName"` // 昵称
    RobotSex int `csv:"robotSex"` // 性别
    IconSrc string `csv:"iconSrc"` // 头像
    MaxGameLevel int `csv:"maxGameLevel"` // 最大关卡数
    AreanRobotLineupArr []*AreanRobotLineupOfAreanRobot `csv:"areanRobotLineup"` // 英雄1配置ID
    S string `csv:"s"` // 
}

// AwardsParamsOfAwards TODO: game config struct
type AwardsParamsOfAwards struct {
    ArrawardsParams1 string `csv:"arr_awardsParams_1"` // 奖励1表 creatures, equipments, items, 
    ArrawardsParams2 string `csv:"arr_awardsParams_2"` // 奖励1id(对应人物配置表ID，装备配置ID或者物品配置ID)
    ArrawardsParams3 UNumber `csv:"arr_awardsParams_3"` // 奖励1数量下限
    ArrawardsParams4 UNumber `csv:"arr_awardsParams_4"` // 奖励1数量上限
    ArrawardsParams5 int `csv:"arr_awardsParams_5"` // 奖励1几率
}
// Awards TODO: game config struct
type Awards struct {
    CfgID string `csv:"cfgId"` // 奖励id
    Name string `csv:"name"` // 奖励名称
    Description string `csv:"description"` // 描述
    IconSrc string `csv:"iconSrc"` // 图标所在图集： creatures, equipments, items, award
    Icon string `csv:"icon"` // 图标
    AwardType string `csv:"awardType"` // 奖励几率类型(不填默认是certainty)： certainty，必掉1组。后续的n项奖励，只奖励1项，多项奖励概率和为100% rand，全独立随机。后续的n项奖励，每项奖励概率独立计算。    即可能掉0~n项
    IsMul bool `csv:"isMul"` // 奖励数量是否乘以任务每秒产出： TRUE 乘以 FALSE 不乘以
    AwardsParamsArr []*AwardsParamsOfAwards `csv:"awardsParams"` // 奖励1表 creatures, equipments, items, 
}

// BattleConfig TODO: game config struct
type BattleConfig struct {
    CfgID string `csv:"cfgId"` // ID(1普通,2副本,3竞技场)
    CreaturesLevel bool `csv:"creaturesLevel"` // 英雄升级
    CreaturesTalents bool `csv:"creaturesTalents"` // 英雄天赋
    CreaturesSpecialAbility bool `csv:"creaturesSpecialAbility"` // 英雄特殊能力
    CreaturesDnhance bool `csv:"creaturesDnhance"` // 英雄强化
    CreaturesEvo bool `csv:"creaturesEvo"` // 英雄超越
    EquipmentBuff bool `csv:"equipmentBuff"` // 装备buff
    EquipmentSuit bool `csv:"equipmentSuit"` // 装备套装
    Privilege bool `csv:"privilege"` // 特权（包含商店类buff,好友助战、VIP特权等)
    PlayerSpell bool `csv:"playerSpell"` // 主动技能
    BattleTime int `csv:"battleTime"` // 战斗时长 单位为秒。0为无限时间
    GameLevelAccelerate bool `csv:"gameLevelAccelerate"` // 关卡等级有几率获取5倍加速
}

// BattleNum TODO: game config struct
type BattleNum struct {
    CfgID string `csv:"cfgId"` // 挑战次数
    DungeonBattleNum UNumber `csv:"dungeonBattleNum"` // 购买副本挑战次数消耗
    ArenaBattleNum UNumber `csv:"arenaBattleNum"` // 购买竞技场挑战次数消耗
}

// Born TODO: game config struct
type Born struct {
    CfgID string `csv:"cfgId"` // 表id
    MaterialType string `csv:"materialType"` // 类型items,creatures,equipments
    MaterialID string `csv:"materialId"` // 材料id
    Num UNumber `csv:"num"` // 数量
}

// EffectsParamsOfBuffs TODO: game config struct
type EffectsParamsOfBuffs struct {
    ArreffectsParams1 string `csv:"arr_effectsParams_1"` // 效果1
    ArreffectsParams2 UNumber `csv:"arr_effectsParams_2"` // 效果1参数
}
// Buffs TODO: game config struct
type Buffs struct {
    CfgID string `csv:"cfgId"` // 配置ID
    Name string `csv:"name"` // buff名称
    Description string `csv:"description"` // buff描述
    ArtEffect string `csv:"artEffect"` // 美术特效名字
    ArtEffectPosition string `csv:"artEffectPosition"` // 美术特效挂载位置 1.bloodStripNode  血条。   2.headNode        头部。 3.bodyNode        身体原点。 4.footNode        脚下。 
    ControllerType string `csv:"controllerType"` // 控制器类型，默认是duration类型，候选如下： duration， counter
    BuffType string `csv:"buffType"` // buff目标求解器类型，默认owner类型，候选如下（以后更多）： owner，coma, frozen，其中，coma表示昏迷类buff，frozen表示冰冻类buff，owner表示普通buff作用于owner自身。
    Duration int `csv:"duration"` // 如果是duration类buff，这里设置buff持续的时间长度（单位毫秒）， 默认是0,0表示无限长
    Counter int `csv:"counter"` // 如果是counter类buff，这里设置计数次数
    Period int `csv:"period"` // 如果是周期buff，这里设置周期，单位是毫秒，默认是0，0表示无限长周期
    ExlusiveBuffIDs string `csv:"exlusiveBuffIds"` // 互斥buff id字符串，多个id用逗号分隔
    ReplacedBuffIDs string `csv:"replacedBuffIds"` // 替换的buff id字符串，多个id用逗号分隔
    DurationMergeBuffIDs string `csv:"durationMergeBuffIds"` // 时间合并的buff id字符串，多个id用逗号分隔
    DurationMergeBuffLimit int `csv:"durationMergeBuffLimit"` // 允许时间合并最大个数，默认是0，表示不限制
    ConcurrentBuffIDs string `csv:"concurrentBuffIds"` // 允许共存最大个数的buff id列表，多个id用逗号分隔
    ConcurrentBuffLimit int `csv:"concurrentBuffLimit"` // 允许共存最大个数
    EffectsParamsArr []*EffectsParamsOfBuffs `csv:"effectsParams"` // 效果1
    Beizhu1 string `csv:"beizhu1"` // 备注
}

// Creatures TODO: game config struct
type Creatures struct {
    CfgID string `csv:"cfgId"` // ID
    Name string `csv:"name"` // 名称
    CreaturesType string `csv:"creaturesType"` // 生物类型(默认hero) hero, monster, boss
    GroupType string `csv:"groupType"` // 生物组ID
    Icon string `csv:"icon"` // 头像框
    Description string `csv:"description"` // 描述
    ModleScale float32 `csv:"modleScale"` // 英雄详情模型缩放大小
    Spine string `csv:"spine"` // 骨骼资源名字
    Voice1 string `csv:"voice1"` // 语音1
    Voice2 string `csv:"voice2"` // 语音2
    Star int `csv:"star"` // 星级,默认1级
    Quality int `csv:"quality"` // 品质,默认1
    Level int `csv:"level"` // 初始等级,默认1级
    EvoLevel int `csv:"evoLevel"` // 超越等级,默认0级
    EnhanceLevel int `csv:"enhanceLevel"` // 进化等级,默认0级
    Sex string `csv:"sex"` // 性别，默认男性，候选： male, female
    Race string `csv:"race"` // 种族，默认hum，候选： hum, orcs, rebel
    Troop string `csv:"troop"` // 兵种类型,默认air，候选： air, land
    AttackType string `csv:"attackType"` // 伤害类型，默认physic，候选： physic， magic
    AttackValue UNumber `csv:"attackValue"` // 攻击
    Hp UNumber `csv:"hp"` // 生命
    PhysicDefend UNumber `csv:"physicDefend"` // 物防
    MagicDefend UNumber `csv:"magicDefend"` // 法防
    AttackSpeed float32 `csv:"attackSpeed"` // 攻速:每秒攻击多少次
    MoveDuration float32 `csv:"moveDuration"` // 移速：多少秒到达目标位置
    MoveSpeed float32 `csv:"moveSpeed"` // 移速：每秒走多少
    HitChance float32 `csv:"hitChance"` // 命中率 （大于等于0，且小于等于1）
    DodgeChance float32 `csv:"dodgeChance"` // 闪避 （大于等于0，且小于等于1）
    CriticalStrikeChance float32 `csv:"criticalStrikeChance"` // 暴率 （大于等于0，且小于等于1）
    CriticalStrikeDamage float32 `csv:"criticalStrikeDamage"` // 暴伤
    Anger UNumber `csv:"anger"` // 怒气
    InitialAnger UNumber `csv:"initialAnger"` // 初始怒气
    NormalSpellID string `csv:"normalSpellId"` // 普攻技能ID
    BigSpellID string `csv:"bigSpellId"` // 大招技能ID
    SpecialAbilityType1 string `csv:"specialAbilityType1"` // 特殊能力1类型id
    SpecialAbilityID1 string `csv:"specialAbilityId1"` // 特殊能力1的id
    SpecialAbilityType2 string `csv:"specialAbilityType2"` // 特殊能力2类型id
    SpecialAbilityID2 string `csv:"specialAbilityId2"` // 特殊能力2的id
    TalentsArr []string `csv:"talents"` // 天赋1
    RespawnTime int `csv:"respawnTime"` // 复活时间
    IsComaImmune bool `csv:"isComaImmune"` // 昏迷免疫,默认false，候选： true, false
    IsFrozenImmune bool `csv:"isFrozenImmune"` // 冰冻免疫,默认false，候选： true, false
    ComaResistant int `csv:"comaResistant"` // 昏迷抗性
    FrozenResistant int `csv:"frozenResistant"` // 冰冻抗性
    IsBoss bool `csv:"isBoss"` // 是否是BOSS,默认false，候选： true, false
    IsPhysicImmune bool `csv:"isPhysicImmune"` // 附带物免罩子,默认false，候选： true, false
    IsMagicImmune bool `csv:"isMagicImmune"` // 附带魔免罩子,默认false，候选： true, false
    EvoHeroID string `csv:"evoHeroId"` // 超越后的英雄ID,已经最高级无法超越时候不填(怪物也不填)
    EvoFrontHeroID string `csv:"evoFrontHeroId"` // 超越前的英雄ID,无前面id时候不填
    EvoConfigID string `csv:"evoConfigId"` // 超越属性和消耗，关联到heroEvolution（怪物不填)
}

// Dungeon TODO: game config struct
type Dungeon struct {
    CfgID string `csv:"cfgId"` // 副本D
    MaterialType string `csv:"materialType"` // 副本名
    UnlockGameLevel int `csv:"unlockGameLevel"` // 副本解锁所需关卡
    MaterialID string `csv:"materialId"` // 副本图
    IsLeft string `csv:"isLeft"` // 是否左边
    IsHum bool `csv:"isHum"` // 是否限制人类false不限制，true限制
    IsOrcs bool `csv:"isOrcs"` // 是否限制兽族false不限制，true限制
    IsEebel bool `csv:"isEebel"` // 是否限制异族false不限制，true限制
    IsMonday bool `csv:"isMonday"` // 副本是否星期一开放0不开放，1开放
    IsTuesday bool `csv:"isTuesday"` // 副本是否星期二开放
    IsWednesday bool `csv:"isWednesday"` // 副本是否星期三开放
    IsThursday bool `csv:"isThursday"` // 副本是否星期四开放
    IsFriday bool `csv:"isFriday"` // 副本是否星期五开放
    IsSaturday bool `csv:"isSaturday"` // 副本是否星期六开放
    IsSunday bool `csv:"isSunday"` // 副本是否星期日开放
}

// DungeonLevel TODO: game config struct
type DungeonLevel struct {
    CfgID string `csv:"cfgId"` // 副本层数id
    DungeonLayer int `csv:"dungeonLayer"` // 副本标识
    DungeonID string `csv:"dungeonID"` // 副本id
    DungeonName string `csv:"dungeonName"` // 副本名称
    MapName string `csv:"mapName"` // 地图名字
    MonsterIDsArr []string `csv:"monsterIds"` // 怪物id1: 不填表示该位置没有怪物
    EnhanceLevel int `csv:"enhanceLevel"` // 怪物强化等级
    BossIndex int `csv:"bossIndex"` // BOSS索引： 以1起始，例如1表示“怪物id1”为boss，不填表示没有boss
    BossEnhanceLevel int `csv:"bossEnhanceLevel"` // BOSS强化等级
    ShowReward string `csv:"showReward"` // 显示奖励
    CertainlyReward string `csv:"certainlyReward"` // 必掉奖池
    RandomReward string `csv:"randomReward"` // 随机奖池
    Difficulty int `csv:"difficulty"` // 难度星级
    Describe1 string `csv:"describe1"` // 地下城难易度
    Describe2 string `csv:"describe2"` // 地下城小怪特点
    Describe3 string `csv:"describe3"` // 地下城BOSS特点
}

// Effects TODO: game config struct
type Effects struct {
    CfgID string `csv:"cfgId"` // 配置ID
    Name string `csv:"name"` // 效果名字
    Description string `csv:"description"` // 效果描述（其实完全没有用）
    Etype string `csv:"etype"` // 效果类型，不填默认是"property"，候选如下（将来更多）：    property, byDamageProperty, byAttackProperty
    PropertyName string `csv:"propertyName"` // 作用的目标属性名字，默认是"hp"，候选如下（将来更多）： hp,血量 hpMax, 最大血量 mana,魔法 manaMax,最大魔法 attack,攻击力 attackSpeed,攻击速度 dodge,闪避率 hitChance,命中率 physicDefend,物防 magicDefend,魔防 shield,护盾 spellKill，技能中断 moveSpeed,移动速度 criticalStrikeDamage,暴伤 criticalStrikeChance,暴击命中率 respawnTime,复活时间
    ValueType string `csv:"valueType"` // 值的类型，默认是百分比，候选如下：    addMul,addFix,direct
}

// BuffStringsOfEquipmentSuitProperties TODO: game config struct
type BuffStringsOfEquipmentSuitProperties struct {
    ArrbuffStrings1 string `csv:"arr_buffStrings_1"` // 套装属性buff1
    ArrbuffStrings2 string `csv:"arr_buffStrings_2"` // 套装属性buff1作用目标
}
// EquipmentSuitProperties TODO: game config struct
type EquipmentSuitProperties struct {
    CfgID string `csv:"cfgId"` // ID
    Description string `csv:"description"` // 描述
    SuitScore int `csv:"suitScore"` // 装积分
    EvoMulStringsArr []float32 `csv:"evoMulStrings"` // 1星套装倍数
    BuffStringsArr []*BuffStringsOfEquipmentSuitProperties `csv:"buffStrings"` // 套装属性buff1
    Beizhu1 string `csv:"beizhu1"` // 备注1
    Beizhu2 string `csv:"beizhu2"` // 备注2
}

// EquipmentSuits TODO: game config struct
type EquipmentSuits struct {
    CfgID string `csv:"cfgId"` // 套装id
    Name string `csv:"name"` // 套装名
    EquipmentIDsArr []string `csv:"equipmentIds"` // 穿戴的装备ID1
    SuitPropertiesArr []string `csv:"suitProperties"` // 套装属性1需要装备数（激活条件）
}

// Equipments TODO: game config struct
type Equipments struct {
    CfgID string `csv:"cfgId"` // ID
    Name string `csv:"name"` // 名称
    Icon string `csv:"icon"` // 图标
    Star int `csv:"star"` // 星级，默认1级
    Quality int `csv:"quality"` // 品质，默认是1
    InitLevel int `csv:"initLevel"` // 初始等级，默认1级
    BuffID1 string `csv:"buffId1"` // 装备buff1ID
    BuffTargetType1 string `csv:"buffTargetType1"` // 装备buff1ID作用目标，默认是self
    BuffEffectValue1 float32 `csv:"buffEffectValue1"` // buff1提升倍数
    BuffEnhance1 float32 `csv:"buffEnhance1"` // 强化提升的buff1倍数
    BuffID2 string `csv:"buffId2"` // 装备buff2ID
    BuffTargetType2 string `csv:"buffTargetType2"` // 装备buff2ID作用目标，默认是self
    BuffEffectValue2 float32 `csv:"buffEffectValue2"` // buff2提升倍数
    BuffEnhance2 float32 `csv:"buffEnhance2"` // 强化提升的buff2倍数
    SuitID string `csv:"suitId"` // 套装id
    EnhanceCostResouceType1 string `csv:"enhanceCostResouceType1"` // 强化需要消耗材料1资源表ID
    EnhanceCostResouceID1 string `csv:"enhanceCostResouceId1"` // 强化需要消耗材料1ID
    EnhanceCostResouceNum1 int `csv:"enhanceCostResouceNum1"` // 每级强化需要消耗材料1增加数量
    EnhanceCostResouceType2 string `csv:"enhanceCostResouceType2"` // 强化需要消耗材料2资源表ID
    EnhanceCostResouceID2 string `csv:"enhanceCostResouceId2"` // 强化需要消耗材料2ID
    EnhanceCostResouceNum2 int `csv:"enhanceCostResouceNum2"` // 每级强化需要消耗材料2增加数量
    DeComposeGainResouceType1 string `csv:"deComposeGainResouceType1"` // 装备分解获得资源表ID1
    DeComposeGainResouceID1 string `csv:"deComposeGainResouceId1"` // 装备分解获得资源ID1
    DeComposeGainResouceNum1 int `csv:"deComposeGainResouceNum1"` // 装备分解获得资源数量1
    DeComposeGainResouceType2 string `csv:"deComposeGainResouceType2"` // 装备分解获得资源表ID2
    DeComposeGainResouceID2 string `csv:"deComposeGainResouceId2"` // 装备分解获得资源ID2
    DeComposeGainResouceNum2 int `csv:"deComposeGainResouceNum2"` // 装备分解获得资源数量2
    Description string `csv:"description"` // 无尽的装备名备注
    DeComposeNeedResouceID string `csv:"deComposeNeedResouceId"` // 分解需要消耗资源ID
    DeComposeNeedResouceNum int `csv:"deComposeNeedResouceNum"` // 分解需要消耗资源数量
    Maxlevel int `csv:"maxlevel"` // 最大强化等级
    EvolutionIDParamsArr []string `csv:"evolutionIdParams"` // 进化等级1配置ID
    Beizhu1 string `csv:"beizhu1"` // 备注1
    Beizhu2 string `csv:"beizhu2"` // 备注2
    Beizhu3 string `csv:"beizhu3"` // 备注3
}

// EquipmentsEvolution TODO: game config struct
type EquipmentsEvolution struct {
    CfgID string `csv:"cfgId"` // 装备进化ID 
    CostResourceParamsArr []string `csv:"costResourceParams"` // 进化消耗资源ID1
    DeComposeGainResouceParamsArr []string `csv:"deComposeGainResouceParams"` // 分解返还资源ID1
    Beizhu1 string `csv:"beizhu1"` // 备注1
}

// EventPoint TODO: game config struct
type EventPoint struct {
    CfgID string `csv:"cfgId"` // ID
    IDentification string `csv:"identification"` // 事件标识
    Description string `csv:"description"` // 事件描述
}

// FirstRecharge TODO: game config struct
type FirstRecharge struct {
    CfgID string `csv:"cfgId"` // 充值id
    RechargeID string `csv:"rechargeID"` // 默认充值金额(元)
    AwardID string `csv:"awardID"` // 奖励id
    HeroID string `csv:"heroID"` // 英雄id
    HeroDes1 string `csv:"heroDes1"` // 描述1
    HeroDes2 string `csv:"heroDes2"` // 描述2
    HeroDes3 string `csv:"heroDes3"` // 描述3
}

// GameLevels TODO: game config struct
type GameLevels struct {
    CfgID string `csv:"cfgId"` // 关卡ID,格式约定如下：关卡编号-波次编号，例如1-1，编号必须大于等于1
    MapName string `csv:"mapName"` // 地图名字
    MonsterIDsArr []string `csv:"monsterIds"` // 怪物id1: 不填表示该位置没有怪物
    BossIndex int `csv:"bossIndex"` // BOSS索引： 以1起始，例如1表示“怪物id1”为boss，不填表示没有boss
    MonsterLevelOverride int `csv:"monsterLevelOverride"` // 怪物等级： 不填表示使用生物配置表中的等级
    BossWarning string `csv:"bossWarning"` // 是否显示boss报警:true，false
}

// Global TODO: game config struct
type Global struct {
    CfgID string `csv:"cfgId"` // 配置ID:本表格只有一行配置，因此该id随便填，但不能留空
    CreatureLevelRaise float32 `csv:"creatureLevelRaise"` // 等级对攻防血的提升倍数: 不填默认是1倍，也即是没有任何提升
    HeroEnhanceLevelRaise float32 `csv:"heroEnhanceLevelRaise"` // 英雄强化等级对攻防血的提升倍数： 不填默认是1倍，也即是没有任何提升
    CreatureLevelIncByGameLevel int `csv:"creatureLevelIncByGameLevel"` // 生物等级的增加量，随着关卡等级的变化： 默认为0
    EquipmentsLineupMax int `csv:"equipmentsLineupMax"` // 装备最大上阵数
    HeroLineupMax int `csv:"heroLineupMax"` // 英雄最大上阵数
    DeComposeProportion float32 `csv:"deComposeProportion"` // 分解装备强化材料返还比例
    GameLevleMaxWave int `csv:"gameLevleMaxWave"` // 一关最大波数
    DropGoldCoefficient float32 `csv:"dropGoldCoefficient"` // 【客户端】金币掉落系数 掉落金币=最大任务每秒产出金币*(金币掉落系数(百分比)+金币关卡系数(百分比)*(关数*10+波数))
    GameLevelDropGoldCoefficient float32 `csv:"gameLevelDropGoldCoefficient"` // 【客户端】金币关卡系数
    ServerDropGoldCoefficient float32 `csv:"serverDropGoldCoefficient"` // 【服务器】金币掉落系数 掉落金币=最大任务每秒产出金币*(金币掉落系数(百分比)+金币关卡系数(百分比)*(关数*10+波数))
    ServerGameLevelDropGoldCoefficient float32 `csv:"serverGameLevelDropGoldCoefficient"` // 【服务器】金币关卡系数
    ServerDropGoldNumMin float32 `csv:"serverDropGoldNumMin"` // 【服务器】金币至少奖励数量
    AdboxMinLevelCompletedNumber int `csv:"adboxMinLevelCompletedNumber"` // 完成波数后生成adbox
    AdboxMaxLevelCompletedNumber int `csv:"adboxMaxLevelCompletedNumber"` // 完成最大波数生成abox
    AdboxBattleTimeInSeconds int `csv:"adboxBattleTimeInSeconds"` // 持续玩多长时间后生成adbox，单位是秒 
    QuestsMaxLevel int `csv:"questsMaxLevel"` // 任务最大等级
    HeroMaxLevel int `csv:"heroMaxLevel"` // 英雄初始最高级
    HeroMaxEnhanceLevel int `csv:"heroMaxEnhanceLevel"` // 英雄强化最高级
    GameLevelMinCirculate int `csv:"gameLevelMinCirculate"` // 【关卡】循环关起始关
    GameLevelMaxCirculate int `csv:"gameLevelMaxCirculate"` // 【关卡】循环关结束关
    OfflineGoldCoe float32 `csv:"offlineGoldCoe"` // 【离线收益】 离线收益=任务每秒获得金币*挂机时间*N(0<N<1）
    OfflineGameLevelCoe float32 `csv:"offlineGameLevelCoe"` // 【离线收益】 离线闯关最多能达波数=N*历史最大关卡波数(0<0<1)
    OfflineMaxDay int `csv:"offlineMaxDay"` // 【离线收益】离线挂机最多挂x天
    RebornNeedGameLevel int `csv:"rebornNeedGameLevel"` // 重生所需最小关卡
    RebornQuestCoinAwardRise float32 `csv:"rebornQuestCoinAwardRise"` //  重生时任务收益增长
    RebornQuestCoinAwardCntLimit int `csv:"rebornQuestCoinAwardCntLimit"` // 重生时任务收益增长增加次数限制
    RebornEmblemAwardRise float32 `csv:"rebornEmblemAwardRise"` // 重生时纹章收益增长
    RebornEmblemAwardCntLimit int `csv:"rebornEmblemAwardCntLimit"` // 重生时纹章增长增加次数限制
    RebornInitCoin UNumber `csv:"rebornInitCoin"` // 重生时重置金币到这个数量
    RebornDoubleEmblemAwardDiamondCnt UNumber `csv:"rebornDoubleEmblemAwardDiamondCnt"` // 重生时双倍纹章奖励需要的钻石数量
    BattleDropGoldMin int `csv:"battleDropGoldMin"` // 战斗掉落金币最小个数
    BattleDropGoldMax int `csv:"battleDropGoldMax"` // 战斗掉落金币最大个数
    BattleDropGoldLowerLimit int `csv:"battleDropGoldLowerLimit"` // 战斗掉落倍数下限
    BattleDropGoldUpperLimit int `csv:"battleDropGoldUpperLimit"` // 战斗掉落倍数上限
    ShopBoxRefreshTime int `csv:"shopBoxRefreshTime"` // 限时商店刷新时间（x小时刷新一次)
    AchieveGameLevelRechargeID string `csv:"achieveGameLevelRechargeID"` // 【等级礼包】等级礼包充值id
    AchieveGameLevelValue int `csv:"achieveGameLevelValue"` // 【等级礼包】价值多少钻石
    FocusWeChatAwards string `csv:"focusWeChatAwards"` // 【分享】关注公众号奖励id
    ShareGameLevels int `csv:"shareGameLevels"` // 【分享】达到x关才算邀请成功
    SystemChannel int `csv:"systemChannel"` // 【聊天】系统频道最多条数
    ChatChannel int `csv:"chatChannel"` // 【聊天】聊天频道最多条数
    GeneralChannel int `csv:"generalChannel"` // 【聊天】综合频道最多条数
    ChatInterval int `csv:"chatInterval"` // 【聊天】聊天时间间隔
    DayRedPoint int `csv:"dayRedPoint"` // 【红点】前x天活动功能需要有红点
    DungeonChallengeNum int `csv:"dungeonChallengeNum"` // 【副本】初始挑战次数
    DungeonBuyChallengeNum int `csv:"dungeonBuyChallengeNum"` // 【副本】初始可购买挑战次数
    DungeonNumCD int `csv:"dungeonNumCD"` // 【副本】x秒恢复挑战次数
    ArenaChallengeNum int `csv:"arenaChallengeNum"` // 【竞技场】初始挑战次数
    ArenaBuyChallengeNum int `csv:"arenaBuyChallengeNum"` // 【竞技场】初始可购买挑战次数
    ArenaNumCD int `csv:"arenaNumCD"` // 【竞技场】x秒恢复挑战次数
    ArenaAttackCoe int `csv:"arenaAttackCoe"` // 【竞技力】攻击力系数
    ArenaPhysicDefenseCoe int `csv:"arenaPhysicDefenseCoe"` // 【竞技力】物防系数
    ArenaMagicDefenseCoe int `csv:"arenaMagicDefenseCoe"` // 【竞技力】魔防系数
    ArenaHpCoe int `csv:"arenaHpCoe"` // 【竞技力】生命系数
    ArenaConversionCoe int `csv:"arenaConversionCoe"` // 【竞技力】转化k值
    LotteryFreeNum int `csv:"lotteryFreeNum"` // 【抽奖】每日初始免费次上限数
    AddLotteryNumInterval int `csv:"addLotteryNumInterval"` // 【抽奖】每X关增加一次免费抽奖次数
    ShareAwardID string `csv:"shareAwardID"` // 【分享】分享每日奖励ID
    ShareAwardNum int `csv:"shareAwardNum"` // 【分享】分享每日奖励次数
    FlauntDiamonds int `csv:"flauntDiamonds"` // 【炫耀】炫耀每日奖励钻石
    FlauntNum int `csv:"flauntNum"` // 【炫耀】炫耀每日奖励次数
    ShowArchieveID string `csv:"showArchieveID"` // 【关卡成就】关卡成就界面单独显示奖励的关卡ID
    SuggestRebornParameter float32 `csv:"suggestRebornParameter"` // 【建议重生】历史最高纪录x%退关时候提示
    GameLevelAccelerateBuffCfgID string `csv:"gameLevelAccelerateBuffCfgId"` // 10关有概率获取系统加速的buffID
    GameLevelAccelerateDuration int `csv:"gameLevelAccelerateDuration"` // 10关有概率获取系统加速的时间，单位为秒
    VideoReleaseNeedGameLevel int `csv:"videoReleaseNeedGameLevel"` // 【视频发布】弹出tips至少关卡
    VideoReleaseAwardsCnt int `csv:"videoReleaseAwardsCnt"` // 【视频发布】奖励钻石数量
    HeroID string `csv:"heroID"` // 【分享奖励】英雄id
    AwardsID string `csv:"awardsID"` // 【分享奖励】奖励id
    SearchSmallProgramAwardsID string `csv:"searchSmallProgramAwardsID"` // 【搜寻小程序】搜索小程序奖励id
    QuestNotClearVipLvl int `csv:"questNotClearVipLvl"` // 【重生任务不清零】激活所需vip等级
    AutoUpgradeQuestVipLvl int `csv:"autoUpgradeQuestVipLvl"` // 【自动升级任务】可开启所需vip等级
    AutoUpgradeHeroVipLvl int `csv:"autoUpgradeHeroVipLvl"` // 【自动升级英雄】可开启所需vip等级
    PowerSavingMode int `csv:"powerSavingMode"` // 【省电模式】过x分钟音效消失
    GameLevelsFightingCapacityCoefficient float32 `csv:"gameLevelsFightingCapacityCoefficient"` // 【战斗校验】冒险战力系数
    AreanFightingCapacityCoefficient float32 `csv:"areanFightingCapacityCoefficient"` // 【战斗校验】竞技场战力系数
    DungeonFightingCapacityCoefficient float32 `csv:"dungeonFightingCapacityCoefficient"` // 【战斗校验】副本系数
    PassingTime float32 `csv:"passingTime"` // 【战斗校验】校验时间（秒)
    PhoneBinding string `csv:"phoneBinding"` // 【手机绑定】绑定奖励
}

// Guide TODO: game config struct
type Guide struct {
    CfgID string `csv:"cfgId"` // 引导ID
    Description string `csv:"description"` // 描述
    LastGuide int `csv:"lastGuide"` // 前置引导
    NextGuide int `csv:"nextGuide"` // 引导结束触发引导
    KeyStep int `csv:"keyStep"` // 关键步骤
    Condition int `csv:"condition"` // 触发条件
    ConditionValue string `csv:"conditionValue"` // 条件值
    GuideSumStep int `csv:"guideSumStep"` // 引导总步数
    StepArr []string `csv:"step"` // 步骤ID1
}

// GuideStep TODO: game config struct
type GuideStep struct {
    CfgID string `csv:"cfgId"` // 步骤ID
    ForceGuide bool `csv:"forceGuide"` // 是否强制
    GuideType int `csv:"guideType"` // 引导步骤类型 1：箭头 2：阿瑞斯文字 3：箭头+小文字 4.箭头+阿瑞斯文字
    ShowView int `csv:"showView"` // 出现界面类型
    ArrowDir int `csv:"arrowDir"` // 箭头方向 1,上 2,下 3,左 4,右
    SkipTime int `csv:"skipTime"` // 引导时长,x秒后自动跳过。0表示不能自动跳过。
    ArrowRectStr string `csv:"arrowRectStr"` // 箭头点击区域的坐标(x,y,w,h)
    OffsetDir int `csv:"offsetDir"` // 偏移方向(用于适配后获取正确位置)
    TipsCoordinateStr string `csv:"tipsCoordinateStr"` // 提示框坐标(x,y)
    TipsContent string `csv:"tipsContent"` // 提示框内容
    Description string `csv:"description"` // 描述
    Text1 string `csv:"text1"` // 所在界面1
    EventPointID string `csv:"eventPointID"` // 事件埋点ID
}

// Help TODO: game config struct
type Help struct {
    CfgID string `csv:"cfgId"` // 配置ID
    ExplainArr []string `csv:"explain"` // 标题1
}

// HeroEnhanceLevel TODO: game config struct
type HeroEnhanceLevel struct {
    CfgID string `csv:"cfgId"` // 配置ID：必须是等级数字
    Cost UNumber `csv:"cost"` // 消耗的纹章数量
}

// CostResourceParamsOfHeroEvolution TODO: game config struct
type CostResourceParamsOfHeroEvolution struct {
    ArrcostResourceParams1 string `csv:"arr_costResourceParams_1"` // 超越消耗材料1类型
    ArrcostResourceParams2 int `csv:"arr_costResourceParams_2"` // 超越消耗材料1数量
}
// HeroEvolution TODO: game config struct
type HeroEvolution struct {
    CfgID string `csv:"cfgId"` // 配置ID:必须跟生物表中的id一致
    PropertyRaiseMultiples float32 `csv:"propertyRaiseMultiples"` // 超越等级对属性的提升百分比: 大于等于0
    EmblemRaiseMultiples float32 `csv:"emblemRaiseMultiples"` // 超越等级对获取纹章几率提升百分比: 大于等于0
    MaxLevelIncrease int `csv:"maxLevelIncrease"` // 对角色最大等级的提升数量
    CostResourceParamsArr []*CostResourceParamsOfHeroEvolution `csv:"costResourceParams"` // 超越消耗材料1类型
    Beizhu1 string `csv:"beizhu1"` // 备注1
}

// HeroLevels TODO: game config struct
type HeroLevels struct {
    CfgID string `csv:"cfgId"` // 配置ID：必须是等级数字
    Cost UNumber `csv:"cost"` // 消耗的金币数量
}

// HeroTalents TODO: game config struct
type HeroTalents struct {
    CfgID string `csv:"cfgId"` // 天赋ID
    Name string `csv:"name"` // 天赋名称
    Description string `csv:"description"` // 天赋描述
    Icon string `csv:"icon"` // 天赋图标
    RequiredLevel int `csv:"requiredLevel"` // 天赋激活等级，默认是0
    BuffID string `csv:"buffId"` // 天赋BuffID
    BuffTargetType string `csv:"buffTargetType"` // 天赋Buff作用目标类型， 默认是self
    Enhance int `csv:"enhance"` // 天赋Buff数值提升，默认是0
    Beizhu1 string `csv:"beizhu1"` // 备注1
    Beizhu2 string `csv:"beizhu2"` // 备注2
    Beizhu3 string `csv:"beizhu3"` // 备注3
    Beizhu4 string `csv:"beizhu4"` // 备注4
    Beizhu5 string `csv:"beizhu5"` // 备注5
    Beizhu6 string `csv:"beizhu6"` // 备注6
    Beizhu7 string `csv:"beizhu7"` // 备注7
    Beizhu8 string `csv:"beizhu8"` // 备注8
    Beizhu9 string `csv:"beizhu9"` // 备注9
    Beizhu10 string `csv:"beizhu10"` // 备注10
    Beizhu11 string `csv:"beizhu11"` // 备注11
}

// Items TODO: game config struct
type Items struct {
    CfgID string `csv:"cfgId"` // 道具ID
    Name string `csv:"name"` // 物品名称
    ItemType string `csv:"itemType"` // 物品类型:支持的类型如下：coin，useableItem，speedItem， diamond
    Description string `csv:"description"` // 物品描述
    Atlas string `csv:"atlas"` // 物品Icon所在图集
    Icon string `csv:"icon"` // 物品Icon (填图片名)
    Quality int `csv:"quality"` // 物品品质 (1~5)，不填默认0级
    StackCount int `csv:"stackCount"` // 可堆叠数量
    ButtonName string `csv:"buttonName"` // 跳转链接按钮名称 (物品详情中，获取物品的跳转按钮名称）
    CommandScheme string `csv:"commandScheme"` // 点击按钮时跳转的界面： instance shopBox shopMagic shopPrivilege shopRecharge arena dungeon
    BackpackTag int `csv:"backpackTag"` // 背包排序索引，例如数值越小的，排在前面
    AwardID string `csv:"awardId"` // 奖励配置ID（对应宝盒打开有概率掉落物品）
    SpellsParamsArr []string `csv:"spellsParams"` // SpellID1
}

// Lottery TODO: game config struct
type Lottery struct {
    CfgID string `csv:"cfgId"` // 奖励id
    Probability int `csv:"probability"` // 概率
    AwardType string `csv:"awardType"` // 奖励1表 creatures, equipments, items, 
    AwardID string `csv:"awardID"` // 奖励id
    Cnt int `csv:"cnt"` // 数量
    Broadcast bool `csv:"broadcast"` // 是否发广播
}

// LotteryBox TODO: game config struct
type LotteryBox struct {
    CfgID string `csv:"cfgId"` // ID
    LotteryCnt int `csv:"lotteryCnt"` // 抽奖次数
    Icon1 string `csv:"icon1"` // 奖励icon
    Icon2 string `csv:"icon2"` // 领奖后图片
    AwardID string `csv:"awardID"` // 奖励id
}

// LotteryRankAward TODO: game config struct
type LotteryRankAward struct {
    CfgID string `csv:"cfgId"` // ID
    OffRank int `csv:"offRank"` // 竞技场排名下限
    AwardID string `csv:"awardID"` // 固定奖励id
    RandomAwardID string `csv:"randomAwardID"` // 固定奖励id
}

// DescriptionOfMailRewardCfg TODO: game config struct
type DescriptionOfMailRewardCfg struct {
    Arrdescription1 string `csv:"arr_description_1"` // 邮件内容通配1
    Arrdescription2 string `csv:"arr_description_2"` // 邮件内容通配1描述
}
// PropOfMailRewardCfg TODO: game config struct
type PropOfMailRewardCfg struct {
    Arrprop1 string `csv:"arr_prop_1"` // 道具1ID类型
    Arrprop2 string `csv:"arr_prop_2"` // 道具1ID
    Arrprop3 int `csv:"arr_prop_3"` // 道具1数量
    Arrprop4 int `csv:"arr_prop_4"` // 道具1单位
}
// MailRewardCfg TODO: game config struct
type MailRewardCfg struct {
    CfgID string `csv:"cfgId"` // 邮件ID
    Mailtype int `csv:"mail_type"` // 邮件类型
    Title string `csv:"title"` // 邮件标题
    Content string `csv:"content"` // 邮件内容
    DescriptionArr []*DescriptionOfMailRewardCfg `csv:"description"` // 邮件内容通配1
    Mailicon string `csv:"mail_icon"` // 邮件图标
    PropArr []*PropOfMailRewardCfg `csv:"prop"` // 道具1ID类型
}

// OperateItems TODO: game config struct
type OperateItems struct {
    CfgID string `csv:"cfgId"` // 道具ID
    Name string `csv:"name"` // 物品名称
    ItemType string `csv:"itemType"` // 物品类型:支持的类型如下：coin，useableItem，speedItem， diamond
    Description string `csv:"description"` // 物品描述
    Atlas string `csv:"atlas"` // 物品Icon所在图集
    Icon string `csv:"icon"` // 物品Icon (填图片名)
    AwardID string `csv:"awardId"` // 奖励配置ID（对应宝盒打开有概率掉落物品）
    PrivilegesArr []string `csv:"privileges"` // 特权1
}

// UpgradeParamsOfPlayerSpells TODO: game config struct
type UpgradeParamsOfPlayerSpells struct {
    ArrupgradeParams1 string `csv:"arr_upgradeParams_1"` // 技能升级材料1的ID（必然在items表)
    ArrupgradeParams2 int `csv:"arr_upgradeParams_2"` // 技能升级材料1数量
}
// PlayerSpells TODO: game config struct
type PlayerSpells struct {
    CfgID string `csv:"cfgId"` // 主动技能ID
    Name string `csv:"name"` // 技能名称
    Description string `csv:"description"` // 技能描述
    Icon string `csv:"icon"` // 技能图标
    UnLockCfgID string `csv:"unLockCfgId"` // 技能解锁ID
    Quality int `csv:"quality"` // 品质
    SpellType string `csv:"spellType"` // 类型(同技能等级不同分为一类）
    Level int `csv:"level"` // 技能等级
    SpellID string `csv:"spellId"` // 技能ID
    UpgradeParamsArr []*UpgradeParamsOfPlayerSpells `csv:"upgradeParams"` // 技能升级材料1的ID（必然在items表)
    NeedGameLevel int `csv:"needGameLevel"` // 所需关卡
    NeedCreaturesNum int `csv:"needCreaturesNum"` // 技能所需英雄数量
    CreaturesTypeArr []string `csv:"creaturesType"` // 激活所需英雄1
    UpgradeSpellsID string `csv:"upgradeSpellsID"` // 升级后的技能id
    Awards string `csv:"awards"` // 奖励ID
}

// Privilege TODO: game config struct
type Privilege struct {
    CfgID string `csv:"cfgId"` // 特权id
    Name string `csv:"name"` // 特权名称
    DescribeValue string `csv:"describeValue"` // 显示数值
    TypeID int `csv:"typeID"` // 特权类型(同类相同)
    TypeLevel int `csv:"typeLevel"` // 等级
    DurationTime int `csv:"durationTime"` // 持续时间 1.永久 2.每天重置
    Value float32 `csv:"value"` // 特权数值
    ClientType int `csv:"clientType"` // 特权分类 1：商店特权 2：VIP特权 3：成就特权 4：VIP特权
    BuffID1 string `csv:"buffID1"` // buff1ID
    BuffTarget1 string `csv:"buffTarget1"` // buff1目标 (self、target、system)
    BuffID2 string `csv:"buffID2"` // buff2ID
    BuffTarget2 string `csv:"buffTarget2"` // buff2目标 (self、target、system)
}

// PvpArena TODO: game config struct
type PvpArena struct {
    CfgID string `csv:"cfgId"` // ID
    OffRank int `csv:"offRank"` // 竞技场排名下限
    OnRank int `csv:"onRank"` // 竞技场排名上限
    AwardID string `csv:"awardID"` // 胜利资源类型
}

// PvpArenaRank TODO: game config struct
type PvpArenaRank struct {
    CfgID string `csv:"cfgId"` // ID
    OffRank int `csv:"offRank"` // 竞技场排名下限
    OnRank int `csv:"onRank"` // 竞技场排名上限
    JumpUpNumMin int `csv:"jumpUpNumMin"` // 挑战位置间隔数量下限
    JumpUpNumMax int `csv:"jumpUpNumMax"` // 挑战位置间隔数量上线
}

// PvpArenaRankAward TODO: game config struct
type PvpArenaRankAward struct {
    CfgID string `csv:"cfgId"` // ID
    OffRank int `csv:"offRank"` // 竞技场排名下限
    OnRank int `csv:"onRank"` // 竞技场排名上限
    AwardID string `csv:"awardID"` // 奖励id
}

// Quests TODO: game config struct
type Quests struct {
    CfgID string `csv:"cfgId"` // 任务id
    Name string `csv:"name"` // 任务名称
    Icon string `csv:"icon"` // 图标
    ActivateCostCoin UNumber `csv:"activateCostCoin"` // 激活需要的金币
    Duration int `csv:"duration"` // 任务过程时长，单位是秒
    UpgradeCostCoin UNumber `csv:"upgradeCostCoin"` // 任务升级消耗金币
    AwardCoin UNumber `csv:"awardCoin"` // 任务奖励金币
    AwardCoinLevel UNumber `csv:"awardCoinLevel"` // 任务奖励金币等级加成
    AsAutoCostCoin UNumber `csv:"asAutoCostCoin"` // 变为自动任务需要消耗的金币
    AsAutoCostDiamond UNumber `csv:"asAutoCostDiamond"` // 变为自动任务需要消耗的钻石
    MaxLevel int `csv:"maxLevel"` // 最高等级
    Tips string `csv:"tips"` // 自动任务花钻石解锁提示
}

// Reborn TODO: game config struct
type Reborn struct {
    CfgID string `csv:"cfgId"` // 重生关卡数
    EmblemCount UNumber `csv:"emblemCount"` // 纹章奖励数量
}

// RechargeConfig TODO: game config struct
type RechargeConfig struct {
    CfgID string `csv:"cfgId"` // id
    Price int `csv:"price"` // 多少钱（人民币）
    ChannelID string `csv:"channelId"` // 需要在哪个渠道里显示，用下划线分割多个渠道(1、默认)
    ProductName string `csv:"productName"` // 商品名字
    IsFirstRecharge bool `csv:"isFirstRecharge"` // 是否计算首充
    RechargeType string `csv:"rechargeType"` // 充值类型 1:月卡 monthcard 2:年卡 yearcard 3:永久卡 forevercard 4:商店  shop 5:等级礼包  activityLevelAward 6:礼包商店 shopGift 7:货币商店 shopCurrency 
    RechargeCfgID string `csv:"rechargeCfgId"` // 配置ID
    Note string `csv:"note"` // 备注
    CommodityID string `csv:"commodityId"` // 苹果充值商品id
    MailTitle string `csv:"mailTitle"` // 邮件标题
    MailContent string `csv:"mailContent"` // 邮件内容
    Atlas string `csv:"atlas"` // 礼包icon所在图集
    Icon string `csv:"icon"` // 物品Icon (填图片名)
    AppID int `csv:"appId"` // AppID
}

// ShareAchieve TODO: game config struct
type ShareAchieve struct {
    CfgID string `csv:"cfgId"` // 成就id
    Name string `csv:"name"` // 功能名
    Description string `csv:"description"` // 成就描述
    AchieveType int `csv:"achieveType"` // 成就类型 1,召集 2,助战 3,分享
    AchieveValue int `csv:"achieveValue"` // 成就值
    AchieveAwards string `csv:"achieveAwards"` // 成就奖励
    PrivilegeID string `csv:"privilegeID"` // 成就特权
}

// ShareSupply TODO: game config struct
type ShareSupply struct {
    CfgID string `csv:"cfgId"` // 补给
    Name string `csv:"name"` // 补给名
    AwardsID string `csv:"awardsID"` // 奖励id
}

// ShopBox TODO: game config struct
type ShopBox struct {
    CfgID string `csv:"cfgId"` // 商品id
    GroupType int `csv:"groupType"` // 页签类型: 1.限时宝箱 2.英雄宝箱 3.装备宝箱
    TypeID int `csv:"typeID"` // 商品类型id
    RefreshType int `csv:"refreshType"` // 刷新类型 （1、永不刷新 2、每隔x小时刷新 3、每天刷新 4、每周刷新 5、每月刷新）
    Probability int `csv:"probability"` // 出现概率（同typeID加起来10000)
    RechargeType int `csv:"rechargeType"` // 类型： 1.观看视频 2.付费 3.分享
    CostResourceID string `csv:"costResourceID"` // 消耗资源类型
    CostNum UNumber `csv:"costNum"` // 消耗资源数量
    BuyCount int `csv:"buyCount"` // 可购买次数（0为无限)
    AwardID1 string `csv:"awardID1"` // 奖励ID1
    AwardID2 string `csv:"awardID2"` // 奖励ID2
    AwardID3 string `csv:"awardID3"` // 奖励ID3
    MinNum int `csv:"minNum"` // 抽x~y次的x
    MaxNum int `csv:"maxNum"` // 抽x~y次的y
    AwardID4 string `csv:"awardID4"` // 低保奖池 （出现低保替换奖励ID3)
    AwardID5 string `csv:"awardID5"` // 首次奖池 (出现低保替换奖励ID3)
    CoolDown int `csv:"coolDown"` // 冷却时间（秒)
}

// ShopGift TODO: game config struct
type ShopGift struct {
    CfgID string `csv:"cfgId"` // 商品id
    GroupType int `csv:"groupType"` // 页签类型: 1.限时礼包 2.指定英雄礼包 3.新手礼包 4.勋章礼包 5.副本材料礼包 6.抽奖券礼包
    GiftType int `csv:"giftType"` // 类型： 1.周礼包 2.月礼包 3.指定英雄礼包 4.新手礼包 5.勋章礼包 6.副本材料礼包 7.抽奖券礼包
    RefreshType int `csv:"refreshType"` // 刷新类型 （1、永不刷新 2、每隔x小时刷新 3、每天刷新 4、每周刷新 5、每月刷新）
    AwardID string `csv:"awardID"` // 奖励id
    RechargeID string `csv:"rechargeID"` // 充值ID
    Recharge int `csv:"recharge"` // 充值金额(人民币)
    BuyCount int `csv:"buyCount"` // 可购买次数（0为无限)
    IsIosdisplay bool `csv:"isIosdisplay"` // IOS渠道上是否显示
}

// ShopMagic TODO: game config struct
type ShopMagic struct {
    CfgID string `csv:"cfgId"` // 商品id
    GroupType int `csv:"groupType"` // 页签类型: 1.魔法商店
    RefreshType int `csv:"refreshType"` // 刷新类型 （1、永不刷新 2、每隔x小时刷新 3、每天刷新 4、每周刷新 5、每月刷新）
    IconSrc string `csv:"iconSrc"` // 图标所在图集： creatures, equipments, items,shop, award
    Icon string `csv:"icon"` // 图标
    RechargeType int `csv:"rechargeType"` // 类型： 1.观看视频 2.付费 3.分享
    Name string `csv:"name"` // 商品名
    Description string `csv:"description"` // 描述
    CostResourceID string `csv:"costResourceID"` // 消耗资源类型
    CostNum UNumber `csv:"costNum"` // 消耗资源数量
    BuyCount int `csv:"buyCount"` // 可购买次数（0为无限)
    AwardID string `csv:"awardID"` // 奖励id
    DurationTime int `csv:"durationTime"` // 持续时间(秒)
    PrivilegeID string `csv:"privilegeId"` // 特权id
    CoolDown int `csv:"coolDown"` // 冷却时间（秒)
}

// ShopRecharge TODO: game config struct
type ShopRecharge struct {
    CfgID string `csv:"cfgId"` // 商品id
    GroupType int `csv:"groupType"` // 页签类型: 1.钻石包 2.荣誉币包
    RefreshType int `csv:"refreshType"` // 刷新类型 （1、永不刷新 2、每隔x小时刷新 3、每天刷新 4、每周刷新 5、每月刷新）
    IconSrc string `csv:"iconSrc"` // 图标所在图集： creatures, equipments, items, award
    Icon string `csv:"icon"` // 图标
    RechargeType int `csv:"rechargeType"` // 类型： 1.观看视频 2.付费 3.分享
    GetResourceID string `csv:"getResourceID"` // 获取资源类型
    GetNum UNumber `csv:"getNum"` // 获取资源数量(显示用)
    Name string `csv:"name"` // 商品名
    Description string `csv:"description"` // 描述
    RechargeID string `csv:"rechargeID"` // 充值ID
    Recharge int `csv:"recharge"` // 充值金额(人民币)
    BuyCount int `csv:"buyCount"` // 可购买次数（0为无限)
    Percent float32 `csv:"percent"` // 打折百分比
    CoolDown int `csv:"coolDown"` // 冷却时间（秒)
}

// ShopTechnologyTree TODO: game config struct
type ShopTechnologyTree struct {
    CfgID string `csv:"cfgId"` // 科技树id
    GroupType int `csv:"groupType"` // 组id
    TypeID int `csv:"typeID"` // 商品id
    TypeLevel int `csv:"typeLevel"` // 商品等级 同类商品id用等级区分
    RefreshType int `csv:"refreshType"` // 刷新类型 （1、永不刷新 2、每隔x小时刷新 3、每天刷新 4、每周刷新 5、每月刷新）
    IconSrc string `csv:"iconSrc"` // 图标所在图集： creatures, equipments, items,shop, award,
    Icon string `csv:"icon"` // 图标
    Name string `csv:"name"` // 商品名
    Description string `csv:"description"` // 描述
    DescriptionDetails string `csv:"descriptionDetails"` // 描述详情
    CostResourceID string `csv:"costResourceID"` // 消耗资源类型
    CostNum UNumber `csv:"costNum"` // 消耗资源数量
    BuyCount int `csv:"buyCount"` // 可购买次数（0为无限)
    PrivilegeID string `csv:"privilegeId"` // 特权id
    GroupID string `csv:"groupId"` // 特权组id
}

// DescribesOfShopTechnologyTreeGroup TODO: game config struct
type DescribesOfShopTechnologyTreeGroup struct {
    Arrdescribes1 string `csv:"arr_describes_1"` // 特权1图片
    Arrdescribes2 string `csv:"arr_describes_2"` // 特权1文本
}
// ShopTechnologyTreeGroup TODO: game config struct
type ShopTechnologyTreeGroup struct {
    CfgID string `csv:"cfgId"` // ID
    RebornNum int `csv:"rebornNum"` // 重生次数
    DescribesArr []*DescribesOfShopTechnologyTreeGroup `csv:"describes"` // 特权1图片
}

// SignInFirst TODO: game config struct
type SignInFirst struct {
    CfgID string `csv:"cfgId"` // 功能ID
    Name string `csv:"name"` // 功能名
    AwardID string `csv:"awardID"` // 奖励ID
}

// SignInLoop TODO: game config struct
type SignInLoop struct {
    CfgID string `csv:"cfgId"` // 功能ID
    Name string `csv:"name"` // 功能名
    AwardID string `csv:"awardID"` // 奖励ID
    DoubleAward int `csv:"doubleAward"` // VIP翻倍 填写VIP ? 可以翻倍
}

// BuffsParamsOfSpecialAbility TODO: game config struct
type BuffsParamsOfSpecialAbility struct {
    ArrbuffsParams1 string `csv:"arr_buffsParams_1"` // BuffID1
    ArrbuffsParams2 string `csv:"arr_buffsParams_2"` // BuffID1目标类型
}
// SpecialAbility TODO: game config struct
type SpecialAbility struct {
    CfgID string `csv:"cfgId"` // 特殊能力ID
    Name string `csv:"name"` // 特殊能力名称
    Description string `csv:"description"` // 特殊能力描述
    BuffsParamsArr []*BuffsParamsOfSpecialAbility `csv:"buffsParams"` // BuffID1
}

// SpellPresentation TODO: game config struct
type SpellPresentation struct {
    CfgID string `csv:"cfgId"` // 配置ID
    AttackAction string `csv:"attackAction"` // 施放动作
    BeAttackedAction string `csv:"beAttackedAction"` // 受击动作
    AttackArtEffect string `csv:"attackArtEffect"` // 施放美术特效
    CastPosition string `csv:"castPosition"` // 施放美术特效位置
    MotionArtEffect string `csv:"motionArtEffect"` // 弹道施放美术特效
    MotionType string `csv:"motionType"` // 弹道类型
    MotionDuration float32 `csv:"motionDuration"` // 弹道施放时长
    BeAttackArtEffect string `csv:"beAttackArtEffect"` // 受击美术特效
    OntoPosition string `csv:"ontoPosition"` // 受击美术特效位置
    SpellReleaseEffect string `csv:"spellReleaseEffect"` // 技能释放音效
    CreatureHitEffect string `csv:"creatureHitEffect"` // 生物受击音效
}

// BuffsParamsOfSpells TODO: game config struct
type BuffsParamsOfSpells struct {
    ArrbuffsParams1 string `csv:"arr_buffsParams_1"` // BuffID1
    ArrbuffsParams2 float32 `csv:"arr_buffsParams_2"` // BuffID1触发概率（大于等于0，且小于等于1）
}
// Spells TODO: game config struct
type Spells struct {
    CfgID string `csv:"cfgId"` // 技能ID
    Name string `csv:"name"` // 技能名称
    SpellType1 string `csv:"spellType1"` // 攻击目标类型
    SpellType2 string `csv:"spellType2"` // 攻击目标类型
    Icon string `csv:"icon"` // 技能图标
    Description string `csv:"description"` // 技能描述
    Presentation string `csv:"presentation"` // 技能表现配置ID
    ApperanceMultiple float32 `csv:"apperanceMultiple"` // 技能动作时间倍率，默认1倍
    Cooldown int `csv:"cooldown"` // 冷却时间（单位毫秒）
    AttackValue2 UNumber `csv:"attackValue2"` // 当技能挂在非生物身上时的攻击力
    AttackType2 string `csv:"attackType2"` // 当技能挂在非生物身上时的攻击类型，候选如下： physic magic
    ManaDecrease UNumber `csv:"manaDecrease"` // 怒气消耗
    ManaIncrease UNumber `csv:"manaIncrease"` // 怒气增加（加到hero身上）
    CastTargetSet string `csv:"castTargetSet"` // 施放目标集合，不填表示不过滤，候选如下： enemy allies
    CastTargetType string `csv:"castTargetType"` // 施放目标类型，不填默认等于"select"，候选如下：         select,         column,         front,         back,         all,         minHpHero,         rand, bounce
    RandCount int `csv:"randCount"` // 随机个数，castTargetType设置： 1.为rand时，这里填写随机个数 2.如果是bounce,填写弹跳次数
    CastTroopType string `csv:"castTroopType"` // 施放目标的兵种过滤，不填表示不过滤，候选如下：         air,         land,
    CastTargetAttackType string `csv:"castTargetAttackType"` // 施放目标的攻击类型过滤，不填表示不过滤，候选如下：         magic,         physic,
    CastTargetSex string `csv:"castTargetSex"` // 施放目标的性别过滤，不填表示不过滤，候选如下：         female,         male
    DamageMultiple UNumber `csv:"damageMultiple"` // 伤害倍数 （大于等于0)，不填或者填0表示没有伤害，后面若干个参数会被忽略
    IgnoreDefenceRate float32 `csv:"ignoreDefenceRate"` // 无视防御百分比 （大于等于0，且小于等于1)
    IgnoreDefenceDamage UNumber `csv:"ignoreDefenceDamage"` // 固定伤害
    Dark bool `csv:"dark"` // 是否黑屏
    SpecialEffects string `csv:"specialEffects"` // 第三阶段特殊效果，例如震屏,多个就逗号分隔，现在支持的:shock, black
    BuffsParamsArr []*BuffsParamsOfSpells `csv:"buffsParams"` // BuffID1
    Beizhu1 string `csv:"beizhu1"` // 备注
}

// SystemLock TODO: game config struct
type SystemLock struct {
    CfgID string `csv:"cfgId"` // 功能ID
    Name string `csv:"name"` // 功能名
    Description string `csv:"description"` // 功能描述
    Tips string `csv:"tips"` // 未解锁时提示
    Icon string `csv:"icon"` // 功能ICON
    Atlas string `csv:"atlas"` // 图集所在目录 技能spell 其他main
    ButtonShowGameLevels int `csv:"buttonShowGameLevels"` // 按钮显示关数
    UnlockGameLevels int `csv:"unlockGameLevels"` // 按钮解锁关数
    NextCfgID string `csv:"nextCfgId"` // 下个功能解锁的ID
    IsCostDiamond bool `csv:"isCostDiamond"` // 是否消耗钻石解锁，不填默认是false
    Cost UNumber `csv:"cost"` // 消耗钻石
    IsShowWindows bool `csv:"isShowWindows"` // 是否有弹窗
    OpenInWechat bool `csv:"openInWechat"` // 在微信中是否开放
    OpenInHeadLine bool `csv:"openInHeadLine"` // 在头条中是否开放
}

// Vipconfig TODO: game config struct
type Vipconfig struct {
    CfgID string `csv:"cfgId"` // ID
    Name string `csv:"name"` // 名称
    Icon string `csv:"icon"` // 游戏icon
    AwardsID string `csv:"awardsID"` // 礼包id
    VipLevel int `csv:"vipLevel"` // vip等级
    ConditionNumber int `csv:"conditionNumber"` // 需要完成条件数量
    PrivilegeMainDesc string `csv:"privilegeMainDesc"` // 主要特权描述
    PrivilegeDescArr []string `csv:"privilegeDesc"` // 特权描述1
    ConditionGroupID string `csv:"conditionGroupID"` // 条件所在的组ID（在achieveVip中查找）
    PrivilegeIDArr []string `csv:"privilegeID"` // 特权1
}

// WelfareCard TODO: game config struct
type WelfareCard struct {
    CfgID string `csv:"cfgId"` // ID
    RechargeID string `csv:"rechargeID"` // 充值id
    Type string `csv:"type"` // 类型:monthcard,yearcard,forevercard
    Awards string `csv:"awards"` // 立即奖励
    EverydayAwards string `csv:"everydayAwards"` // 每日可领
    Privilege1 string `csv:"privilege1"` // 特权1
    Privilege2 string `csv:"privilege2"` // 特权2
    Privilege3 string `csv:"privilege3"` // 特权3
    Privilege4 string `csv:"privilege4"` // 特权4
    Privilege5 string `csv:"privilege5"` // 特权5
}

