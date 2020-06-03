/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/negrel/gnotify"
	"github.com/spf13/cobra"
)

// Notification manager
var manager gnotify.Manager
var timer time.Duration = 0

// Initialize notification manager
func init() {
	var err error = nil
	manager, err = gnotify.New("TIMER")
	if err != nil {
		log.Fatal(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "timer",
	Short:   "a simple cli timer tool written in Go.",
	Long: `Timer cli tool that start a timer and show a desktop notification
when your timer is over.
`,
	Example: `  timer 25m15s - start a timer of 25 minutes and 25 seconds.`,
	// Parse args to check errors
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)
		if err != nil {
			return err
		}

		value := strings.Join(args, " ")
		timer, err = time.ParseDuration(value)
		if err != nil {
			e := fmt.Sprintf("error while parsing the given duration")
			return errors.New(e)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Timer start notification.
		notif := gnotify.NewNotification(gnotify.Option{
			Title:         "Starting a timer",
			Body:          fmt.Sprintf("Starting a timer of %v", args[0]),
			ExpireTimeout: 3000,
		})
		manager.Push(notif)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			done := time.Until(time.Now().Add(timer))
			<-time.After(done)

			// never expire until user click on "Ok".
			notif.ExpireTimeout = 0
			notif.Actions["Ok"] = func() {}
			// update title and body
			notif.Title = "Time is up"
			notif.Body = fmt.Sprintf("Your timer of %v is over.", args[0])
			manager.Push(notif)

			wg.Done()
		}()

		wg.Wait()
	},
}
