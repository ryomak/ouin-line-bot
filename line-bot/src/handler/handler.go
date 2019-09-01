package handler

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	myLineRequest, err := line.UnmarshalLineRequest([]byte(request.Body))
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(
		os.Getenv("channelSecret"),
		os.Getenv("channelToken"),
	)
	if err != nil {
		log.Fatal(err)
	}
	message := line.MakeMessge(myLineRequest.Events[0].Message.ID, myLineRequest.Events[0].Message.Text)
  log.Println(message)
	if _, err = bot.ReplyMessage(myLineRequest.Events[0].ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Fatal(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}
