package env

import (
	"os"

	"github.com/joho/godotenv"
)

var envLoadedBefore = false

func lazyInit() {
	if envLoadedBefore {
		return
	}

	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
	envLoadedBefore = true
}
