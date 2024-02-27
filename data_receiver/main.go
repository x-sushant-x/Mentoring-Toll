package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sushant102004/TollCalculator/types"
	"log"
	"net/http"
)

type DataReceiver struct {
	webSCon *websocket.Conn
}

func NewDataReciever() DataReceiver {
	return DataReceiver{}
}

func main() {
	dataReceiver := NewDataReciever()
	http.HandleFunc("/ws", dataReceiver.wsHandler)

	go func() {
		if err := http.ListenAndServe(":6443", nil); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Waiting for data from truck...")
	select {}
}

func (dr *DataReceiver) wsHandler(w http.ResponseWriter, req *http.Request) {
	u := websocket.Upgrader{ReadBufferSize: 1028, WriteBufferSize: 1028}

	conn, err := u.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.webSCon = conn

	go dr.wsLoop()
}

func (dr *DataReceiver) wsLoop() {
	fmt.Println("Client Connected")
	for {
		var data types.OBUData
		err := dr.webSCon.ReadJSON(&data)
		if err != nil {
			fmt.Println("error: ", err.Error())
			continue
		}
	}
}
