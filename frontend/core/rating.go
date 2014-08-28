package core

import (
	"encoding/json"
	"frontend/structs"
	"net/http"
)

func GetRating(w http.ResponseWriter) {
	var result []*structs.BotRich
	result = make([]*structs.BotRich, 1)

	var bot1 = &structs.BotRich{}
	bot1.BID = "123"

	result[0] = bot1

	b, _ := json.Marshal(result)
	w.Write(b)
}
