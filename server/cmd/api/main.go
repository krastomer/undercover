package main

import "github.com/krastomer/undercover/server/internal/di"

func main() {
	ctn, err := di.NewContainer()
	if err != nil {
		panic(err)
	}

	ctn.Start()
}
