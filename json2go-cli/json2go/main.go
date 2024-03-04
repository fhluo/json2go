package main

import (
	"log"

	"github.com/fhluo/json2go/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
