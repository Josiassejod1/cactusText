package helpers

import (
  "net/http"
  "log"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "bytes"
)

type Lyric struct {
  Title string
  Lyric string
  Song_ID int
  Line_Number int
}

func MakeHttpRequest(url string) *http.Response{
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    log.Fatalln(err)
  }

  resp := LogHttpRequest(req)

  return resp
}

func GetLyrics(url string) [] Lyric {
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", url)
  resp := MakeHttpRequest(urlStr)
  lyric_array := DecodeLyrics(resp)

  fmt.Printf("\n Lyrics : %+v", lyric_array)
  return lyric_array
}

func DecodeLyrics(resp *http.Response) []Lyric {
  var lyrics []Lyric
  err := json.NewDecoder(resp.Body).Decode(&lyrics)
	if err != nil {
		log.Fatal(err)
	}

  return lyrics
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

  bodyString := string(body)

  log.Printf(bodyString)

  return resp
}
