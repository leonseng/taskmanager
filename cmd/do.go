/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		doTask(args)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doTask(args []string) {
	if len(args) != 1 {
		log.Fatal("Provide a task ID\n")
		return
	}

	taskNum, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Provide a task ID\n")
		return
	}

	db, err := bolt.Open("task.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// build map of displayed task number to
	taskNumToId := make(map[int][]byte)
	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			fmt.Println("Task list has not been created.")
			return nil
		}

		c := b.Cursor()
		index := 0
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			index++
			taskNumToId[index] = k
		}

		return nil
	}); err != nil {
		log.Fatal(err)
		return
	}

	if taskId, ok := taskNumToId[taskNum]; ok {
		// delete task from DB
		if err := db.Update(func(tx *bolt.Tx) error {
			return tx.Bucket([]byte("tasks")).Delete(taskId)
		}); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		fmt.Printf("Task %d does not exist.\n", taskNum)
		return
	}
}
