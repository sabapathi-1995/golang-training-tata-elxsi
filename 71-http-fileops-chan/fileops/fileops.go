package fileops

import (
	"errors"
	"log/slog"
	"os"
)

var ChanData chan []byte

func init() { // automatically called
	ChanData = make(chan []byte, 10)
}

func Init(filename string) {
	go Save(filename)
}

func Save(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	defer file.Close()

	for data := range ChanData {
		if data == nil || len(data) == 0 {
			//return errors.New("nil data or no data")
			slog.Error("no data of the user")
		}

		data = append(data, '\n')
		_, err = file.Write(data)
		if err != nil {
			slog.Error(err.Error(), "data:", data)
		}

	}
}

func SaveToFile(fileName string, bytes []byte) error {

	if bytes == nil || len(bytes) == 0 {
		return errors.New("nil data or no data")
	}

	bytes = append(bytes, '\n')
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil

}
