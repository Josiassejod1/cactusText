package trav

import (
  "fmt"
  "github.com/spf13/cobra"
  twilio "github.com/Josiassejod1/cactusText/twilio"
  api "github.com/Josiassejod1/cactusText/api"
  "os"
)

var SendSMS = &cobra.Command{
  Use: "sms-random",
  Short: `
    Send Random Lyrics to Your Phone -- required params
    CONTACT NUMBER
    ex: (
        string +9979731234,
      )
  `,
  Run: func(cmd *cobra.Command, args []string) {
    if len(os.Args) > 1{
      toSender := args[0]
      randomLyric := api.GetRandomLyrics()
      twilio.SendMsgEZ(toSender, randomLyric)
    } else {
      fmt.Println(`
          Missing contact number, refer to Twilio API for correct format
        `, args[0])
      os.Exit(1)
    }

},
}
