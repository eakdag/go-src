package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Match struct {
	Sequence int    `json="seq"`
	Score    string `json="s"`
	Playing  int    `json="p"`
	Status   string `json="st"`
}

func getMatches() {
	resp, err := http.Get("http://www.mackolik.com/Match/MatchStatusHandler.ashx")

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s", body)

	var matchList map[string]Match

	err = json.Unmarshal(body, &matchList)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%+v", matchList)
}

func handler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	matchId := v.Get("id")

	fmt.Printf("url=%s, matchId=%s ", r.URL, matchId)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	serveSingle("/sitemap.xml", "./sitemap.xml")
	serveSingle("/favicon.ico", "./favicon.ico")
	serveSingle("/robots.txt", "./robots.txt")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	//getMatches()
}
