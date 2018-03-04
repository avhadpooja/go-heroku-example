package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
	"io"
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
	//fmt.Fprintln(res, "hello, heroku")
}*/

func hello(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
    // in your case file would be fileupload
    file, header, err := r.FormFile("myfile")
    if err != nil {
        panic(err)
    }
    defer myfile.Close()
    name := strings.Split(header.Filename, ".")
    fmt.Printf("File name %s\n", name[0])
    // Copy the file data to my buffer
    io.Copy(&Buf, file)
    // do something with the contents...
    // I normally have a struct defined and unmarshal into a struct, but this will
    // work as an example
    contents := Buf.String()
    fmt.Println(contents)
    // I reset the buffer in case I want to use it again
    // reduces memory allocations in more intense projects
    Buf.Reset()
    // do something else
    // etc write header
    return
}
