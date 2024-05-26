package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lakshhtaneja/cryptoprice/models"
)

const apiUrl = "https://cex.io/api/ticker/%s/INR"

func GetRate(currency string) (*models.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	var cryptoRate Response
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 200 { //http.StatusOK
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &cryptoRate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code recieved: %v", res.StatusCode)
	}

	rate := &models.Rate{Currency: currency, Price: cryptoRate.Bid}
	return rate, nil
}
