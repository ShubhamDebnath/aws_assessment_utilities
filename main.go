package main

func main() {

	records := readCsv("./multiTimeline.csv")
	// fmt.Println(records)

	// creating aws sdk client session
	// with hardcoded credentials
	sess := configureAWS()

	// creating table
	createTable(sess)
	batchDump(sess, records)

	// deleteTable(sess)
}
