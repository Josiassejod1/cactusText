package trav

import (
  "fmt"
  "github.com/spf13/cobra"
  api "github.com/Josiassejod1/cactusText/api"
  "strings"
  "os"
)

var GetLyricsCmd = &cobra.Command{
  Use: "lyrics",
  Short: "Get Lyrics",
  Run: func(cmd *cobra.Command, args []string) {

    if len(os.Args) < 3 {
      fmt.Println(`
          Your Missing Some Search Terms
        `)
      os.Exit(1)
    }
    search := strings.Join(args, " ")
    fmt.Println("Searching For Lyrics 🔎")
    lyrics := api.GetLyrics(search)


    _ = api.LyricLoop(lyrics)
  },
}
