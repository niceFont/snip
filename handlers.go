package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

//IndexHandler sending index.html and pushing client.js
func IndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if pusher, ok := w.(http.Pusher); ok {

		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": r.Header["Accept-Encoding"],
			},
		}

		if err := pusher.Push("./static/client.js", options); err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)

}

//SnipHandler handling POST requests to shorten URL's
func SnipHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var shortenedURL bytes.Buffer
	resStruct := ResponseData{}

	shortenedURL.WriteString(DOMAIN)
	dataS := new(Data)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading Body")
		http.Error(w, "Cant read body", http.StatusBadRequest)
	}

	err = json.Unmarshal([]byte(body), dataS)
	if err != nil {
		fmt.Println(err)
	}

	key := HashURL(dataS.Val)

	List[key] = dataS.Val
	shortenedURL.WriteString(key)

	resStruct.OriginalURL = dataS.Val
	resStruct.ShortURL = shortenedURL.String()

	resString, _ := json.Marshal(resStruct)

	fmt.Fprintln(w, string(resString))
	fmt.Printf("%s", string(resString))
}

//RedirectHandler matches looks up Url in Map and redirects or sends http.Error() when it doesn't exist
func RedirectHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if url := List[ps[0].Value]; url != "" {

		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	http.Error(w, "URL Doesn't Exist", http.StatusBadRequest)

}
