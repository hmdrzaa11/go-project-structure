package main

import (
	"github.com/joho/godotenv"
)

func main() {
	//load our environemt variuable
	if err := godotenv.Load(); err != nil {
		panic("no .env found!")
	}
}
