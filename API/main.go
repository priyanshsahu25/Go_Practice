package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Joke struct {
	ID        int    `json:"id"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	res, err := http.Get("https://official-joke-api.appspot.com/jokes/ten")
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println(err)
	} else {
		fmt.Println("Data is fetched")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var jokes []Joke
	err1 := json.Unmarshal(data, &jokes)
	if err1 != nil {
		fmt.Println(err1)
	}

	for _, val := range jokes {
		fmt.Println("ID:", val.ID, "\n Setup:", val.Setup, "\n Punchline:", strings.ToUpper(val.Punchline))
	}
}
