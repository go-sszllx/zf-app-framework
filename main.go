package main

import (
	"fmt"
	"framework/app"
)

func main() {
	app.AddAPPStartHook(func() error {
		// init DB
		fmt.Println("init DB")
		return nil
	}, func() error {
		// other init actions
		fmt.Println("init other actions")
		return nil
	})

	app.Run("mono", "data", "/tmp/csv")
}
