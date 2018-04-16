//http 包创建web服务器
package main

import (
    "fmt"
    "log"
    "net/http"
)
type Mux struct {}

func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/" {
        says(w,r)
        return
    }

    http.NotFound(w,r)
    return
}

func says(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    fmt.Println("path:",r.URL.Path)
    fmt.Fprintf(w,"hello go")
}

func main() {
   // http.HandleFunc("/",says)
   mux := &Mux{}
   err :=http.ListenAndServe(":9091",mux)
    if err !=nil{
        log.Fatal("ListenServer:",err)
    }
}
