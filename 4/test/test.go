package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

var jsonMessage = []byte(`
{"Name": "Anatolij","Body":"Negr","Time":6756577}
`)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Fprintf(w, "Hi there, I love "+path[1:]+"!")
		m := Message{}
		err := json.Unmarshal(jsonMessage, &m)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(m.Body))
	})
	http.ListenAndServe(":1111", nil)

}
