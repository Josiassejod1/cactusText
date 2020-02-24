package helpers

import (
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  "bytes"
)


func MakeHttpRequest(url string) *http.Response{
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    log.Fatalln(err)
  }

  resp := LogHttpRequest(req)

  return resp
}

func DeconstructJson(i interface {}, resp *http.Response) interface{} {
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     log.Fatal(err)
   }
   json.Unmarshal([]byte(body), &i)
   return i
}

func LogHttpRequest(req *http.Request ) *http.Response {
  client := &http.Client{}

  resp, _ := client.Do(req)
  var body [] byte

  if resp != nil {
      body, _ = ioutil.ReadAll(resp.Body)
  } else {
    log.Printf("Response is Nil")
  }

  resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
  //use when trouble shooting
 //fmt.Println(resp.Body)
  return resp
}
