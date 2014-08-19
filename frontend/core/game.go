package game

import (
	"encoding/json"
	"frontend/structs"
	"net/http"
	"time"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var result []*structs.Game

	var bots []string
	bots[0] = "123"
	bots[1] = "1234"
	bots[2] = "12345"

	var notes []string
	notes[0] = "123"
	notes[1] = "1234"
	notes[2] = "12345"

	var res map[string]float64
	res["123"] = 0.5
	res["1234"] = 0.5
	res["12345"] = 0.0

	var game1 = &structs.Game{
		GID:       "123",
		START:     time.Now(),
		FINISH:    time.Now(),
		BOTID:     bots,
		RESULTMAP: res,
		LISTING:   "bla-bla-bla",
		NOTES:     notes}

	result[0] = game1

	b, _ := json.Marshal(result)
	w.Write(b)
}
