package main

import (
  "fmt"
  helpers "github.com/Josiassejod1/http/helpers"
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

/*
   Access one record
   md[0].(map[string]interface{})["title"]
   md[0].(map[string]interface{})["song_id"]
   md[0].(map[string]interface{})["line_number"]
   md[0].(map[string]interface{})["lyric"]
*/


func getLyrics(url string) {
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", url)

  var lyric []Lyric
  resp := helpers.MakeHttpRequest(urlStr)
  lyric_array := helpers.DeconstructJson(lyric, resp)

  fmt.Printf("\n Lyrics : %+v", lyric_array)
}
