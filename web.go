package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
)

func main() {
	http.HandleFunc("/#", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

/*func hello(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("client.html")
        t.Execute(res, nil)
	//fmt.Fprintln(res, "hello, heroku")
}*/

func hello(res http.ResponseWriter, req *http.Request) {
	err :=r.ParseForm()
	if err != nil{}
	domains := r.PostFormFile("myfile")
	fmt.Fprintf(w, "List is %s", myfile)
	
	//fmt.Fprintln(res, "hello, heroku")
