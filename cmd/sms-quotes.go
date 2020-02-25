package trav

import (
  "fmt"
  "github.com/spf13/cobra"
   twilio "github.com/Josiassejod1/cactusText/twilio"
   api "github.com/Josiassejod1/cactusText/api"
    "os"
)

var SendSMSSearch = &cobra.Command{
  Use: "sms-quotes",
  Short: `
    Search For Quotes -- required params
    [CONTACT NUMBER, BODY]
    ex: (
        string +9979731234,
        string"LETS GOOOOOO"
      )
  `,
  Run: func(cmd *cobra.Command, args []string) {

    if len(os.Args) > 3 {
      toSender := args[1]
      bodyMessage := args[2]

      quote := api.GetLyrics(bodyMessage)
      results := api.LyricLoop(quote)
      twilio.SendMsgEZ(toSender,string(results))
    } else {
      fmt.Println(`
          Paramater [0] : Phone Number
          Missing contact number, refer to Twilio API for correct format

          Paramater [1] : Keyword
          Type in Some Lyrics
        `)
      os.Exit(1)
    }


},
}
