package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ecutdavid/codeforcesreport"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please tell me the user handle :)")
	}
	handle := os.Args[1]
	res := codeforcesreport.FetchSubmissions(handle)

	counter, lines := 0, []string{}
	for _, v := range codeforcesreport.ParseUniqOkSubmissions(res) {
		counter++
		id := fmt.Sprintf("%d/%s", v.Problem.ContestId, v.Problem.Index)
		y, m, d := time.Unix(v.CreationTimeSeconds, 0).Date()
		date := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		lines = append(lines, fmt.Sprintf("%-9s%-50s%-12s", id, v.Problem.Name, date))
	}

	fmt.Printf("%s solved %d problems.\n", handle, counter)
	fmt.Printf("%-9s%-50s%-12s\n", "ID", "Name", "Date")
	for _, l := range lines {
		fmt.Println(l)
	}
}
