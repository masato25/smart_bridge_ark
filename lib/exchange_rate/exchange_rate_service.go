package exchange_rate

import (
	"fmt"

	"github.com/emirpasic/gods/maps"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

const apiUrl = "https://api.coinmarketcap.com/v1/ticker"

var ReqClient *gorequest.SuperAgent
var SupportCurrency maps.Map

func Init() {
	ReqClient = gorequest.New()
	SupportCurrency = hashmap.New()
	SupportCurrency.Put("bitcoin", float64(0))
	SupportCurrency.Put("ethereum", float64(0))
	SupportCurrency.Put("ark", float64(0))
	log.Info("init exchange_rate_service")
}

type Raging struct {
	FromUSD float64
	ToUSD   float64
}

func (this *Raging) Rate(fromAmount float32) float64 {
	return this.FromUSD / this.ToUSD
}

func parseCurrencyRate(body gjson.Result) (rate float64, err error) {
	jsonArray := body.Array()
	if len(jsonArray) == 0 {
		err = fmt.Errorf("no data found")
		return
	} else {
		usdRate := body.Get("0.price_usd")
		if !usdRate.Exists() {
			err = fmt.Errorf("no data found")
			return
		}
		rate = usdRate.Float()

	}
	return
}

type CoinObj struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetAllRate() []CoinObj {
	var coinobj []CoinObj
	for _, k := range SupportCurrency.Keys() {
		reqUrl := fmt.Sprintf("%s/%s", apiUrl, k)
		log.Debug(reqUrl)
		_, body, _ := ReqClient.Get(reqUrl).End()
		jbody := gjson.Parse(body)
		rate, _ := parseCurrencyRate(jbody)
		coinobj = append(coinobj, CoinObj{k.(string), rate})
	}
	return coinobj
}

func GetRate(fromCurrencyCode string, toCurrencyCode string) (rate Raging, err error) {
	if _, ok := SupportCurrency.Get(fromCurrencyCode); ok {
		reqUrl := fmt.Sprintf("%s/%s", apiUrl, fromCurrencyCode)
		log.Debug(reqUrl)
		_, body, errs := ReqClient.Get(reqUrl).End()
		if len(errs) != 0 {
			err = errs[0]
			return
		}
		jbody := gjson.Parse(body)
		var fromusd float64
		fromusd, err = parseCurrencyRate(jbody)
		rate.FromUSD = fromusd
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("%s is not supported.", fromCurrencyCode)
		return
	}

	if _, ok := SupportCurrency.Get(toCurrencyCode); ok {
		reqUrl := fmt.Sprintf("%s/%s", apiUrl, toCurrencyCode)
		log.Debug(reqUrl)
		_, body, errs := ReqClient.Get(reqUrl).End()
		if len(errs) != 0 {
			err = errs[0]
			return
		}
		jbody := gjson.Parse(body)
		var tousd float64
		tousd, err = parseCurrencyRate(jbody)
		rate.ToUSD = tousd
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("%s is not supported.", toCurrencyCode)
		return
	}
	return
}
