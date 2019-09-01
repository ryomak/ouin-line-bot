package repository

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Ouin struct {
	Heading string `json:"heading"`
	Text    string `json:"text"`
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
}

func GetOuinList(tango string) []Ouin {
  v := url.Values{}
  v.Add("api","1")
  v.Add("dict","広辞苑")
  v.Add("q",tango)
	url := "https://sakura-paris.org/dict/?" + v.Encode()
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return []Ouin{}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Ouin{}
	}
	ouinList := make([]Ouin, 0)
	err = json.Unmarshal(body, &ouinList)
	if err != nil {
		return []Ouin{}
	}
	defer resp.Body.Close()
	return ouinList
}
