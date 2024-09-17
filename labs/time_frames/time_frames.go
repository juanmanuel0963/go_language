/*
Given calendar (meeting) event start/end datetime in JSON with ISO8601 format:

Given working hours are 8am - 5pm.
Print 30 minute or longer contiguous blocks of time that are free and without events/meetings.
Example output:

2024-02-09T08:00:00-0700 - 2024-02-09T10:00:00-0700
2024-02-09T10:30:00-0700 - 2024-02-09T11:00:00-0700

workdayStart, _ := time.Parse(time.RFC3339, "2024-02-09T08:00:00-0700")
*/
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeBlock struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func main() {

	// Modified JSON object with meetings to match RFC3339 format
	myMeetings := []byte(
		`[{"start":"2024-02-09T10:00:00-07:00","end":"2024-02-09T10:30:00-07:00"},
    	  {"start":"2024-02-09T11:00:00-07:00", "end":"2024-02-09T12:00:00-07:00"},
    	  {"start":"2024-02-09T12:00:00-07:00", "end":"2024-02-09T17:00:00-07:00"}]`)

	lMeetings := []TimeBlock{}

	err := json.Unmarshal(myMeetings, &lMeetings)

	if err != nil {
		fmt.Println(err)
	}

	//workdayStart, err := time.Parse(time.RFC3339, lMeetings[0].Start)

	for i := 0; i < len(lMeetings); i++ {
		// Using RFC3339 since the JSON format has been adjusted
		workdayStart, err := time.Parse(time.RFC3339, lMeetings[i].Start)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}
		fmt.Println(workdayStart)

		workdayEnd, err := time.Parse(time.RFC3339, lMeetings[i].End)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}
		fmt.Println(workdayEnd)

		fmt.Println("-----")
	}
}
