package line

import (
	"encoding/json"
	"strings"

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
	tokens := repository.GetToken(str)
	out := ""
	for _, v := range tokens.MaResult.WordList.Word {
		if strings.Index(v.Pos, "助") != -1 {
			out += v.Surface
		} else {
			out += trim(repository.GetOuinList(v.Reading)[0].Heading)
		}
	}
	return out
}

func trim(str string) string {
	first := 0
	runeStr := []rune(str)
	for i, c := range runeStr {
		if c == '【' {
			first = i
		}
	}
	return string(runeStr[first+1 : len(runeStr)-1])
}
