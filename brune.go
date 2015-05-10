package main

import (
    "fmt"
	"encoding/json"
    "net/http"
	"io/ioutil"
	"log"
)

var file []byte
var jsontype jsonobject

func bruneHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><link href='http://fonts.googleapis.com/css?family=Reenie+Beanie'rel='stylesheet' type='text/css'><style>span { font-family: 'Reenie Beanie', cursive; } div {font-size: 62; position: fixed; top: 50%%;  left: 50%%;  /* bring your own prefixes */  transform: translate(-50%%, -50%%);}</style><div><span>Brune blamed:&nbsp<em>%v</em></div></html>", jsontype.Object.Counter)
	data, err := json.Marshal(jsontype)
	jsontype.Object.Counter++
	err = ioutil.WriteFile("config.json", data, 755)
	if err != nil {
		fmt.Println("Can't write file.")
		return
	}
	log.Printf("Brunes counted: %v", jsontype)
}

func handlerICon(w http.ResponseWriter, r *http.Request) {} 

type jsonobject struct {
    Object ObjectType
}

type ObjectType struct {
	Counter int64
}

func main() {
	file, e := ioutil.ReadFile("config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
    }
    json.Unmarshal(file, &jsontype)
	http.HandleFunc("/favicon.ico", handlerICon) 
	http.HandleFunc("/", bruneHandler)
    http.ListenAndServe(":8080", nil)
	fmt.Print(jsontype)
}
