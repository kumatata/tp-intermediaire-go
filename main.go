package main

import (
"fmt"
"net/http"
"io"
"time"
"io/ioutil"
"os"
"log"
)

func main() {
	http.HandleFunc("/",getDateHandler)
	http.HandleFunc("/entries",getEntriesHandler)
	http.HandleFunc("/add", queryHandler)
	fmt.Println("Ecrivez vos params dans l'url de votre navigateur :)")
	http.ListenAndServe(":4567", nil)
   }

func getDateHandler(w http.ResponseWriter, req *http.Request) {

 if req.Method == http.MethodGet{
	getDate()
 }
	 
   }

func getDate(){
	t := time.Now()
	h := t.Hour()
	m := t.Minute()

	fmt.Println(h,":",m)
}

func getEntriesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost{
		if err := req.ParseForm(); err != nil {
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
		return
		}
		for key, value := range req.PostForm {
		fmt.Println(key+":",value)
		}
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
		}
		
	  }

func save(mydata []string) {
	for _, value := range mydata {

		f, err := os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(value); err != nil {
			log.Println(err)
		}
}
}

func queryHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    io.WriteString(res, req.FormValue("author") + ": ")
	io.WriteString(res, req.FormValue("entry"))
}

func getEntries() string{
	data, err := ioutil.ReadFile("save.data")
	if err == nil {
		fmt.Println(string(data))
		return string(data)
	}
	return ""
}
