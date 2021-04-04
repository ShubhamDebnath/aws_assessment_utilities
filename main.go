package main

func main() {

	records := readCsv("./multiTimeline.csv")
	// jsonBody, _ := json.Marshal(records)
	// fmt.Println(string(jsonBody))

	// creating aws sdk client session
	// with hardcoded credentials
	sess := configureAWS()

	// creating table
	// createTable(sess)
	// batchDump(sess, records)

	// deleteTable(sess)

	// streamArn := listStreams(sess)
	// shardIds := describeStream(sess, streamArn)

	// for idx, shardId := range shardIds {
	// 	fmt.Printf("shardIterator index:: %d", idx)
	// 	shardIterator := getShardIterator(sess, streamArn, shardId)
	// 	getRecords(sess, *shardIterator)
	// }

	// createEsDomain(sess)

	// createLambdaFunction(sess)
	// createLambdaTrigger(sess, streamArn)

	scheduleUpdates(sess, records)

}
