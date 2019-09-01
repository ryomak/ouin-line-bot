package line

import (
	"encoding/json"
	"strings"

	"github.com/ryomak/ouin-line-bot/line-bot/src/ouin"
	"github.com/ryomak/ouin-line-bot/line-bot/src/repository"
)

func UnmarshalLineRequest(data []byte) (LineRequest, error) {
	var r LineRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LineRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LineRequest struct {
	Events      []Event `json:"events"`
	Destination string  `json:"destination"`
}

type Event struct {
	Type       string  `json:"type"`
	ReplyToken string  `json:"replyToken"`
	Source     Source  `json:"source"`
	Timestamp  int64   `json:"timestamp"`
	Message    Message `json:"message"`
}

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Source struct {
	UserID string `json:"userId"`
	Type   string `json:"type"`
}

func MakeMessge(str string) string {
	tokens := ouin.GetHiraganas(str)
	out := ""
	for _, v := range tokens {
		if strings.Index(v.Type, "助") != -1 {
			out += v.Surface
		} else {
			out += trim(repository.GetOuinList(v.Hiragana)[0].Heading)
		}
	}
  return out
}

func trim(str string) string {
	first := 0
	for i, c := range str {
		if c == '【' {
			first = i
		}
	}
	return str[first+1 : len(str)-1]
}
