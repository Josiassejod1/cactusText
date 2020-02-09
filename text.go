package main

import (
  "net/http"
  "fmt"
  "log"
  "io/ioutil"
)

func main() {
  getLyrics("hey")
}

func getLyrics(url string) {
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", url)

  client := &http.Client{}
  req, err := http.NewRequest("GET", urlStr, nil)

  if err != nil {
    log.Fatalln(err)
  }

//  defer req.Body.Close()

  resp, _ := client.Do(req)

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }
  bodyString := string(body)
  fmt.Printf(bodyString)
}
