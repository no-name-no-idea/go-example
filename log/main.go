package main

import (
	"fmt"
	"log"
	"os"
)

func WriteLogs(logs string) (err error) {
	data, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	log.SetOutput(data)
	log.Println(logs)

	return
}

func main() {
	logs := "Start Log"
	WriteLogs(logs)

	logs = "End Log"
	WriteLogs(logs)
}
