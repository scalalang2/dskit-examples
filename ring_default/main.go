package main

import "dskit-examples/ring_default/app"

func main() {
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	if err := application.Run(); err != nil {
		app.Logger.Log("msg", "application is terminated", "err", err)
	}
}
