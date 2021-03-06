package core

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/xupyprmv/aicontest-go/frontend/structs"
	"net/http"
)

func SaveBot(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var bot structs.Bot
	err := decoder.Decode(&bot)
	if err != nil {
		panic(err.Error())
	}

	if bot.BID == "" {
		uuid, err := uuid.NewV4()
		if err != nil {
			panic(err.Error())
		}
		bot.BID = uuid.String()
	}

	stmtIns, err := GetConnection().Prepare("INSERT INTO BOT (BID, SOURCE, LANGUAGE) VALUES(?, ?, ? )")
	defer stmtIns.Close()
	fmt.Println(bot.SOURCE)
	_, err = stmtIns.Exec(bot.BID, bot.SOURCE, bot.LANGUAGE)
	if err != nil {
		panic(err.Error())
	}

	w.Write([]byte(`{"result":"OK"}`))
}
