package main

import (
	"fmt"
	"os"
	"strings"
	"log"

	"github.com/urfave/cli"

	lib "./lib"
)

func main() {
	app := cli.NewApp()

	app.Name = "createx"
	app.Usage = "Create tex template app"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "main.tex",
			Usage: "Create tex file named `FILENAME`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "latex",
			Aliases: []string{"l"},
			Usage:   "create latex paper",
			Action:  CreateLatex,
		},
		{
			Name:    "beamer",
			Aliases: []string{"b"},
			Usage:   "Create Slide",
			Action:  CreateSlide,
		},
	}

	app.Run(os.Args)
}

func CreateLatex(c *cli.Context) error {
	filename := CreateFileName(c)
	log.Print("Create latex file")

	if isExist(filename) {
		log.Print("This filename already exists.")
		return nil
	}

	lib.LatexTemplateCopy(filename)

	return nil
}

func CreateSlide(c *cli.Context) error {
	filename := CreateFileName(c)
	fmt.Println("Create beamer file")
	
	if isExist(filename) {
		log.Print("This filename already exists.")
		return nil
	}

	lib.BeamerTemplateCopy(filename)

	return nil
}

func NotTexSuffix(s string) bool {
	return !strings.HasSuffix(s, ".tex")
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func CreateFileName(c *cli.Context) string {
	filename := "main.tex"
	if c.NArg() > 0 {
		filename = c.Args()[0]
	}
	if NotTexSuffix(filename) {
		filename = filename + ".tex"
	}
	return filename
}
