package main

import (
  "fmt"
  "time"
  "encoding/json"
)

type LogMessage struct {
  ExampleTime time.Time
  ExampleMessage string
}
func main() {
  for {
    msg := &LogMessage{
      ExampleMessage: "Hi",
      ExampleTime: time.Now(),
    }
    json, _  := json.Marshal(msg)
    fmt.Println(string(json))
    time.Sleep(3 * time.Second)
  }
}
