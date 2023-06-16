package currensee

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"reflect" // for iterating in the struct

)
type data struct {
	AvailableCurrencies struct {
		Aed string `json:"AED"`
		All string `json:"ALL"`
		Aoa string `json:"AOA"`
		Ars string `json:"ARS"`
		Aud string `json:"AUD"`
		Bdt string `json:"BDT"`
		Bgn string `json:"BGN"`
		Bhd string `json:"BHD"`
		Brl string `json:"BRL"`
		Cad string `json:"CAD"`
		Chf string `json:"CHF"`
		Clp string `json:"CLP"`
		Cnh string `json:"CNH"`
		Cny string `json:"CNY"`
		Cop string `json:"COP"`
		Czk string `json:"CZK"`
		Dkk string `json:"DKK"`
		Egp string `json:"EGP"`
		Eur string `json:"EUR"`
		Gbp string `json:"GBP"`
		Ghs string `json:"GHS"`
		Hkd string `json:"HKD"`
		Hrk string `json:"HRK"`
		Huf string `json:"HUF"`
		Idr string `json:"IDR"`
		Ils string `json:"ILS"`
		Inr string `json:"INR"`
		Isk string `json:"ISK"`
		Jod string `json:"JOD"`
		Jpy string `json:"JPY"`
		Kes string `json:"KES"`
		Krw string `json:"KRW"`
		Kwd string `json:"KWD"`
		Lbp string `json:"LBP"`
		Lkr string `json:"LKR"`
		Mad string `json:"MAD"`
		Mur string `json:"MUR"`
		Mxn string `json:"MXN"`
		Myr string `json:"MYR"`
		Ngn string `json:"NGN"`
		Nok string `json:"NOK"`
		Nzd string `json:"NZD"`
		Omr string `json:"OMR"`
		Pen string `json:"PEN"`
		Php string `json:"PHP"`
		Pkr string `json:"PKR"`
		Pln string `json:"PLN"`
		Qar string `json:"QAR"`
		Ron string `json:"RON"`
		Rub string `json:"RUB"`
		Sar string `json:"SAR"`
		Sek string `json:"SEK"`
		Sgd string `json:"SGD"`
		Thb string `json:"THB"`
		Tnd string `json:"TND"`
		Try string `json:"TRY"`
		Twd string `json:"TWD"`
		Usd string `json:"USD"`
		Vnd string `json:"VND"`
		Xaf string `json:"XAF"`
		Xag string `json:"XAG"`
		Xau string `json:"XAU"`
		Xof string `json:"XOF"`
		Xpd string `json:"XPD"`
		Xpt string `json:"XPT"`
		Zar string `json:"ZAR"`
		Zwl string `json:"ZWL"`
	} `json:"available_currencies"`
	Endpoint string `json:"endpoint"`
}

func Getlist(){
	apikey:= "N03byn9cUhvo7ay5szSx"
	url := "https://marketdata.tradermade.com/api/v1/live_currencies_list?api_key="+apikey
	req, err:= http.Get(url)
	if err!=nil{
		fmt.Println("error",err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
  	if readErr != nil {
      log.Fatal(readErr)
    }
	data_obj := data{}    
	jsonErr := json.Unmarshal(body, &data_obj)
      if jsonErr != nil {
         log.Fatal(jsonErr)
	  }

	v := reflect.ValueOf(data_obj.AvailableCurrencies)
	typeOfd := v.Type()
	for i :=0 ; i<v.NumField();i++{
		fmt.Println(typeOfd.Field(i).Name, v.Field(i).Interface())
	}
	}
