package main

import (
	//"crypto/md5"
	"io"
	//"strconv"
	//"time"
	"fmt"
	"net/http"
	"os"
	"html/template"
	
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

/*func hello(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("client.html")
        t.Execute(res, nil)
	fmt.Fprintln(res, "hello, heroku")
}
*/

func hello(w http.ResponseWriter, r *http.Request) {
 
        t, _ := template.ParseFiles("client.html")
        t.Execute(w, nil)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
 
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
 
        io.Copy(f, file)
	fmt.Println(f) 
 
}

 

