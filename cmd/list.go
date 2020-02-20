package trav

import (
  "fmt"
  "github.com/spf13/cobra"
  http "github.com/Josiassejod1/cactusText/http/helpers"
  "strings"
)

var GetLyricsCmd = &cobra.Command{
  Use: "lyrics",
  Short: "Get Lyrics",
  Run: func(cmd *cobra.Command, args []string) {
    search := strings.Join(args, " ")
    fmt.Println("Searching For Lyrics 🔎")
    lyrics := http.GetLyrics(search)


    if len(lyrics) == 0 {
      fmt.Println("No Lyrics Found 🙁")
    } else {
      fmt.Println("found: ", len(lyrics), " lyrics ")
      for i, lyric := range lyrics {
        fmt.Printf("%d: %v\n", i + 1, lyric.Lyric)
      }
    }
  },
}
