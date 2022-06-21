package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getCurrencyOverview() string {
	rt := ""
	list := []string{"GBPUSD", "EURUSD", "USDCAD", "USDCHF", "USDJPY", "AUDUSD", "USDCNY"}
	pairs := strings.Join(list, ",")

	resp, err := http.Get("https://www.freeforexapi.com/api/live?pairs=" + pairs)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Rates map[string]struct {
			Rate float64
			// Timestamp uint64
		}
		// Code uint64
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return err.Error()
	}
	for pair, value := range data.Rates {
		rt += fmt.Sprintf("%s.freeforexapi: %f\n", pair, value.Rate)
	}
	return rt
}
