package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data01 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data02 `json:"data"`
}

type Data02 struct {
	ProductGroup       string   `json:"product_group"`
	ProductGroupID     int      `json:"product_group_id"`
	ProductGroupType   string   `json:"product_group_type"`
	ProductGroupTypeID int      `json:"product_group_type_id"`
	ProductPrefix      []string `json:"product_prefix"`
	ProductImage       string   `json:"product_image"`
	ProductList        []Data03 `json:"product_list"`
}

type Data03 struct {
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
	ProductCode  string `json:"product_code"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	ProductImage string `json:"product_image"`
}

type KudoFilter struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func main() {

	url := "https://testapi.io/api/holmes4869/kudo_filter/430"

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

	data01 := Data01{}
	jsonErr := json.Unmarshal(body, &data01)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	result := []KudoData{}
	productName := ""

	for _, a := range data01.Data.ProductList {
		productName += data01.Data.ProductGroup + " " + a.ProductName
		productPrice := a.ProductB2BPrice
		s := KudoData{ProductName: productName, ProductCode: a.ProductCode, ProductPrice: productPrice, ProductImage: data01.Data.ProductImage}
		result = append(result, s)
	}

	fmt.Println(result)

}
