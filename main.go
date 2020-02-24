package main

import (
  trav "github.com/josiassejod1/cactusText/cmd"
  "github.com/joho/godotenv"
   "log"
   "github.com/spf13/cobra"
)

func init() {
  if err := godotenv.Load();
  err != nil {
    log.Print("No .env file found")
  }
    RootCmd.AddCommand(trav.GetLyricsCmd)
    RootCmd.AddCommand(trav.SendSMS)
}

var RootCmd = &cobra.Command{
  Use: "cactus",
  Short: `
    🌵🌵 Cactus Jack Lyrics 🌵🌵
    A tool to search Travis Scott Lyrics 🎶
  `,
}


func main() {
  RootCmd.Execute()
}
