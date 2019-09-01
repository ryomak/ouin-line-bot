package line

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/ryomak/login-bonus-manager/line-bot/src/repository"
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

func MakeMessge(id, str string) string {
	if strings.Index(str, "一覧表示") != -1 {
		homelist := repository.GetHomeList(id)
		res := ""
		for _, v := range homelist {
			res += v.Value + "\n"
		}
		return res
	}
	log.Println("ぉっl")
	repMessage := []string{
		"できたのはすごいね",
		"はやばすぎる",
		"もうオンリーワンだね",
	}
	repository.SetHome(&repository.Home{ID: id, Value: str})
	return "「" + str + "」" + repMessage[int(time.Now().UnixNano())%len(repMessage)]
}
