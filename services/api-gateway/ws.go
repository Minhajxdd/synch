package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Minhajxdd/Synch/shared/contracts"
	"github.com/Minhajxdd/Synch/shared/util"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRidersWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed : %v", err)
		return
	}

	defer conn.Close()

	userId := r.URL.Query().Get("userID")
	if userId == "" {
		log.Println("User id is required")
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message : %v", err)
			break
		}

		fmt.Printf("Received Message of is %s", p)
	}
}

func handleDriversWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	userId := r.URL.Query().Get("userID")
	if userId == "" {
		log.Println("User id is required")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		log.Println("packageSlug is required")
		return
	}

	type Driver struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		ProfilePicture string `json:"profilePicture"`
		CarPlate       string `json:"carPlate"`
		PackageSlug    string `json:"packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver{
			Id:             userId,
			Name:           "minhaj",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate:       "KA 43",
			PackageSlug:    packageSlug,
		},
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("Error sending message : %v", err)
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message : %v", err)
			break
		}

		fmt.Printf("Received Message of  %s", p)
	}
}
