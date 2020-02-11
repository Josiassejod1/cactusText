package main

import (
  "net/http"
  "fmt"
  "log"
  "io/ioutil"
  "encoding/json"
  "bytes"
)

type Lyric struct {
  Title string
  Lyric string
  Song_ID int
  Line_Number int
}

 /*
 lyricJson := `[{
     "title": "Highest in The Room",
     "lyric": "She give me bright ideas",
     "song_id": 123,
     "line_number": 2
   }]`
 */

func main() {
  getLyrics("sicko")
}

func getLyrics(url string) {
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", url)

  /*
     Access one record
     md[0].(map[string]interface{})["title"]
     md[0].(map[string]interface{})["song_id"]
     md[0].(map[string]interface{})["line_number"]
     md[0].(map[string]interface{})["lyric"]
  */

  var lyric []Lyric

  resp := make_http_request(urlStr)
  lyric_array := deconstructJson(lyric, resp)

  fmt.Printf("\n Lyrics : %+v", lyric_array)
}


func make_http_request(url string) *http.Response{
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    log.Fatalln(err)
  }

  resp := log_http_request(req)

  return resp
}


func deconstructJson(i interface {}, resp *http.Response) interface{} {
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     log.Fatal(err)
   }
   json.Unmarshal([]byte(body), &i)
   return i
}

func log_http_request(req *http.Request ) *http.Response {
  client := &http.Client{}

  resp, _ := client.Do(req)
  var body [] byte

  if resp != nil {
      body, _ = ioutil.ReadAll(resp.Body)
  } else {
    log.Printf("Response is Nil")
  }

  resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  bodyString := string(body)

  log.Printf(bodyString)

  return resp
}
