package main

import (
	"github.com/fhluo/json2go/cmd"
	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
