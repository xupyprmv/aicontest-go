package core

import (
	"encoding/json"
	"fmt"
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

	stmtIns, err := GetConnection().Prepare("INSERT INTO BOT (SOURCE, LANGUAGE) VALUES( ?, ? )")
	//	defer stmtIns.Close()
	fmt.Println(bot.SOURCE)
	_, err = stmtIns.Exec(bot.SOURCE, bot.LANGUAGE)
	if err != nil {
		panic(err.Error())
	}

	w.Write([]byte(`{"result":"OK"}`))
}
