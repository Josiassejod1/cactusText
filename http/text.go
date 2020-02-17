package main

import (
  "fmt"
  helpers "github.com/Josiassejod1/http/helpers"
  "github.com/joho/godotenv"
  "os"
  "log"
  "net/url"
  "strings"
  "net/http"
  "encoding/json"
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

func init(){
  if err := godotenv.Load(); err != nil {
    log.Print("No .env file foud")
  }
}

func main() {

  key, exists := os.LookupEnv("ACCOUNT_SSID")

  if exists {
    fmt.Println("key exists")
    fmt.Println(key)
  }

  pass, exists := os.LookupEnv("ACCOUNT_TOKEN")

  if exists {
    fmt.Println("pass exists")
    fmt.Println(pass)
  }
  md := getLyrics("sicko")
  lyric := md.([]interface{})[0].(map[string]interface{})["lyric"]
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/"+ key +"/Messages.json"

  fmt.Println(lyric.(string))
  msgData := url.Values{}
  msgData.Set("To","") //Number to send
  msgData.Set("From","") //Number from
  msgData.Set("Body",  "TEST: " + lyric.(string))
  msgDataReader := *strings.NewReader(msgData.Encode())
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(key, pass)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  resp, _ := client.Do(req)

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
  var data map[string]interface{}
  decoder := json.NewDecoder(resp.Body)
  err := decoder.Decode(&data)
  if (err == nil) {
    fmt.Println(data["sid"])
  }
} else {
  fmt.Println(resp.Status);
}
}

/*
   Access one record
   md[0].(map[string]interface{})["title"]
   md[0].(map[string]interface{})["song_id"]
   md[0].(map[string]interface{})["line_number"]
   md[0].(map[string]interface{})["lyric"]
*/


func getLyrics(url string) interface{}{
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", url)

  var lyric []Lyric
  resp := helpers.MakeHttpRequest(urlStr)
  lyric_array := helpers.DeconstructJson(lyric, resp)

  fmt.Printf("\n Lyrics : %+v", lyric_array)
  return lyric_array
}
