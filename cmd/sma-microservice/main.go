package main

import "sma-microservice/internal/api"

func main() {
	api := api.New()
	api.Start()
	// req, _ := sma.CreateSMARequest("coinbasepro", "BTC-USD", "1m")
	// indicator, err := sma.GetIndicator(req)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(indicator)
}
