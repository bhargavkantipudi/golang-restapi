package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sf-heating-process-service/logger"
	"strconv"
	"time"
)

func TimeStampGetStartDayOfWeek() int64 {
	tm := time.Now()
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return currentZeroDay.Add(-1 * (weekday) * 24 * time.Hour).Unix()
}

// CurrentTimestamp ...
func CurrentTimestamp() int {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	currentTimestamp, _ := strconv.Atoi(timestamp)
	//logger.Info.Println("Current Timestamp ", currentTimestamp)
	return currentTimestamp
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func ReadFile(filename string) string {
	tdata, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Info.Println("File reading error", err)
	}
	return string(tdata)
}

// ModelToString ...
func ModelToString(m interface{}) string {
	e, _ := json.Marshal(m)
	e = bytes.Replace(e, []byte("\\u003c"), []byte("<"), -1)
	e = bytes.Replace(e, []byte("\\u003e"), []byte(">"), -1)
	e = bytes.Replace(e, []byte("\\u0026"), []byte("&"), -1)
	return string(e)
}
func CurrentHourInEpoach() int64 {
	now := time.Now().UTC().Unix()
	now = now - now%3600
	return now
}
