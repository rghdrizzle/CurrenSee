package currensee

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"reflect" // for iterating in the struct

)
type data2 struct {
	AvailableCryptos struct {
		Ada   string `json:"ADA"`
		Atom  string `json:"ATOM"`
		Avax  string `json:"AVAX"`
		Axs   string `json:"AXS"`
		Bch   string `json:"BCH"`
		Bnb   string `json:"BNB"`
		Btc   string `json:"BTC"`
		Btg   string `json:"BTG"`
		Busd  string `json:"BUSD"`
		Dai   string `json:"DAI"`
		Dash  string `json:"DASH"`
		Doge  string `json:"DOGE"`
		Dot   string `json:"DOT"`
		Egld  string `json:"EGLD"`
		Enj   string `json:"ENJ"`
		Eos   string `json:"EOS"`
		Etc   string `json:"ETC"`
		Eth   string `json:"ETH"`
		Fil   string `json:"FIL"`
		Flow  string `json:"FLOW"`
		Ftm   string `json:"FTM"`
		Ftt   string `json:"FTT"`
		Gala  string `json:"GALA"`
		Hbar  string `json:"HBAR"`
		Hnt   string `json:"HNT"`
		Icp   string `json:"ICP"`
		Link  string `json:"LINK"`
		Lrc   string `json:"LRC"`
		Ltc   string `json:"LTC"`
		Luna  string `json:"LUNA"`
		Mana  string `json:"MANA"`
		Matic string `json:"MATIC"`
		Near  string `json:"NEAR"`
		Neo   string `json:"NEO"`
		Rose  string `json:"ROSE"`
		Sand  string `json:"SAND"`
		Shib  string `json:"SHIB"`
		Sol   string `json:"SOL"`
		Theta string `json:"THETA"`
		Trx   string `json:"TRX"`
		Uni   string `json:"UNI"`
		Usdt  string `json:"USDT"`
		Ust   string `json:"UST"`
		Vet   string `json:"VET"`
		Xlm   string `json:"XLM"`
		Xmr   string `json:"XMR"`
		Xrp   string `json:"XRP"`
		Xtz   string `json:"XTZ"`
	} `json:"available_currencies"`
	Endpoint string `json:"endpoint"`
}
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

func Getcrypto(){
	apikey:= "N03byn9cUhvo7ay5szSx"
	url := "https://marketdata.tradermade.com/api/v1/live_crypto_list?api_key="+apikey
	req, err:= http.Get(url)
	if err!=nil{
		fmt.Println("error",err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
  	if readErr != nil {
      log.Fatal(readErr)
    }
	data_obj := data2{}    
	jsonErr := json.Unmarshal(body, &data_obj)
      if jsonErr != nil {
         log.Fatal(jsonErr)
	  }

	v := reflect.ValueOf(data_obj.AvailableCryptos)
	typeOfd := v.Type()
	for i :=0 ; i<v.NumField();i++{
		fmt.Println(typeOfd.Field(i).Name, v.Field(i).Interface())
	}

}