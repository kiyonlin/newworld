package main

import (
	"log"
	"net/http"

	"github.com/kiyonlin/newworld/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("无法监听 5000 端口 %v", err)
	}
}
