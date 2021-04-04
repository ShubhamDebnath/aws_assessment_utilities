package main

import (
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
)

func createRandomRecords(data []TimeLineData) []TimeLineData {
	var randomData []TimeLineData
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(data)
	idxMap := make(map[int]int)

	for i := 0; i < 100; {
		idx := r.Intn(size)
		val := r.Intn(100)

		_, prs := idxMap[idx]

		if prs == false {
			idxMap[idx] = val
			randomData = append(randomData, TimeLineData{data[idx].Date, val})
			i++
		}

	}

	return randomData
}

// Infinite loop
// running every 5 mins
// and uploading random data
func scheduleUpdates(sess *session.Session, data []TimeLineData) {
	for {
		records := createRandomRecords(data)
		batchDump(sess, records)
		// fmt.Println(records)
		time.Sleep(5 * time.Minute)
	}
}
