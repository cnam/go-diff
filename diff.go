package main

import (
    "fmt"
    "github.com/sergi/go-diff/diffmatchpatch"
    "net/http"
    "io/ioutil"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Hello world")
    router := gin.Default()
    dmp := diffmatchpatch.New()
    s1 := readFile("examples/speedtest1.txt")
    s2 := readFile("examples/speedtest2.txt")

    router.GET("/", func(c *gin.Context) {

        diffs := dmp.DiffMain(s1, s2, true)

        c.HTMLString(http.StatusOK, dmp.DiffPrettyHtml(diffs))
    })

    router.Run(":8080")
}

func readFile(filename string) string {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        
    }
    return string(bytes)
}