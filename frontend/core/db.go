package core

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

var connection *sql.DB

func GetConnection() *sql.DB {
	fmt.Print("*1")
	var err error
	//	if connection == nil {
	fmt.Print("*2")
	connection, err = sql.Open("mysql", "aicontest:aicontest_pass@/aicontest")
	//	connection = db
	if err != nil {
		panic(err.Error())
	}
	_, err = connection.Exec(`CREATE TABLE IF NOT EXISTS aicontest.BOT (
  							BID VARCHAR(36)  NOT NULL,
  							CID VARCHAR(36)  NOT NULL,
  							UID VARCHAR(36)  NOT NULL,
  							LANGUAGE VARCHAR(10)  NOT NULL,
							SOURCE TEXT  NOT NULL,
  							COMPILED BOOLEAN  NOT NULL,
  							COMPILEOUTPUT TEXT ,
 							INDEX b(BID),
  							INDEX c(CID),
  							INDEX u(UID),
  							INDEX cu(CID, UID),
  							INDEX comp(COMPILED)
							) ENGINE = InnoDB;`)
	if err != nil {
		panic(err.Error())
	}
	//	}
	return connection
}
