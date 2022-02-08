package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getCommodityOverview() string {
	rt := ""

	resp, err := http.Get("https://api-pub.bitfinex.com/v2/tickers?symbols=tBTCUSD")
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var btcData [][]interface{}
	if err := json.Unmarshal(body, &btcData); err != nil {
		return err.Error()
	}
	rt += fmt.Sprintf("BTCUSD.bitfinex: %f\n", btcData[0][7].(float64))

	resp, err = http.Get("https://www.freeforexapi.com/api/live?pairs=USDXAU")
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var xauData struct {
		Rates map[string]struct {
			Rate float64
			// Timestamp uint64
		}
		// Code uint64
	}
	if err := json.Unmarshal(body, &xauData); err != nil {
		return err.Error()
	}
	rt += fmt.Sprintf("XAUUSD.freeforexapi: %f\n", 1.0/xauData.Rates["USDXAU"].Rate)

	return rt
}
