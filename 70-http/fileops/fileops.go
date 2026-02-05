package fileops

import (
	"errors"
	"os"
)

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
