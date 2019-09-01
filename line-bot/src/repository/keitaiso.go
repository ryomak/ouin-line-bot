package repository

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Token struct {
	XMLName        xml.Name `xml:"ResultSet"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	MaResult       struct {
		Text          string `xml:",chardata"`
		TotalCount    string `xml:"total_count"`
		FilteredCount string `xml:"filtered_count"`
		WordList      struct {
			Text string `xml:",chardata"`
			Word []struct {
				Text     string `xml:",chardata"`
				Surface  string `xml:"surface"`
				Reading  string `xml:"reading"`
				Pos      string `xml:"pos"`
				Baseform string `xml:"baseform"`
			} `xml:"word"`
		} `xml:"word_list"`
	} `xml:"ma_result"`
	UniqResult struct {
		Text          string `xml:",chardata"`
		TotalCount    string `xml:"total_count"`
		FilteredCount string `xml:"filtered_count"`
		WordList      struct {
			Text string `xml:",chardata"`
			Word []struct {
				Text     string `xml:",chardata"`
				Surface  string `xml:"surface"`
				Reading  string `xml:"reading"`
				Pos      string `xml:"pos"`
				Baseform string `xml:"baseform"`
				Count    string `xml:"count"`
			} `xml:"word"`
		} `xml:"word_list"`
	} `xml:"uniq_result"`
}

func GetToken(sent string) Token {
	v := url.Values{}
	v.Add("appid", os.Getenv("yahooAppId"))
	v.Add("results", "ma")
	v.Add("sentence", sent)
	url := "https://jlp.yahooapis.jp/MAService/V1/parse?" + v.Encode()
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return Token{}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Token{}
	}
	token := new(Token)
	err = xml.Unmarshal(body, &token)
	if err != nil {
		return Token{}
	}
	defer resp.Body.Close()
	return *token
}
