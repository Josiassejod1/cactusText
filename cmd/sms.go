package trav

import (
  "fmt"
  "github.com/spf13/cobra"
  http "github.com/Josiassejod1/cactusText/http/helpers"
  "strings"
)

var SendSMS = &cobra.Command{
  Use: "sms",
  Short: `
    Send Lyric -- required params
    [ACCOUNT_SSID, ACCOUNT_TOKEN, CONTACT NUMBER, TWILIO NUMBER, BODY]

    ex: (
        string afdasfa,
        string key_23435,
        string +1234567893,
        string +9979731234,
        string"LETS GOOOOOO"
      )
  `,
  Run: func(cmd *cobra.Command, args []string) {
    key := args[0]
    pass := args[1]
    to := args[2]
    from := args[3]
    body := args[4]

    fmt.Printf(args)
}
