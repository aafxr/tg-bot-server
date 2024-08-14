package main

import (
	"encoding/json"
	"log"
)

func main() {
	s := []string{"str 1", "str 2"}

	d, _ := json.Marshal(s)
	log.Println(string(d))

}
