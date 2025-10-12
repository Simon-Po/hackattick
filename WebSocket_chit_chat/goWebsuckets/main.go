package main

import (
	"log"

	"hackattic/WebsocketsGo/ws"
)

func main() {
	log.Println(ws.Hi())
	ctx, cancelCtx, conn := ws.Connect(ws.GetToken())
	defer cancelCtx()
	ws.Handle(&ctx, conn)
}
