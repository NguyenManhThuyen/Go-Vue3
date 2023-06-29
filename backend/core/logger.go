package core

import (
	"api/config"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteLog(message string) {
	time.LoadLocation(config.Config("APP_TIME_ZONE"))
	dir_path := "./assets/log"
	file_name := fmt.Sprintf("%s/%s.txt", dir_path, time.Now().Format("2006-01-02"))

	log_time := time.Now().Format("2006-01-02 15:04:05")
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	defer f.Close()

	fmt.Println(123123)

	msg := fmt.Sprintf("[%s] | %s \n", log_time, message)
	f.WriteString(msg)

}
