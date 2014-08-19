package core

import (
	"encoding/json"
	"frontend/structs"
	"net/http"
	"time"
)

func GetGames(w http.ResponseWriter) {
	var result []*structs.Game
	result = make([]*structs.Game, 3)

	var bots []string
	bots = make([]string, 3)
	bots[0] = "123"
	bots[1] = "1234"
	bots[2] = "12345"

	var notes []string
	notes = make([]string, 3)
	notes[0] = "123"
	notes[1] = "1234"
	notes[2] = "12345"

	var res map[string]float64
	res = make(map[string]float64)
	res["123"] = 0.5
	res["1234"] = 0.5
	res["12345"] = 0.0

	var game1 = &structs.Game{
		GID:       "id1",
		START:     time.Now(),
		FINISH:    time.Now(),
		BOTID:     bots,
		RESULTMAP: res,
		LISTING:   "bla-bla-bla",
		NOTES:     notes}

	var game2 = &structs.Game{
		GID:       "id2",
		START:     time.Now(),
		FINISH:    time.Now(),
		BOTID:     bots,
		RESULTMAP: res,
		LISTING:   "bla-bla-bla2",
		NOTES:     notes}

	var game3 = &structs.Game{
		GID:       "id3",
		START:     time.Now(),
		FINISH:    time.Now(),
		BOTID:     bots,
		RESULTMAP: res,
		LISTING:   "bla-bla-bla3",
		NOTES:     notes}

	result[0] = game1
	result[1] = game2
	result[2] = game3

	b, _ := json.Marshal(result)
	w.Write(b)
}
