package main

import (
	"log"

	"github.com/gorsuch/sns"
)

func main() {
	c := sns.New("http://sns.us-east-1.amazonaws.com/")
	err := c.Publish(
		"arn:aws:sns:us-east-1:854436987475:aggregates",
		"test",
		"message",
	)

	if err != nil {
		log.Fatal(err)
	}
}
