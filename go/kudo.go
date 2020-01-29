package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Data0 struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    []Data1 `json:"data"`
}

type Data1 struct {
	ProductType       string  `json:"product_type"`
	ProductTypeID     int     `json:"product_type_id"`
	ProductCategories []Data2 `json:"product_categories"`
}

type Data2 struct {
	ProductCategory   string  `json:"product_category"`
	ProductCategoryID int     `json:"product_category_id"`
	ProductGroups     []Data3 `json:"product_groups"`
}

type Data3 struct {
	ProductGroup       string   `json:"product_group"`
	ProductGroupID     int      `json:"product_group_id"`
	ProductGroupType   string   `json:"product_group_type"`
	ProductGroupTypeID int      `json:"product_group_type_id"`
	ProductPrefix      []string `json:"product_prefix"`
	ProductImage       string   `json:"product_image"`
	ProductList        []Data4  `json:"product_list"`
}

type Data4 struct {
	CountryOrigin           string `json:"country_origin"`
	ProductCode             string `json:"product_code"`
	ProductName             string `json:"product_name"`
	ProductDescription      string `json:"product_description"`
	ProductNominal          int    `json:"product_nominal"`
	ProductNominalCurrency  string `json:"product_nominal_currency"`
	ProductB2BPrice         int    `json:"product_b2b_price"`
	ProductB2BPriceCurrency string `json:"product_b2b_price_currency"`
	ProductPrice            int    `json:"product_price"`
	ProductPriceCurrency    string `json:"product_price_currency"`
	ProductCurrencySymbol   string `json:"product_currency_symbol"`
	ProductStatus           string `json:"product_status"`
}

type KudoData struct {
	ProductCode   string   `json:"product_code"`
	ProductName   string   `json:"product_name"`
	ProductPrice  int      `json:"product_price"`
	ProductPrefix []string `json:"product_prefix"`
	ProductImage  string   `json:"product_image"`
}

type KudoFilter struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func inArray(val interface{}, array interface{}) (exists bool) {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

func main() {

	url := "https://testapi.io/api/holmes4869/kudo_products"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people0 := Data0{}
	jsonErr := json.Unmarshal(body, &people0)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//result := make([]*KudoData, 0)
	result := []KudoData{}
	//var productName string
	productName := ""

	for _, a := range people0.Data {
		for _, b := range a.ProductCategories {
			if b.ProductCategoryID == 7 || b.ProductCategoryID == 8 {
				for _, c := range b.ProductGroups {
					for _, d := range c.ProductList {
						fmt.Println(c.ProductPrefix)
						productName += c.ProductGroup + " " + d.ProductName
						n := KudoData{ProductName: productName, ProductCode: d.ProductCode, ProductPrice: d.ProductB2BPrice, ProductPrefix: c.ProductPrefix, ProductImage: c.ProductImage}
						result = append(result, n)
						productName = ""
					}
				}
			}
		}
	}

	phone := "+628123456789"
	phone = phone[3:len(phone)]
	phone = "0" + phone
	prefixPhone := phone[0:4]

	resp := []KudoFilter{}

	limit := 100000

	for _, v := range result {

		if inArray(prefixPhone, v.ProductPrefix) && limit >= v.ProductPrice {
			//fmt.Println(v.ProductPrice)
			s := KudoFilter{Code: v.ProductCode, Name: v.ProductName, Price: v.ProductPrice, Image: v.ProductImage}
			resp = append(resp, s)
		}
	}

	//telkomselPrefix:=

	//fmt.Println(resp)
	//fmt.Println(result)
	//fmt.Println(reflect.TypeOf(result))
}
