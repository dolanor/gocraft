package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dolanor/rip"
	"github.com/dolanor/rip/encoding/html"
	"github.com/dolanor/rip/encoding/json"
)

func startRIP() {
	// We need to wait for the game to start to have an initalized `game` pointer
	time.Sleep(1 * time.Second)

	bp := BlockProvider{
		game: game,
	}

	ro := rip.NewRouteOptions().
		WithCodecs(
			json.Codec,
			html.Codec,
			html.FormCodec,
		)

	http.HandleFunc(rip.HandleEntities("/blocks/", &bp, ro))
	log.Println("running RIP")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
