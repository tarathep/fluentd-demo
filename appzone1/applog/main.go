package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fluent/fluent-logger-golang/fluent"
)

func main() {
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: string(os.Getenv("HOST"))})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	tag := "fluentd.test"

	var data = map[string]string{
		"from": "GO",
		"to":   "Monitor",
	}

	error := logger.Post(tag, data)
	if error != nil {
		panic(error)
	}
	fmt.Println("LOG APP")

	//---------- NAS ------------

	var dataNas = map[string]string{
		"from": "NAS",
		"to":   "Monitor",
	}
	//CONVERT TO JSON
	jsonString, err := json.Marshal(dataNas)
	if err != nil {
		panic(err)
	}
	//WRITE FILE
	writeLog("fluentd.test", string(jsonString))
	fmt.Println("LOG NAS")

}

func writeLog(fileName string, data string) {
	f, err := os.OpenFile("log/"+fileName+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data + "\n"); err != nil {
		log.Println(err)
	}
}
