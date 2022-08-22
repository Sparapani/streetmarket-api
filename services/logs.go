package services

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Sparapani/streetmarketapi/models"
)

var logChannel chan logType

type logType struct {
	Sm     models.StreetMarket
	Oper   string
	Result string
}

func InitLog() {
	logChannel = make(chan logType, 100)
	logChannel <- logType{models.StreetMarket{}, "INIT", "STARTING API"}
	go logStatus(logChannel)
}
func FinishLog() {
	close(logChannel)
}

func logStatus(logChannel <-chan logType) {
	logFile, err := os.OpenFile("./logs/smlogs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("could not create file. err [%s]", err)
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			log.Printf("could not close file. err [%s]", err)
		}
	}()

	logWriter := bufio.NewWriter(logFile)

	for {
		value, hasNext := <-logChannel
		if !hasNext {
			break
		}
		_, err = logWriter.WriteString(fmt.Sprintf("%s,%s,%s\n", value.Sm, value.Oper, value.Result))
		if err != nil {
			log.Printf("could not write to file. err [%s]", err)
		}
		err = logWriter.Flush()
		if err != nil {
			log.Printf("could not flush writer. err [%s]", err)
		}
	}
}
