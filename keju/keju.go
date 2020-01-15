package keju

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HMGetByKeyword(keyword string) (QAPairCollection, error) {
	const queryURLPattern = "https://haimanchajian.com/api/jx3/search-data/keju?q=%s&offset=0&limit=100"
	url := fmt.Sprintf(queryURLPattern, keyword)
	resp, err := http.Get(url)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Println("Close Response Body Fail: ", err.Error())
			}
		}()
	}
	if err != nil {
		return QAPairCollection{}, errors.New("Request Failed: " + err.Error())
	}
	var parsedResp struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Data    []struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		} `json:"data"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&parsedResp); err != nil {
		return QAPairCollection{}, errors.New("JSON Decode For Response Failed: " + err.Error())
	}
	if parsedResp.Errcode != 0 {
		return QAPairCollection{}, errors.New("Response With Error " + strconv.Itoa(parsedResp.Errcode) + " : " + parsedResp.Errmsg)
	}
	col := QAPairCollection{}
	for _, v := range parsedResp.Data {
		col.Insert(QAPair(v))
	}
	return col, nil
}

func MYGetByKeyword(keyword string) (QAPairCollection, error) {
	const queryURLPattern = "https://j3cx.com/exam/?m=1&q=%s&csrf="
	url := fmt.Sprintf(queryURLPattern, keyword)
	resp, err := http.Get(url)
	if resp != nil {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Println("Close Response Body Fail: ", err.Error())
			}
		}()
	}
	if err != nil {
		return QAPairCollection{}, errors.New("Request Failed: " + err.Error())
	}
	var parsedResp struct {
		Code    string    `json:"code"`
		Keyword string `json:"question"`
		Data    []struct {
			Question string `json:"ques"`
			Answer   string `json:"answ"`
		} `json:"result"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&parsedResp); err != nil {
		return QAPairCollection{}, errors.New("JSON Decode For Response Failed: " + err.Error())
	}
	if parsedResp.Code != "200" {
		return QAPairCollection{}, errors.New("Response With Error " + parsedResp.Code)
	}
	col := QAPairCollection{}
	for _, v := range parsedResp.Data {
		col.Insert(QAPair(v))
	}
	return col, nil
}
