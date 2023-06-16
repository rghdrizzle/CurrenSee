package currensee

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"

)

type rate struct {
	Endpoint       string                   `json:'endpoint'`
	Quotes         []map[string]interface{} `json:'quotes'`
	Requested_time string                   `json:'requested_time'`
	Timestamp      int32                    `json:'timestamp'`
}

func ExchangeRate(input string){
	currencies:=input
	apikey:= "N03byn9cUhvo7ay5szSx"
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + apikey
	req, err:= http.Get(url)
	if err!=nil{
		fmt.Println("error",err)
	}
	body, readErr := ioutil.ReadAll(req.Body)
  	if readErr != nil {
      log.Fatal(readErr)
    }
	data_obj := rate{}    
	jsonErr := json.Unmarshal(body, &data_obj)
      if jsonErr != nil {
         log.Fatal(jsonErr)
	  }
	  fmt.Println("requested time", data_obj.Requested_time, "timestamp", data_obj.Timestamp)
	  for key, value := range data_obj.Quotes {
			  fmt.Println(key)
			  fmt.Println("symbol", value["base_currency"] ,value["quote_currency"], "bid", value["bid"], "ask", value["ask"],"mid", value["mid"])


}
}