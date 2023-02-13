package main

import app2 "dskit-examples/ring_default/app"

func main() {
	app, err := app2.NewApp()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		app2.Logger.Log("msg", "application is terminated", "err", err)
	}
}
