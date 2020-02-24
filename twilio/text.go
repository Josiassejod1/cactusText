package twilio

import (
  "fmt"
  "os"
  "net/url"
  "strings"
  "net/http"
  "encoding/json"
)

func validateTwilio() (string, string, string){
  key, key_err := os.LookupEnv("ACCOUNT_SSID")

  if key_err {
    fmt.Println("Account SSID exists, %s", key)
  } else {
    fmt.Println("Account SSID not found")
    os.Exit(1)
  }

  pass, pass_err := os.LookupEnv("ACCOUNT_TOKEN")

  if pass_err {
    fmt.Println("Account Token exists, %s", pass)
  } else {
    fmt.Println("Account Token not found")
    os.Exit(1)
  }

  twilio, twilio_err := os.LookupEnv("ACCOUNT_NUMBER")

  if twilio_err {
    fmt.Println("Account Number exists %s", twilio)
  } else {
    fmt.Println("Account Number not found")
    os.Exit(1)
  }

  return  key, pass, twilio
}



func SendMsgEZ(to string, body string) {
  key, pass, from  := validateTwilio()
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + key + "/Messages.json"
  msgData := url.Values {}
  msgData.Set("To", to) //Number to send
  msgData.Set("From", from) //Number from
  msgData.Set("Body", body)
  msgDataReader := *strings.NewReader(msgData.Encode())

  client := & http.Client {}
  req, _ := http.NewRequest("POST", urlStr, & msgDataReader)
  req.SetBasicAuth(key, pass)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  resp,
  _ := client.Do(req)

  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string] interface {}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode( & data)
    if (err == nil) {
      fmt.Println("🌵: Message Sent")
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}

func SendMsg(key string, pass string, to string, from string, body string) {
    urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + key + "/Messages.json"
    msgData := url.Values {}
    msgData.Set("To", to) //Number to send
    msgData.Set("From", from) //Number from
    msgData.Set("Body", body)
    msgDataReader := *strings.NewReader(msgData.Encode())

    client := &http.Client {}
    req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
    req.SetBasicAuth(key, pass)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp,
    _ := client.Do(req)

    if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
      var data map[string] interface {}
      decoder := json.NewDecoder(resp.Body)
      err := decoder.Decode( & data)
      if (err == nil) {
        fmt.Println("🌵: Message Sent")
        fmt.Println(data["sid"])
      }
    } else {
      fmt.Println(resp.Status);
    }
}
