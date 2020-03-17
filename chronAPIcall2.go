package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "github.com/jasonlvhit/gocron"
)

// go run chronapi3.go

func main() {
    s := gocron.NewScheduler()
    s.Every(1).Minute().Do(task)
    <- s.Start()
}

func task() {
  resp, err := http.Get("http://0.0.0.0/api/news/csv")
  if err != nil {
    log.Fatalln(err)
    fmt.Println(err)
  }
  //fmt.Println(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  log.Println(string(body))
}

