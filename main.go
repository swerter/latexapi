package main

import (
        "fmt"
        "os"
        "io/ioutil"
        // "strings"
        // "math/rand"
        // "path/filepath"
        "os/exec"
        // "strconv"
        "github.com/gin-gonic/gin"
        "net/http"
)

const DeleteFolderConst string = "DELETE"
const RootPathConst string = "/var/mails"




// taken from http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang#10510783
// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func CompileLatex(text string) (string) {
        fmt.Println(text)
        dir := "/latexapi_tex"
        prefix := "latex_"

        file, err := ioutil.TempFile(dir, prefix)

        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Temp filename: %s\n", file.Name())
        err = ioutil.WriteFile(file.Name(), []byte(text), 0644)
        if err != nil {
                fmt.Printf("Error writing latex file: %s", err)
        }

        out, err := exec.Command("/usr/bin/xelatex", "-halt-on-error" ,file.Name()).Output()
        if err != nil {
                fmt.Printf("Error compiling latex file: %s", err)
        }
        fmt.Printf("%s", out)
        filename := fmt.Sprintf("%s.pdf", file.Name())
        // Remove all generated files
        err = os.Remove(file.Name())
        if err != nil {
                fmt.Printf("Error deleting file: %s", err)
        }
        err = os.Remove(fmt.Sprintf("%s.aux", file.Name()))
        if err != nil {
                fmt.Printf("Error deleting file: %s", err)
        }
        err = os.Remove(fmt.Sprintf("%s.log", file.Name()))
        if err != nil {
                fmt.Printf("Error deleting file: %s", err)
        }
        err = os.Remove(fmt.Sprintf("%s.out", file.Name()))
        if err != nil {
                fmt.Printf("Error deleting file: %s", err)
        }

        return filename
}



func CompileLatexEndpoint(c *gin.Context) {
        text := c.Query("text")
        filename := CompileLatex(text)
        c.File(filename)
        err := os.Remove(filename)
        if err != nil {
                fmt.Printf("Error deleting file: %s", err)
        }
}



func main() {
        router := gin.Default()

        router.LoadHTMLGlob("templates/*")

        // Global middlewares
        router.Use(gin.Recovery())

        router.GET("/ping", func(c *gin.Context) {
                c.String(200, "pong")
        })
        router.GET("/compile", CompileLatexEndpoint)
        router.POST("/compile", CompileLatexEndpoint)

        router.GET("/", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.tmpl", gin.H{})
        })

        // Listen and serve on 0.0.0.0:8080
        router.Run(":8080")
}
