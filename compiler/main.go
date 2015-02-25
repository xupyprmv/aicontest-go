/* compiler project main.go
Simple utility:
	- Get uncompiled bots from database
    - Try to compile each bot
	- Update database with result of compilation:
		COMPILE:
			0 - NOT COMPILED
			1 - COMPILE SUCCESS
			2 - COMPILE ERROR
		COMPILEOUTPUT:
			output of compiler
*/
package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"github.com/xupyprmv/aicontest-go/frontend/structs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

import _ "github.com/go-sql-driver/mysql"

func main() {

	config, err := getConfiguration()
	if err != nil {
		log.Fatalf("getConfiguration: %v", err)
		return
	}

	err = os.MkdirAll(config.Result_dir, os.ModePerm)
	if err != nil {
		log.Fatalf("os.MkdirAll: %v", err)
		return
	}

	log.Println("CONNECT TO > " + config.Database_url)
	db, err := sql.Open("mysql", config.Database_url)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
		return
	}

	result, err := db.Query(`SELECT BID, LANGUAGE, SOURCE FROM aicontest.BOT WHERE COMPILED=FALSE`)
	if err != nil {
		log.Fatalf("db.Query: %v", err)
		return
	}

	defer result.Close()
	for result.Next() {
		bot := new(structs.Bot)
		err = result.Scan(&bot.BID, &bot.LANGUAGE, &bot.SOURCE)
		if err != nil {
			log.Fatalf("result.Scan: %v", err)
			return
		}

		fileName, err := copySourceToTemporaryFile(bot)
		if err != nil {
			log.Fatalf("copySourceToTemporaryFile: %v", err)
			return
		}

		for _, compiler := range config.Compilers {
			if compiler.Language == bot.LANGUAGE {
				m := make(map[string]string)
				m["$1"] = fileName
				m["$2"] = config.Result_dir + "/" + bot.BID

				precompileString := addParamsToString(m, compiler.Precompile)
				execString := addParamsToString(m, compiler.Compile_string)
				postcompileString := addParamsToString(m, compiler.Postcompile)

				executeString(precompileString)
				log.Println("COMPILE STRING > " + execString)
				es := strings.Fields(execString)
				cmd := exec.Command(es[0], es[1:len(es)]...)

				stdout, err := cmd.StdoutPipe()
				if err != nil {
					log.Fatalf("cmd.StdoutPipe: %v", err)
					return
				}
				inout := bufio.NewScanner(stdout)

				stderr, err := cmd.StderrPipe()
				if err != nil {
					log.Fatalf("cmd.StderrPipe: %v", err)
					return
				}
				inerr := bufio.NewScanner(stderr)

				if err := cmd.Start(); err != nil {
					log.Fatalf("cmd.Start: %v")
				}

				errormsg := ""
				for inerr.Scan() {
					errormsg += inerr.Text() + "\n"
				}
				outmsg := ""
				for inout.Scan() {
					outmsg += inout.Text() + "\n"
				}

				if err := cmd.Wait(); err != nil {
					if exiterr, ok := err.(*exec.ExitError); ok {
						// The program has exited with an exit code != 0
						if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
							log.Printf("Exit Status: %d", status.ExitStatus())
							stmtUpd, err := db.Prepare(`UPDATE aicontest.BOT SET COMPILED = 2, COMPILEOUTPUT = ? WHERE BID = ?;`)
							if err != nil {
								log.Fatalf("db.Prepare: %v", err)
								return
							}
							_, err = stmtUpd.Exec("Exit Status: "+strconv.Itoa(status.ExitStatus())+"\nERROR LOG:\n"+errormsg+"\nOUTPUT LOG:\n"+outmsg, bot.BID)
							if err != nil {
								log.Fatalf("stmtUpd.Exec: %v", err)
								return
							}
						}

					} else {
						log.Fatalf("cmd.Wait: %v", err)
					}
				} else {
					executeString(postcompileString)
					stmtUpd, err := db.Prepare(`UPDATE aicontest.BOT SET COMPILED = 1, COMPILEOUTPUT = ? WHERE BID = ?;`)
					if err != nil {
						log.Fatalf("db.Prepare: %v", err)
						return
					}
					_, err = stmtUpd.Exec("ERROR LOG:\n"+errormsg+"\nOUTPUT LOG:\n"+outmsg, bot.BID)
					if err != nil {
						log.Fatalf("stmtUpd.Exec: %v", err)
						return
					}
				}
			}
		}
	}

	defer db.Close()
}

type Lang struct {
	Language       string
	Precompile     string
	Compile_string string
	Postcompile    string
}

type Configuration struct {
	Database_url string
	Result_dir   string
	Compilers    []Lang
}

func getConfiguration() (conf Configuration, err error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	return conf, err
}

func addParamsToString(m map[string]string, str string) (res string) {
	for key, value := range m {
		str = strings.Replace(str, key, value, -1)
	}
	return str
}

func copySourceToTemporaryFile(bot *structs.Bot) (fileName string, err error) {
	file, err := ioutil.TempFile(os.TempDir(), "aicontest")
	defer file.Close()
	ioutil.WriteFile(file.Name(), []byte(bot.SOURCE), os.ModeExclusive)
	return file.Name(), err
}

func executeString(execString string) {
	if execString == "" {
		return
	}
	commands := strings.Split(execString, " && ")

	for _, command := range commands {
		es := strings.Fields(command)
		log.Println("EXECUTE STRING > " + command)
		cmd := exec.Command(es[0], es[1:len(es)]...)
		if err := cmd.Start(); err != nil {
			log.Fatalf("cmd.Start: %v", err)
		}
	}
	return
}
