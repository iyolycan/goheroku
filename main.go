package main

import (
   "net/http"
    "fmt"
    "strconv"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {

    d := "<html><table border=1>" ;

    for y:=0; y<10; y++{
        d += "<tr>"
        for x:=0; x<10; x++{
            d += "<td>" + strconv.Itoa(y) + "x" + strconv.Itoa(x) + "=" + strconv.Itoa(x*y ) + "</td>"
        }
        d += "</tr>"
    }

    d += "</table></html>"

    fmt.Fprint(w, d)
}

func main() {
    http.HandleFunc( "/" , indexHandle)
    http.ListenAndServe( ":80" , nil)
}