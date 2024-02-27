/*
	PoF -> This file will act as a virtual truck that will send data to our server every 30 seconds.
		   This data is basically the current location of truck.
*/

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sushant102004/TollCalculator/types"
	"log"
	"math/rand"
	"time"
)

const wsEndpoint = "ws://localhost:6443/ws"

func main() {
	// This is truck ID which is sending data.
	obuID := "HR-PB-6990"

	con, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	for {
		lat, lon := generateTruckLocation()

		obuData := types.OBUData{
			ObuID: obuID,
			Lat:   lat,
			Long:  lon,
		}

		err := con.WriteJSON(obuData)
		if err != nil {
			fmt.Println("error: unable to send data to websocket: ", err.Error())
		}

		fmt.Println("Data Sent: ", obuData)

		time.Sleep(time.Second * 5)
	}
}

// Generate Random Coordinate: - 39.65477667
func generateCoord() float64 {
	n := float64(rand.Intn(100)) + 1
	f := rand.Float64()
	return n + f
}

func generateTruckLocation() (float64, float64) {
	// latitude, longitude
	return generateCoord(), generateCoord()
}
