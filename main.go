package main

import (
  // "net/http"
  // "fmt"
  "github.com/joho/godotenv"
   "log"
  // "io/ioutil"
  // "encoding/json"
  // "bytes"
  helpers "github.com/josiassejod1/cactusText/http/helpers"
)

func init() {
  if err := godotenv.Load();
  err != nil {
    log.Print("No .env file found")
  }
}


func main() {
  helpers.GetLyrics("sicko")
}
