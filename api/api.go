package api

import (
  "net/http"
  helpers "github.com/Josiassejod1/cactusText/http/helpers"
  "fmt"
  "time"
  "math/rand"
  "log"
  "encoding/json"
)

type Lyric struct {
  Title string
  Lyric string
  Song_ID int
  Line_Number int
}

type FullLyrics struct {
  Songid int `json:"id"`
  Regulartitle string `json:"regular_title"`
  Aztitle string `json:"az_title"`
  Songlyric string `json:"song_lyrics"`
  Created string `json:"created_at"`
  Updated string `json:"updated_at"`
}

func GetLyrics(keyword string) []Lyric {
  urlStr := fmt.Sprintf("https://cactus-chorus.herokuapp.com/api/v1/lyrics/%s", keyword)
  resp := helpers.MakeHttpRequest(urlStr)
  lyric_array := DecodeLyrics(resp)
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

func LyricLoop(lyric_array []Lyric) string {
  var str string
  if len(lyric_array) == 0 {
    fmt.Println("No Lyrics Found 🙁")
  } else {
    fmt.Println("found: ", len(lyric_array), " lyrics ")
    for i, lyric := range lyric_array {

      str += fmt.Sprintf("%d: %v\n", i + 1, lyric.Lyric)
    }
  }

  fmt.Println(str)

  return str
}

func DecodeAllLyrics(resp *http.Response) []FullLyrics {
  var all []FullLyrics

  err := json.NewDecoder(resp.Body).Decode(&all)
  fmt.Println("%d", err)
	if err != nil {
		log.Fatal(err)
	}

  return all
}

func GetAllLyrics() []FullLyrics {
  urlStr := "https://cactus-chorus.herokuapp.com/api/v1/lyrics"
  resp := helpers.MakeHttpRequest(urlStr)
  lyric_array := DecodeAllLyrics(resp)
  return lyric_array
}



func GetRandomLyrics() string {
  var lyric = GetAllLyrics()
  len  := len(lyric)
  if len == 0 {
      return "Lyrics Not Found"
  }

  rand.Seed(time.Now().UTC().UnixNano())
  song := lyric[rand.Intn(len - 1)]

  output := fmt.Sprintf("%s \n \n %s", song.Regulartitle, song.Songlyric)
  return output
}
