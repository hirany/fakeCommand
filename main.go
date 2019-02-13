package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fake"
	app.Usage = "dammy command"
	app.Action = func(c *cli.Context) {
		err := fake()
		if err != nil {
			log.Fatal(err)
		}
	}
	app.Run(os.Args)
}

func fake() error {
	var str string
	err := showDir()
	if err != nil {
		return err
	}
	for {
		str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		if str == "fake\r\n" {
			return nil
		}
		showDir()
	}
	return nil
}

func showDir() error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("PS %s> ", curDir)
	return nil
}
