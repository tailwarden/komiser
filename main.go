package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/tailwarden/komiser/cmd"
)

func main() {
	func() {
		defer func() {
			err := recover()
			fmt.Println(err)

			if err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
			}
		}()

		cmd.Execute()
	}()
}
