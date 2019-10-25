package gamecfg

import (
    "encoding/csv"
    "bufio"
    "github.com/sirupsen/logrus"
)
var (
    // AppConfigMap TODO
    AppConfigMap map[string]*AppConfig
    AppConfigArray []*AppConfig
    // ApplePayProductIDCfgMap TODO
    ApplePayProductIDCfgMap map[string]*ApplePayProductIDCfg
    ApplePayProductIDCfgArray []*ApplePayProductIDCfg
    // CompensationAwardCfgMap TODO
    CompensationAwardCfgMap map[string]*CompensationAwardCfg
    CompensationAwardCfgArray []*CompensationAwardCfg
    // ContinueRechargeCfgMap TODO
    ContinueRechargeCfgMap map[string]*ContinueRechargeCfg
    ContinueRechargeCfgArray []*ContinueRechargeCfg
    // CumulateConsumeCfgMap TODO
    CumulateConsumeCfgMap map[string]*CumulateConsumeCfg
    CumulateConsumeCfgArray []*CumulateConsumeCfg
    // CumulateRechargeCfgMap TODO
    CumulateRechargeCfgMap map[string]*CumulateRechargeCfg
    CumulateRechargeCfgArray []*CumulateRechargeCfg
    // CumulateRechargeForeverCfgMap TODO
    CumulateRechargeForeverCfgMap map[string]*CumulateRechargeForeverCfg
    CumulateRechargeForeverCfgArray []*CumulateRechargeForeverCfg
    // LotteryRankRewardCfgMap TODO
    LotteryRankRewardCfgMap map[string]*LotteryRankRewardCfg
    LotteryRankRewardCfgArray []*LotteryRankRewardCfg
    // AchieveMap TODO
    AchieveMap map[string]*Achieve
    AchieveArray []*Achieve
    // AchieveBoxMap TODO
    AchieveBoxMap map[string]*AchieveBox
    AchieveBoxArray []*AchieveBox
    // AchieveGameLevelMap TODO
    AchieveGameLevelMap map[string]*AchieveGameLevel
    AchieveGameLevelArray []*AchieveGameLevel
    // AchieveVIPMap TODO
    AchieveVIPMap map[string]*AchieveVIP
    AchieveVIPArray []*AchieveVIP
    // ActivityLevelAwardMap TODO
    ActivityLevelAwardMap map[string]*ActivityLevelAward
    ActivityLevelAwardArray []*ActivityLevelAward
    // AdboxMap TODO
    AdboxMap map[string]*Adbox
    AdboxArray []*Adbox
    // AnnounceConditionMap TODO
    AnnounceConditionMap map[string]*AnnounceCondition
    AnnounceConditionArray []*AnnounceCondition
    // AppDungeonActiviesAchieveMap TODO
    AppDungeonActiviesAchieveMap map[string]*AppDungeonActiviesAchieve
    AppDungeonActiviesAchieveArray []*AppDungeonActiviesAchieve
    // AppDungeonActiviesRankAwardMap TODO
    AppDungeonActiviesRankAwardMap map[string]*AppDungeonActiviesRankAward
    AppDungeonActiviesRankAwardArray []*AppDungeonActiviesRankAward
    // AppEquipmentSuitActiviesAchieveMap TODO
    AppEquipmentSuitActiviesAchieveMap map[string]*AppEquipmentSuitActiviesAchieve
    AppEquipmentSuitActiviesAchieveArray []*AppEquipmentSuitActiviesAchieve
    // AppEquipmentSuitActiviesRankAwardMap TODO
    AppEquipmentSuitActiviesRankAwardMap map[string]*AppEquipmentSuitActiviesRankAward
    AppEquipmentSuitActiviesRankAwardArray []*AppEquipmentSuitActiviesRankAward
    // AppNewServiceActiviesAchieveMap TODO
    AppNewServiceActiviesAchieveMap map[string]*AppNewServiceActiviesAchieve
    AppNewServiceActiviesAchieveArray []*AppNewServiceActiviesAchieve
    // AppNewServiceActiviesAchieveShopMap TODO
    AppNewServiceActiviesAchieveShopMap map[string]*AppNewServiceActiviesAchieveShop
    AppNewServiceActiviesAchieveShopArray []*AppNewServiceActiviesAchieveShop
    // AppNewServiceActiviesDailyObjMap TODO
    AppNewServiceActiviesDailyObjMap map[string]*AppNewServiceActiviesDailyObj
    AppNewServiceActiviesDailyObjArray []*AppNewServiceActiviesDailyObj
    // AppNewServiceActiviesRankAwardMap TODO
    AppNewServiceActiviesRankAwardMap map[string]*AppNewServiceActiviesRankAward
    AppNewServiceActiviesRankAwardArray []*AppNewServiceActiviesRankAward
    // AreanRobotMap TODO
    AreanRobotMap map[string]*AreanRobot
    AreanRobotArray []*AreanRobot
    // AwardsMap TODO
    AwardsMap map[string]*Awards
    AwardsArray []*Awards
    // BattleConfigMap TODO
    BattleConfigMap map[string]*BattleConfig
    BattleConfigArray []*BattleConfig
    // BattleNumMap TODO
    BattleNumMap map[string]*BattleNum
    BattleNumArray []*BattleNum
    // BornMap TODO
    BornMap map[string]*Born
    BornArray []*Born
    // BuffsMap TODO
    BuffsMap map[string]*Buffs
    BuffsArray []*Buffs
    // CreaturesMap TODO
    CreaturesMap map[string]*Creatures
    CreaturesArray []*Creatures
    // DungeonMap TODO
    DungeonMap map[string]*Dungeon
    DungeonArray []*Dungeon
    // DungeonLevelMap TODO
    DungeonLevelMap map[string]*DungeonLevel
    DungeonLevelArray []*DungeonLevel
    // EffectsMap TODO
    EffectsMap map[string]*Effects
    EffectsArray []*Effects
    // EquipmentSuitPropertiesMap TODO
    EquipmentSuitPropertiesMap map[string]*EquipmentSuitProperties
    EquipmentSuitPropertiesArray []*EquipmentSuitProperties
    // EquipmentSuitsMap TODO
    EquipmentSuitsMap map[string]*EquipmentSuits
    EquipmentSuitsArray []*EquipmentSuits
    // EquipmentsMap TODO
    EquipmentsMap map[string]*Equipments
    EquipmentsArray []*Equipments
    // EquipmentsEvolutionMap TODO
    EquipmentsEvolutionMap map[string]*EquipmentsEvolution
    EquipmentsEvolutionArray []*EquipmentsEvolution
    // EventPointMap TODO
    EventPointMap map[string]*EventPoint
    EventPointArray []*EventPoint
    // FirstRechargeMap TODO
    FirstRechargeMap map[string]*FirstRecharge
    FirstRechargeArray []*FirstRecharge
    // GameLevelsMap TODO
    GameLevelsMap map[string]*GameLevels
    GameLevelsArray []*GameLevels
    // GlobalMap TODO
    GlobalMap map[string]*Global
    GlobalArray []*Global
    // GuideMap TODO
    GuideMap map[string]*Guide
    GuideArray []*Guide
    // GuideStepMap TODO
    GuideStepMap map[string]*GuideStep
    GuideStepArray []*GuideStep
    // HelpMap TODO
    HelpMap map[string]*Help
    HelpArray []*Help
    // HeroEnhanceLevelMap TODO
    HeroEnhanceLevelMap map[string]*HeroEnhanceLevel
    HeroEnhanceLevelArray []*HeroEnhanceLevel
    // HeroEvolutionMap TODO
    HeroEvolutionMap map[string]*HeroEvolution
    HeroEvolutionArray []*HeroEvolution
    // HeroLevelsMap TODO
    HeroLevelsMap map[string]*HeroLevels
    HeroLevelsArray []*HeroLevels
    // HeroTalentsMap TODO
    HeroTalentsMap map[string]*HeroTalents
    HeroTalentsArray []*HeroTalents
    // ItemsMap TODO
    ItemsMap map[string]*Items
    ItemsArray []*Items
    // LotteryMap TODO
    LotteryMap map[string]*Lottery
    LotteryArray []*Lottery
    // LotteryBoxMap TODO
    LotteryBoxMap map[string]*LotteryBox
    LotteryBoxArray []*LotteryBox
    // LotteryRankAwardMap TODO
    LotteryRankAwardMap map[string]*LotteryRankAward
    LotteryRankAwardArray []*LotteryRankAward
    // MailRewardCfgMap TODO
    MailRewardCfgMap map[string]*MailRewardCfg
    MailRewardCfgArray []*MailRewardCfg
    // OperateItemsMap TODO
    OperateItemsMap map[string]*OperateItems
    OperateItemsArray []*OperateItems
    // PlayerSpellsMap TODO
    PlayerSpellsMap map[string]*PlayerSpells
    PlayerSpellsArray []*PlayerSpells
    // PrivilegeMap TODO
    PrivilegeMap map[string]*Privilege
    PrivilegeArray []*Privilege
    // PvpArenaMap TODO
    PvpArenaMap map[string]*PvpArena
    PvpArenaArray []*PvpArena
    // PvpArenaRankMap TODO
    PvpArenaRankMap map[string]*PvpArenaRank
    PvpArenaRankArray []*PvpArenaRank
    // PvpArenaRankAwardMap TODO
    PvpArenaRankAwardMap map[string]*PvpArenaRankAward
    PvpArenaRankAwardArray []*PvpArenaRankAward
    // QuestsMap TODO
    QuestsMap map[string]*Quests
    QuestsArray []*Quests
    // RebornMap TODO
    RebornMap map[string]*Reborn
    RebornArray []*Reborn
    // RechargeConfigMap TODO
    RechargeConfigMap map[string]*RechargeConfig
    RechargeConfigArray []*RechargeConfig
    // ShareAchieveMap TODO
    ShareAchieveMap map[string]*ShareAchieve
    ShareAchieveArray []*ShareAchieve
    // ShareSupplyMap TODO
    ShareSupplyMap map[string]*ShareSupply
    ShareSupplyArray []*ShareSupply
    // ShopBoxMap TODO
    ShopBoxMap map[string]*ShopBox
    ShopBoxArray []*ShopBox
    // ShopGiftMap TODO
    ShopGiftMap map[string]*ShopGift
    ShopGiftArray []*ShopGift
    // ShopMagicMap TODO
    ShopMagicMap map[string]*ShopMagic
    ShopMagicArray []*ShopMagic
    // ShopRechargeMap TODO
    ShopRechargeMap map[string]*ShopRecharge
    ShopRechargeArray []*ShopRecharge
    // ShopTechnologyTreeMap TODO
    ShopTechnologyTreeMap map[string]*ShopTechnologyTree
    ShopTechnologyTreeArray []*ShopTechnologyTree
    // ShopTechnologyTreeGroupMap TODO
    ShopTechnologyTreeGroupMap map[string]*ShopTechnologyTreeGroup
    ShopTechnologyTreeGroupArray []*ShopTechnologyTreeGroup
    // SignInFirstMap TODO
    SignInFirstMap map[string]*SignInFirst
    SignInFirstArray []*SignInFirst
    // SignInLoopMap TODO
    SignInLoopMap map[string]*SignInLoop
    SignInLoopArray []*SignInLoop
    // SpecialAbilityMap TODO
    SpecialAbilityMap map[string]*SpecialAbility
    SpecialAbilityArray []*SpecialAbility
    // SpellPresentationMap TODO
    SpellPresentationMap map[string]*SpellPresentation
    SpellPresentationArray []*SpellPresentation
    // SpellsMap TODO
    SpellsMap map[string]*Spells
    SpellsArray []*Spells
    // SystemLockMap TODO
    SystemLockMap map[string]*SystemLock
    SystemLockArray []*SystemLock
    // VipconfigMap TODO
    VipconfigMap map[string]*Vipconfig
    VipconfigArray []*Vipconfig
    // WelfareCardMap TODO
    WelfareCardMap map[string]*WelfareCard
    WelfareCardArray []*WelfareCard
)
func loadAppConfig(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppConfigMap = make(map[string]*AppConfig)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppConfig{}
        p.unmarshalRow(header, row, av)
        AppConfigMap[av.CfgID] = av
        AppConfigArray = append(AppConfigArray,av)
    }
}
func loadApplePayProductIDCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ApplePayProductIDCfgMap = make(map[string]*ApplePayProductIDCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ApplePayProductIDCfg{}
        p.unmarshalRow(header, row, av)
        ApplePayProductIDCfgMap[av.CfgID] = av
        ApplePayProductIDCfgArray = append(ApplePayProductIDCfgArray,av)
    }
}
func loadCompensationAwardCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    CompensationAwardCfgMap = make(map[string]*CompensationAwardCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &CompensationAwardCfg{}
        p.unmarshalRow(header, row, av)
        CompensationAwardCfgMap[av.CfgID] = av
        CompensationAwardCfgArray = append(CompensationAwardCfgArray,av)
    }
}
func loadContinueRechargeCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ContinueRechargeCfgMap = make(map[string]*ContinueRechargeCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ContinueRechargeCfg{}
        p.unmarshalRow(header, row, av)
        ContinueRechargeCfgMap[av.CfgID] = av
        ContinueRechargeCfgArray = append(ContinueRechargeCfgArray,av)
    }
}
func loadCumulateConsumeCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    CumulateConsumeCfgMap = make(map[string]*CumulateConsumeCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &CumulateConsumeCfg{}
        p.unmarshalRow(header, row, av)
        CumulateConsumeCfgMap[av.CfgID] = av
        CumulateConsumeCfgArray = append(CumulateConsumeCfgArray,av)
    }
}
func loadCumulateRechargeCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    CumulateRechargeCfgMap = make(map[string]*CumulateRechargeCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &CumulateRechargeCfg{}
        p.unmarshalRow(header, row, av)
        CumulateRechargeCfgMap[av.CfgID] = av
        CumulateRechargeCfgArray = append(CumulateRechargeCfgArray,av)
    }
}
func loadCumulateRechargeForeverCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    CumulateRechargeForeverCfgMap = make(map[string]*CumulateRechargeForeverCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &CumulateRechargeForeverCfg{}
        p.unmarshalRow(header, row, av)
        CumulateRechargeForeverCfgMap[av.CfgID] = av
        CumulateRechargeForeverCfgArray = append(CumulateRechargeForeverCfgArray,av)
    }
}
func loadLotteryRankRewardCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    LotteryRankRewardCfgMap = make(map[string]*LotteryRankRewardCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &LotteryRankRewardCfg{}
        p.unmarshalRow(header, row, av)
        LotteryRankRewardCfgMap[av.CfgID] = av
        LotteryRankRewardCfgArray = append(LotteryRankRewardCfgArray,av)
    }
}
func loadAchieve(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AchieveMap = make(map[string]*Achieve)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Achieve{}
        p.unmarshalRow(header, row, av)
        AchieveMap[av.CfgID] = av
        AchieveArray = append(AchieveArray,av)
    }
}
func loadAchieveBox(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AchieveBoxMap = make(map[string]*AchieveBox)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AchieveBox{}
        p.unmarshalRow(header, row, av)
        AchieveBoxMap[av.CfgID] = av
        AchieveBoxArray = append(AchieveBoxArray,av)
    }
}
func loadAchieveGameLevel(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AchieveGameLevelMap = make(map[string]*AchieveGameLevel)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AchieveGameLevel{}
        p.unmarshalRow(header, row, av)
        AchieveGameLevelMap[av.CfgID] = av
        AchieveGameLevelArray = append(AchieveGameLevelArray,av)
    }
}
func loadAchieveVIP(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AchieveVIPMap = make(map[string]*AchieveVIP)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AchieveVIP{}
        p.unmarshalRow(header, row, av)
        AchieveVIPMap[av.CfgID] = av
        AchieveVIPArray = append(AchieveVIPArray,av)
    }
}
func loadActivityLevelAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ActivityLevelAwardMap = make(map[string]*ActivityLevelAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ActivityLevelAward{}
        p.unmarshalRow(header, row, av)
        ActivityLevelAwardMap[av.CfgID] = av
        ActivityLevelAwardArray = append(ActivityLevelAwardArray,av)
    }
}
func loadAdbox(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AdboxMap = make(map[string]*Adbox)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Adbox{}
        p.unmarshalRow(header, row, av)
        AdboxMap[av.CfgID] = av
        AdboxArray = append(AdboxArray,av)
    }
}
func loadAnnounceCondition(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AnnounceConditionMap = make(map[string]*AnnounceCondition)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AnnounceCondition{}
        p.unmarshalRow(header, row, av)
        AnnounceConditionMap[av.CfgID] = av
        AnnounceConditionArray = append(AnnounceConditionArray,av)
    }
}
func loadAppDungeonActiviesAchieve(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppDungeonActiviesAchieveMap = make(map[string]*AppDungeonActiviesAchieve)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppDungeonActiviesAchieve{}
        p.unmarshalRow(header, row, av)
        AppDungeonActiviesAchieveMap[av.CfgID] = av
        AppDungeonActiviesAchieveArray = append(AppDungeonActiviesAchieveArray,av)
    }
}
func loadAppDungeonActiviesRankAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppDungeonActiviesRankAwardMap = make(map[string]*AppDungeonActiviesRankAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppDungeonActiviesRankAward{}
        p.unmarshalRow(header, row, av)
        AppDungeonActiviesRankAwardMap[av.CfgID] = av
        AppDungeonActiviesRankAwardArray = append(AppDungeonActiviesRankAwardArray,av)
    }
}
func loadAppEquipmentSuitActiviesAchieve(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppEquipmentSuitActiviesAchieveMap = make(map[string]*AppEquipmentSuitActiviesAchieve)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppEquipmentSuitActiviesAchieve{}
        p.unmarshalRow(header, row, av)
        AppEquipmentSuitActiviesAchieveMap[av.CfgID] = av
        AppEquipmentSuitActiviesAchieveArray = append(AppEquipmentSuitActiviesAchieveArray,av)
    }
}
func loadAppEquipmentSuitActiviesRankAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppEquipmentSuitActiviesRankAwardMap = make(map[string]*AppEquipmentSuitActiviesRankAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppEquipmentSuitActiviesRankAward{}
        p.unmarshalRow(header, row, av)
        AppEquipmentSuitActiviesRankAwardMap[av.CfgID] = av
        AppEquipmentSuitActiviesRankAwardArray = append(AppEquipmentSuitActiviesRankAwardArray,av)
    }
}
func loadAppNewServiceActiviesAchieve(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppNewServiceActiviesAchieveMap = make(map[string]*AppNewServiceActiviesAchieve)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppNewServiceActiviesAchieve{}
        p.unmarshalRow(header, row, av)
        AppNewServiceActiviesAchieveMap[av.CfgID] = av
        AppNewServiceActiviesAchieveArray = append(AppNewServiceActiviesAchieveArray,av)
    }
}
func loadAppNewServiceActiviesAchieveShop(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppNewServiceActiviesAchieveShopMap = make(map[string]*AppNewServiceActiviesAchieveShop)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppNewServiceActiviesAchieveShop{}
        p.unmarshalRow(header, row, av)
        AppNewServiceActiviesAchieveShopMap[av.CfgID] = av
        AppNewServiceActiviesAchieveShopArray = append(AppNewServiceActiviesAchieveShopArray,av)
    }
}
func loadAppNewServiceActiviesDailyObj(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppNewServiceActiviesDailyObjMap = make(map[string]*AppNewServiceActiviesDailyObj)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppNewServiceActiviesDailyObj{}
        p.unmarshalRow(header, row, av)
        AppNewServiceActiviesDailyObjMap[av.CfgID] = av
        AppNewServiceActiviesDailyObjArray = append(AppNewServiceActiviesDailyObjArray,av)
    }
}
func loadAppNewServiceActiviesRankAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AppNewServiceActiviesRankAwardMap = make(map[string]*AppNewServiceActiviesRankAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AppNewServiceActiviesRankAward{}
        p.unmarshalRow(header, row, av)
        AppNewServiceActiviesRankAwardMap[av.CfgID] = av
        AppNewServiceActiviesRankAwardArray = append(AppNewServiceActiviesRankAwardArray,av)
    }
}
func loadAreanRobot(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AreanRobotMap = make(map[string]*AreanRobot)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &AreanRobot{}
        p.unmarshalRow(header, row, av)
        AreanRobotMap[av.CfgID] = av
        AreanRobotArray = append(AreanRobotArray,av)
    }
}
func loadAwards(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    AwardsMap = make(map[string]*Awards)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Awards{}
        p.unmarshalRow(header, row, av)
        AwardsMap[av.CfgID] = av
        AwardsArray = append(AwardsArray,av)
    }
}
func loadBattleConfig(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    BattleConfigMap = make(map[string]*BattleConfig)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &BattleConfig{}
        p.unmarshalRow(header, row, av)
        BattleConfigMap[av.CfgID] = av
        BattleConfigArray = append(BattleConfigArray,av)
    }
}
func loadBattleNum(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    BattleNumMap = make(map[string]*BattleNum)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &BattleNum{}
        p.unmarshalRow(header, row, av)
        BattleNumMap[av.CfgID] = av
        BattleNumArray = append(BattleNumArray,av)
    }
}
func loadBorn(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    BornMap = make(map[string]*Born)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Born{}
        p.unmarshalRow(header, row, av)
        BornMap[av.CfgID] = av
        BornArray = append(BornArray,av)
    }
}
func loadBuffs(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    BuffsMap = make(map[string]*Buffs)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Buffs{}
        p.unmarshalRow(header, row, av)
        BuffsMap[av.CfgID] = av
        BuffsArray = append(BuffsArray,av)
    }
}
func loadCreatures(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    CreaturesMap = make(map[string]*Creatures)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Creatures{}
        p.unmarshalRow(header, row, av)
        CreaturesMap[av.CfgID] = av
        CreaturesArray = append(CreaturesArray,av)
    }
}
func loadDungeon(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    DungeonMap = make(map[string]*Dungeon)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Dungeon{}
        p.unmarshalRow(header, row, av)
        DungeonMap[av.CfgID] = av
        DungeonArray = append(DungeonArray,av)
    }
}
func loadDungeonLevel(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    DungeonLevelMap = make(map[string]*DungeonLevel)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &DungeonLevel{}
        p.unmarshalRow(header, row, av)
        DungeonLevelMap[av.CfgID] = av
        DungeonLevelArray = append(DungeonLevelArray,av)
    }
}
func loadEffects(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EffectsMap = make(map[string]*Effects)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Effects{}
        p.unmarshalRow(header, row, av)
        EffectsMap[av.CfgID] = av
        EffectsArray = append(EffectsArray,av)
    }
}
func loadEquipmentSuitProperties(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EquipmentSuitPropertiesMap = make(map[string]*EquipmentSuitProperties)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &EquipmentSuitProperties{}
        p.unmarshalRow(header, row, av)
        EquipmentSuitPropertiesMap[av.CfgID] = av
        EquipmentSuitPropertiesArray = append(EquipmentSuitPropertiesArray,av)
    }
}
func loadEquipmentSuits(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EquipmentSuitsMap = make(map[string]*EquipmentSuits)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &EquipmentSuits{}
        p.unmarshalRow(header, row, av)
        EquipmentSuitsMap[av.CfgID] = av
        EquipmentSuitsArray = append(EquipmentSuitsArray,av)
    }
}
func loadEquipments(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EquipmentsMap = make(map[string]*Equipments)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Equipments{}
        p.unmarshalRow(header, row, av)
        EquipmentsMap[av.CfgID] = av
        EquipmentsArray = append(EquipmentsArray,av)
    }
}
func loadEquipmentsEvolution(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EquipmentsEvolutionMap = make(map[string]*EquipmentsEvolution)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &EquipmentsEvolution{}
        p.unmarshalRow(header, row, av)
        EquipmentsEvolutionMap[av.CfgID] = av
        EquipmentsEvolutionArray = append(EquipmentsEvolutionArray,av)
    }
}
func loadEventPoint(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    EventPointMap = make(map[string]*EventPoint)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &EventPoint{}
        p.unmarshalRow(header, row, av)
        EventPointMap[av.CfgID] = av
        EventPointArray = append(EventPointArray,av)
    }
}
func loadFirstRecharge(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    FirstRechargeMap = make(map[string]*FirstRecharge)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &FirstRecharge{}
        p.unmarshalRow(header, row, av)
        FirstRechargeMap[av.CfgID] = av
        FirstRechargeArray = append(FirstRechargeArray,av)
    }
}
func loadGameLevels(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    GameLevelsMap = make(map[string]*GameLevels)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &GameLevels{}
        p.unmarshalRow(header, row, av)
        GameLevelsMap[av.CfgID] = av
        GameLevelsArray = append(GameLevelsArray,av)
    }
}
func loadGlobal(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    GlobalMap = make(map[string]*Global)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Global{}
        p.unmarshalRow(header, row, av)
        GlobalMap[av.CfgID] = av
        GlobalArray = append(GlobalArray,av)
    }
}
func loadGuide(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    GuideMap = make(map[string]*Guide)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Guide{}
        p.unmarshalRow(header, row, av)
        GuideMap[av.CfgID] = av
        GuideArray = append(GuideArray,av)
    }
}
func loadGuideStep(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    GuideStepMap = make(map[string]*GuideStep)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &GuideStep{}
        p.unmarshalRow(header, row, av)
        GuideStepMap[av.CfgID] = av
        GuideStepArray = append(GuideStepArray,av)
    }
}
func loadHelp(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    HelpMap = make(map[string]*Help)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Help{}
        p.unmarshalRow(header, row, av)
        HelpMap[av.CfgID] = av
        HelpArray = append(HelpArray,av)
    }
}
func loadHeroEnhanceLevel(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    HeroEnhanceLevelMap = make(map[string]*HeroEnhanceLevel)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &HeroEnhanceLevel{}
        p.unmarshalRow(header, row, av)
        HeroEnhanceLevelMap[av.CfgID] = av
        HeroEnhanceLevelArray = append(HeroEnhanceLevelArray,av)
    }
}
func loadHeroEvolution(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    HeroEvolutionMap = make(map[string]*HeroEvolution)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &HeroEvolution{}
        p.unmarshalRow(header, row, av)
        HeroEvolutionMap[av.CfgID] = av
        HeroEvolutionArray = append(HeroEvolutionArray,av)
    }
}
func loadHeroLevels(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    HeroLevelsMap = make(map[string]*HeroLevels)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &HeroLevels{}
        p.unmarshalRow(header, row, av)
        HeroLevelsMap[av.CfgID] = av
        HeroLevelsArray = append(HeroLevelsArray,av)
    }
}
func loadHeroTalents(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    HeroTalentsMap = make(map[string]*HeroTalents)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &HeroTalents{}
        p.unmarshalRow(header, row, av)
        HeroTalentsMap[av.CfgID] = av
        HeroTalentsArray = append(HeroTalentsArray,av)
    }
}
func loadItems(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ItemsMap = make(map[string]*Items)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Items{}
        p.unmarshalRow(header, row, av)
        ItemsMap[av.CfgID] = av
        ItemsArray = append(ItemsArray,av)
    }
}
func loadLottery(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    LotteryMap = make(map[string]*Lottery)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Lottery{}
        p.unmarshalRow(header, row, av)
        LotteryMap[av.CfgID] = av
        LotteryArray = append(LotteryArray,av)
    }
}
func loadLotteryBox(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    LotteryBoxMap = make(map[string]*LotteryBox)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &LotteryBox{}
        p.unmarshalRow(header, row, av)
        LotteryBoxMap[av.CfgID] = av
        LotteryBoxArray = append(LotteryBoxArray,av)
    }
}
func loadLotteryRankAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    LotteryRankAwardMap = make(map[string]*LotteryRankAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &LotteryRankAward{}
        p.unmarshalRow(header, row, av)
        LotteryRankAwardMap[av.CfgID] = av
        LotteryRankAwardArray = append(LotteryRankAwardArray,av)
    }
}
func loadMailRewardCfg(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    MailRewardCfgMap = make(map[string]*MailRewardCfg)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &MailRewardCfg{}
        p.unmarshalRow(header, row, av)
        MailRewardCfgMap[av.CfgID] = av
        MailRewardCfgArray = append(MailRewardCfgArray,av)
    }
}
func loadOperateItems(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    OperateItemsMap = make(map[string]*OperateItems)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &OperateItems{}
        p.unmarshalRow(header, row, av)
        OperateItemsMap[av.CfgID] = av
        OperateItemsArray = append(OperateItemsArray,av)
    }
}
func loadPlayerSpells(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    PlayerSpellsMap = make(map[string]*PlayerSpells)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &PlayerSpells{}
        p.unmarshalRow(header, row, av)
        PlayerSpellsMap[av.CfgID] = av
        PlayerSpellsArray = append(PlayerSpellsArray,av)
    }
}
func loadPrivilege(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    PrivilegeMap = make(map[string]*Privilege)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Privilege{}
        p.unmarshalRow(header, row, av)
        PrivilegeMap[av.CfgID] = av
        PrivilegeArray = append(PrivilegeArray,av)
    }
}
func loadPvpArena(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    PvpArenaMap = make(map[string]*PvpArena)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &PvpArena{}
        p.unmarshalRow(header, row, av)
        PvpArenaMap[av.CfgID] = av
        PvpArenaArray = append(PvpArenaArray,av)
    }
}
func loadPvpArenaRank(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    PvpArenaRankMap = make(map[string]*PvpArenaRank)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &PvpArenaRank{}
        p.unmarshalRow(header, row, av)
        PvpArenaRankMap[av.CfgID] = av
        PvpArenaRankArray = append(PvpArenaRankArray,av)
    }
}
func loadPvpArenaRankAward(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    PvpArenaRankAwardMap = make(map[string]*PvpArenaRankAward)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &PvpArenaRankAward{}
        p.unmarshalRow(header, row, av)
        PvpArenaRankAwardMap[av.CfgID] = av
        PvpArenaRankAwardArray = append(PvpArenaRankAwardArray,av)
    }
}
func loadQuests(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    QuestsMap = make(map[string]*Quests)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Quests{}
        p.unmarshalRow(header, row, av)
        QuestsMap[av.CfgID] = av
        QuestsArray = append(QuestsArray,av)
    }
}
func loadReborn(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    RebornMap = make(map[string]*Reborn)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Reborn{}
        p.unmarshalRow(header, row, av)
        RebornMap[av.CfgID] = av
        RebornArray = append(RebornArray,av)
    }
}
func loadRechargeConfig(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    RechargeConfigMap = make(map[string]*RechargeConfig)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &RechargeConfig{}
        p.unmarshalRow(header, row, av)
        RechargeConfigMap[av.CfgID] = av
        RechargeConfigArray = append(RechargeConfigArray,av)
    }
}
func loadShareAchieve(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShareAchieveMap = make(map[string]*ShareAchieve)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShareAchieve{}
        p.unmarshalRow(header, row, av)
        ShareAchieveMap[av.CfgID] = av
        ShareAchieveArray = append(ShareAchieveArray,av)
    }
}
func loadShareSupply(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShareSupplyMap = make(map[string]*ShareSupply)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShareSupply{}
        p.unmarshalRow(header, row, av)
        ShareSupplyMap[av.CfgID] = av
        ShareSupplyArray = append(ShareSupplyArray,av)
    }
}
func loadShopBox(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopBoxMap = make(map[string]*ShopBox)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopBox{}
        p.unmarshalRow(header, row, av)
        ShopBoxMap[av.CfgID] = av
        ShopBoxArray = append(ShopBoxArray,av)
    }
}
func loadShopGift(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopGiftMap = make(map[string]*ShopGift)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopGift{}
        p.unmarshalRow(header, row, av)
        ShopGiftMap[av.CfgID] = av
        ShopGiftArray = append(ShopGiftArray,av)
    }
}
func loadShopMagic(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopMagicMap = make(map[string]*ShopMagic)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopMagic{}
        p.unmarshalRow(header, row, av)
        ShopMagicMap[av.CfgID] = av
        ShopMagicArray = append(ShopMagicArray,av)
    }
}
func loadShopRecharge(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopRechargeMap = make(map[string]*ShopRecharge)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopRecharge{}
        p.unmarshalRow(header, row, av)
        ShopRechargeMap[av.CfgID] = av
        ShopRechargeArray = append(ShopRechargeArray,av)
    }
}
func loadShopTechnologyTree(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopTechnologyTreeMap = make(map[string]*ShopTechnologyTree)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopTechnologyTree{}
        p.unmarshalRow(header, row, av)
        ShopTechnologyTreeMap[av.CfgID] = av
        ShopTechnologyTreeArray = append(ShopTechnologyTreeArray,av)
    }
}
func loadShopTechnologyTreeGroup(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    ShopTechnologyTreeGroupMap = make(map[string]*ShopTechnologyTreeGroup)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &ShopTechnologyTreeGroup{}
        p.unmarshalRow(header, row, av)
        ShopTechnologyTreeGroupMap[av.CfgID] = av
        ShopTechnologyTreeGroupArray = append(ShopTechnologyTreeGroupArray,av)
    }
}
func loadSignInFirst(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SignInFirstMap = make(map[string]*SignInFirst)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &SignInFirst{}
        p.unmarshalRow(header, row, av)
        SignInFirstMap[av.CfgID] = av
        SignInFirstArray = append(SignInFirstArray,av)
    }
}
func loadSignInLoop(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SignInLoopMap = make(map[string]*SignInLoop)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &SignInLoop{}
        p.unmarshalRow(header, row, av)
        SignInLoopMap[av.CfgID] = av
        SignInLoopArray = append(SignInLoopArray,av)
    }
}
func loadSpecialAbility(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SpecialAbilityMap = make(map[string]*SpecialAbility)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &SpecialAbility{}
        p.unmarshalRow(header, row, av)
        SpecialAbilityMap[av.CfgID] = av
        SpecialAbilityArray = append(SpecialAbilityArray,av)
    }
}
func loadSpellPresentation(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SpellPresentationMap = make(map[string]*SpellPresentation)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &SpellPresentation{}
        p.unmarshalRow(header, row, av)
        SpellPresentationMap[av.CfgID] = av
        SpellPresentationArray = append(SpellPresentationArray,av)
    }
}
func loadSpells(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SpellsMap = make(map[string]*Spells)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Spells{}
        p.unmarshalRow(header, row, av)
        SpellsMap[av.CfgID] = av
        SpellsArray = append(SpellsArray,av)
    }
}
func loadSystemLock(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    SystemLockMap = make(map[string]*SystemLock)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &SystemLock{}
        p.unmarshalRow(header, row, av)
        SystemLockMap[av.CfgID] = av
        SystemLockArray = append(SystemLockArray,av)
    }
}
func loadVipconfig(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    VipconfigMap = make(map[string]*Vipconfig)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Vipconfig{}
        p.unmarshalRow(header, row, av)
        VipconfigMap[av.CfgID] = av
        VipconfigArray = append(VipconfigArray,av)
    }
}
func loadWelfareCard(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    WelfareCardMap = make(map[string]*WelfareCard)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &WelfareCard{}
        p.unmarshalRow(header, row, av)
        WelfareCardMap[av.CfgID] = av
        WelfareCardArray = append(WelfareCardArray,av)
    }
}
func init(){
    loadFuncMap["AppConfig.csv"] = loadAppConfig
    loadFuncMap["ApplePayProductIDCfg.csv"] = loadApplePayProductIDCfg
    loadFuncMap["CompensationAwardCfg.csv"] = loadCompensationAwardCfg
    loadFuncMap["ContinueRechargeCfg.csv"] = loadContinueRechargeCfg
    loadFuncMap["CumulateConsumeCfg.csv"] = loadCumulateConsumeCfg
    loadFuncMap["CumulateRechargeCfg.csv"] = loadCumulateRechargeCfg
    loadFuncMap["CumulateRechargeForeverCfg.csv"] = loadCumulateRechargeForeverCfg
    loadFuncMap["LotteryRankRewardCfg.csv"] = loadLotteryRankRewardCfg
    loadFuncMap["achieve.csv"] = loadAchieve
    loadFuncMap["achieveBox.csv"] = loadAchieveBox
    loadFuncMap["achieveGameLevel.csv"] = loadAchieveGameLevel
    loadFuncMap["achieveVIP.csv"] = loadAchieveVIP
    loadFuncMap["activityLevelAward.csv"] = loadActivityLevelAward
    loadFuncMap["adbox.csv"] = loadAdbox
    loadFuncMap["announceCondition.csv"] = loadAnnounceCondition
    loadFuncMap["appDungeonActiviesAchieve.csv"] = loadAppDungeonActiviesAchieve
    loadFuncMap["appDungeonActiviesRankAward.csv"] = loadAppDungeonActiviesRankAward
    loadFuncMap["appEquipmentSuitActiviesAchieve.csv"] = loadAppEquipmentSuitActiviesAchieve
    loadFuncMap["appEquipmentSuitActiviesRankAward.csv"] = loadAppEquipmentSuitActiviesRankAward
    loadFuncMap["appNewServiceActiviesAchieve.csv"] = loadAppNewServiceActiviesAchieve
    loadFuncMap["appNewServiceActiviesAchieveShop.csv"] = loadAppNewServiceActiviesAchieveShop
    loadFuncMap["appNewServiceActiviesDailyObj.csv"] = loadAppNewServiceActiviesDailyObj
    loadFuncMap["appNewServiceActiviesRankAward.csv"] = loadAppNewServiceActiviesRankAward
    loadFuncMap["areanRobot.csv"] = loadAreanRobot
    loadFuncMap["awards.csv"] = loadAwards
    loadFuncMap["battleConfig.csv"] = loadBattleConfig
    loadFuncMap["battleNum.csv"] = loadBattleNum
    loadFuncMap["born.csv"] = loadBorn
    loadFuncMap["buffs.csv"] = loadBuffs
    loadFuncMap["creatures.csv"] = loadCreatures
    loadFuncMap["dungeon.csv"] = loadDungeon
    loadFuncMap["dungeonLevel.csv"] = loadDungeonLevel
    loadFuncMap["effects.csv"] = loadEffects
    loadFuncMap["equipmentSuitProperties.csv"] = loadEquipmentSuitProperties
    loadFuncMap["equipmentSuits.csv"] = loadEquipmentSuits
    loadFuncMap["equipments.csv"] = loadEquipments
    loadFuncMap["equipmentsEvolution.csv"] = loadEquipmentsEvolution
    loadFuncMap["eventPoint.csv"] = loadEventPoint
    loadFuncMap["firstRecharge.csv"] = loadFirstRecharge
    loadFuncMap["gameLevels.csv"] = loadGameLevels
    loadFuncMap["global.csv"] = loadGlobal
    loadFuncMap["guide.csv"] = loadGuide
    loadFuncMap["guideStep.csv"] = loadGuideStep
    loadFuncMap["help.csv"] = loadHelp
    loadFuncMap["heroEnhanceLevel.csv"] = loadHeroEnhanceLevel
    loadFuncMap["heroEvolution.csv"] = loadHeroEvolution
    loadFuncMap["heroLevels.csv"] = loadHeroLevels
    loadFuncMap["heroTalents.csv"] = loadHeroTalents
    loadFuncMap["items.csv"] = loadItems
    loadFuncMap["lottery.csv"] = loadLottery
    loadFuncMap["lotteryBox.csv"] = loadLotteryBox
    loadFuncMap["lotteryRankAward.csv"] = loadLotteryRankAward
    loadFuncMap["mailRewardCfg.csv"] = loadMailRewardCfg
    loadFuncMap["operateItems.csv"] = loadOperateItems
    loadFuncMap["playerSpells.csv"] = loadPlayerSpells
    loadFuncMap["privilege.csv"] = loadPrivilege
    loadFuncMap["pvpArena.csv"] = loadPvpArena
    loadFuncMap["pvpArenaRank.csv"] = loadPvpArenaRank
    loadFuncMap["pvpArenaRankAward.csv"] = loadPvpArenaRankAward
    loadFuncMap["quests.csv"] = loadQuests
    loadFuncMap["reborn.csv"] = loadReborn
    loadFuncMap["rechargeConfig.csv"] = loadRechargeConfig
    loadFuncMap["shareAchieve.csv"] = loadShareAchieve
    loadFuncMap["shareSupply.csv"] = loadShareSupply
    loadFuncMap["shopBox.csv"] = loadShopBox
    loadFuncMap["shopGift.csv"] = loadShopGift
    loadFuncMap["shopMagic.csv"] = loadShopMagic
    loadFuncMap["shopRecharge.csv"] = loadShopRecharge
    loadFuncMap["shopTechnologyTree.csv"] = loadShopTechnologyTree
    loadFuncMap["shopTechnologyTreeGroup.csv"] = loadShopTechnologyTreeGroup
    loadFuncMap["signInFirst.csv"] = loadSignInFirst
    loadFuncMap["signInLoop.csv"] = loadSignInLoop
    loadFuncMap["specialAbility.csv"] = loadSpecialAbility
    loadFuncMap["spellPresentation.csv"] = loadSpellPresentation
    loadFuncMap["spells.csv"] = loadSpells
    loadFuncMap["systemLock.csv"] = loadSystemLock
    loadFuncMap["vipconfig.csv"] = loadVipconfig
    loadFuncMap["welfareCard.csv"] = loadWelfareCard
}
