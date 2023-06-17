package currensee 

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
)

type res struct {
	BaseCurrency  string  `json:"base_currency"`
	Quote         float64 `json:"quote"`
	QuoteCurrency string  `json:"quote_currency"`
	RequestedTime string  `json:"requested_time"`
	Timestamp     int     `json:"timestamp"`
	Total         float64    `json:"total"`
}

func Convert(amount string,from string, to string){
	apikey:="N03byn9cUhvo7ay5szSx"
	url:="https://marketdata.tradermade.com/api/v1/convert?api_key="+apikey+"&from="+from+"&to="+to+"&amount="+amount
	req, err:= http.Get(url)
	if err!=nil{
		fmt.Println("error",err)
	}
	body, readErr := ioutil.ReadAll(req.Body)
  	if readErr != nil {
      log.Fatal(readErr)
    }
	data_obj := res{}    
	jsonErr := json.Unmarshal(body, &data_obj)
      if jsonErr != nil {
         log.Fatal(jsonErr)
	  }
	fmt.Println("Base currency:",data_obj.BaseCurrency)
	fmt.Println("Amount:",amount)
	fmt.Println("To:",data_obj.QuoteCurrency)
	fmt.Println("The amount after converting is:",data_obj.Total)
}