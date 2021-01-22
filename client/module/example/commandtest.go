package main

import (
	"Faygo/client/module/command"
	"log"
)

func main() {
	_, out, err := command.NewCommand().Exec("ls -ll /")
	if err != nil {
		log.Panic(err)
	}

	log.Println(out)

}
