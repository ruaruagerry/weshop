package gamecfg

import (
    "encoding/csv"
    "bufio"
    "github.com/sirupsen/logrus"
)
var (
    // Catalog_categoryMap TODO
    Catalog_categoryMap map[string]*Catalog_category
    Catalog_categoryArray []*Catalog_category
    // Catalog_navMap TODO
    Catalog_navMap map[string]*Catalog_nav
    Catalog_navArray []*Catalog_nav
    // GlobalMap TODO
    GlobalMap map[string]*Global
    GlobalArray []*Global
    // GoodsMap TODO
    GoodsMap map[string]*Goods
    GoodsArray []*Goods
    // Goods_specificationMap TODO
    Goods_specificationMap map[string]*Goods_specification
    Goods_specificationArray []*Goods_specification
    // TestMap TODO
    TestMap map[string]*Test
    TestArray []*Test
)
func loadCatalog_category(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    Catalog_categoryMap = make(map[string]*Catalog_category)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Catalog_category{}
        p.unmarshalRow(header, row, av)
        Catalog_categoryMap[av.CfgID] = av
        Catalog_categoryArray = append(Catalog_categoryArray,av)
    }
}
func loadCatalog_nav(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    Catalog_navMap = make(map[string]*Catalog_nav)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Catalog_nav{}
        p.unmarshalRow(header, row, av)
        Catalog_navMap[av.CfgID] = av
        Catalog_navArray = append(Catalog_navArray,av)
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
func loadGoods(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    GoodsMap = make(map[string]*Goods)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Goods{}
        p.unmarshalRow(header, row, av)
        GoodsMap[av.CfgID] = av
        GoodsArray = append(GoodsArray,av)
    }
}
func loadGoods_specification(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    Goods_specificationMap = make(map[string]*Goods_specification)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Goods_specification{}
        p.unmarshalRow(header, row, av)
        Goods_specificationMap[av.CfgID] = av
        Goods_specificationArray = append(Goods_specificationArray,av)
    }
}
func loadTest(file *bufio.Reader, p *parser) {
    csvReader := csv.NewReader(file)
    csvReader.Read()
    csvReader.Read()
    header, err := csvReader.Read()
    if err != nil {
        logrus.Panic(err)
    }
    TestMap = make(map[string]*Test)
    for {
        row, err := csvReader.Read()
        if err != nil {
            break
        }

        av := &Test{}
        p.unmarshalRow(header, row, av)
        TestMap[av.CfgID] = av
        TestArray = append(TestArray,av)
    }
}
func init(){
    loadFuncMap["catalog_category.csv"] = loadCatalog_category
    loadFuncMap["catalog_nav.csv"] = loadCatalog_nav
    loadFuncMap["global.csv"] = loadGlobal
    loadFuncMap["goods.csv"] = loadGoods
    loadFuncMap["goods_specification.csv"] = loadGoods_specification
    loadFuncMap["test.csv"] = loadTest
}
