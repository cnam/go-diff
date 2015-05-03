package main

import (
    "fmt"
    "github.com/sergi/go-diff/diffmatchpatch"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Hello world")
    router := gin.Default();
    router.GET("/", func(c *gin.Context) {
        dmp := diffmatchpatch.New()
        diffs := []diffmatchpatch.Diff{
            diffmatchpatch.Diff{diffmatchpatch.DiffEqual, "jump"},
            diffmatchpatch.Diff{diffmatchpatch.DiffDelete, "s"},
            diffmatchpatch.Diff{diffmatchpatch.DiffInsert, "ed"},
            diffmatchpatch.Diff{diffmatchpatch.DiffEqual, " over "},
            diffmatchpatch.Diff{diffmatchpatch.DiffDelete, "the"},
            diffmatchpatch.Diff{diffmatchpatch.DiffInsert, "a"},
            diffmatchpatch.Diff{diffmatchpatch.DiffEqual, " lazy"}}

        c.HTMLString(http.StatusOK, dmp.DiffPrettyHtml(diffs))
    })

    router.Run(":8080")
}