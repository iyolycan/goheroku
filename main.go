package main

import (
	"log"
	"net/http"
	"os"
    "fmt"
    "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
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
	
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
        // c.HTML(http.StatusOK, "index.tmpl.html", nil)
        http.HandleFunc( "/" , indexHandle)
	})

	router.Run(":" + port)
}
