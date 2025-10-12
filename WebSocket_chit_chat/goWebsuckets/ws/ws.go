package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func Hi() string {
	return "Hey there"
}

type responseBody struct {
	Token string `json:"token"`
}

func Connect(token string) (context.Context, context.CancelFunc, *websocket.Conn) {
	url := "wss://hackattic.com/_/ws/" + token

	// connect to ws server
	// normaly i should handle the cancel func here but oops
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	conn, resp, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Body)
	return ctx, cancel, conn
}

func Handle(ctx *context.Context, conn *websocket.Conn) {
	log.Println("Started Reading from Websocket")
	last := time.Now().UnixMilli()
	for {
		msgT, buf, err := conn.Read(*ctx)
		log.Println(msgT)
		log.Println(string(buf))
		now := time.Now().UnixMilli()
		if err != nil {
			log.Fatal(err)
		}
		switch string(buf[:5]) {
		case "ping!":
			handlePing(*ctx, conn, now-last)
			last = now
		case "good!":
			fmt.Println(string(buf))
		}
		fmt.Println(string(buf))

	}
}

func handlePing(ctx context.Context, conn *websocket.Conn, timeSinceLast int64) error {
	fmt.Println("got: " + fmt.Sprint(timeSinceLast))
	err := conn.Write(ctx, websocket.MessageText, fmt.Append(nil, getClosestTime(timeSinceLast)))
	return err
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func getClosestTime(tsl int64) int64 {
	times := []int64{700, 1500, 2000, 2500, 3000}
	var closest int64
	for _, t := range times {
		if Abs(tsl-t) < Abs(tsl-closest) {
			closest = t
		}
	}

	return closest
}

func GetToken() string {
	log.Println("Getting challenge token")
	response, err := http.Get("https://hackattic.com/challenges/websocket_chit_chat/problem?access_token=be8466c39c975877")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rb responseBody
	json.Unmarshal(rawBody, &rb)

	log.Println(rb.Token)
	return rb.Token
}
