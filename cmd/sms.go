package trav

import (
  "fmt"
  "github.com/spf13/cobra"
  twilio "github.com/Josiassejod1/cactusText/twilio"
  http "github.com/Josiassejod1/cactusText/http/helpers"
  "os"
)

var SendSMS = &cobra.Command{
  Use: "sms",
  Short: `
    Send Lyric -- required params
    [CONTACT NUMBER, BODY]
    ex: (
        string +9979731234,
        string"LETS GOOOOOO"
      )
  `,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args[0]) < 0 {
      fmt.Println(`
          You input : %s
          Missing contact number, refer to Twilio API for correct format
        `, args[0])
      os.Exit(1)
    }

    if len(args[1]) < 0 {
      fmt.Println(`
        You input : %s
        Missing Seach Keyword, Type in Some Lyrics
        `, args[1])
      os.Exit(1)
    }

    toSender := args[0]
    bodyMessage := args[1]
    randomLyric := http.GetRandomLyric(bodyMessage)
    twilio.SendMsgEZ(toSender, randomLyric)
},
}
