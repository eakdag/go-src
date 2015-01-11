package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
type Match struct {
	Sequence int    `json="seq"`
	Score    string `json="s"`
	Playing  int    `json="p"`
	Status   string `json="st"`
}*/
var ch = make(chan string)
var EventList [][]interface{}
var LatestEventId int

func main() {
	serveSingle("/sitemap.xml", "./sitemap.xml")
	serveSingle("/favicon.ico", "./favicon.ico")
	serveSingle("/robots.txt", "./robots.txt")
	http.HandleFunc("/", handler)
	LatestEventId = 0
	go getMatches()
	http.ListenAndServe(":9000", nil)

}

func getMatches() {
	fmt.Print("getMatches")
	resp, err := http.Get("http://www.mackolik.com/LiveScores/EventData.ashx?eId=0" + strconv.Itoa(LatestEventId))

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(body))

	var eventList [][]interface{}

	err = json.Unmarshal(body, &eventList)
	//if len(EventList)>0 && EventList[EventList.length - 1]

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%+v", eventList)

	ch <- string(body)

	for j, _ := range eventList {
		LatestEventId = eventList[j][0].(int)

		if len(EventList) == 20 {
			for i, _ := range EventList {
				if len(EventList)-1 != i {
					EventList[i] = EventList[i+1]
				}
			}
		}
		EventList[len(EventList)-1] = eventList[j]
	}

	time.Sleep(1000)

	getMatches()

}

func handler(w http.ResponseWriter, r *http.Request) {
	/*v := r.URL.Query()
	matchId := v.Get("id") */

	resp := <-ch

	fmt.Fprintf(w, resp)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
