package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
    "log"
  //  "gopkg.in/yaml.v2"

)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, this is not a valid resource %s!", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
  // Standard format for API /api/v1/<service>/<extra params1>/...
    fmt.Fprintf(w, "this is the second api root")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/", handler)
    http.HandleFunc("/api", handler2)
    http.HandleFunc("/cmd", cmdHandler)
    fmt.Print("printing to systemout i think")
    filename := "somefile.txt"
    ioutil.WriteFile(filename, []byte("this line is written"), 0600)
    http.ListenAndServe(":8081", nil)
    fmt.Print("is the above line blocking this line?")
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Print("executing command: " )
  cmd :=exec.Command("touch", "thisisafile")
  err := cmd.Run()
  log.Printf("Command finished with error: %v", err)
}

/*
c style commenting
*/
// This is almost an object
type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
