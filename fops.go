package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "File ops"
	app.Usage = ""
	app.Description = ""
	app.Version = "v1.0.0"
	app.UsageText = "fops [flags]\n   fops [command] "
	var filename string
	var md5Flag bool
	var sha1Flag bool
	var sha256Flag bool
	app.Commands = []cli.Command{
		{
			Name:      "linecount",
			Usage:     "Print linecount of file",
			UsageText: "fops lincecount [flag]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file,f",
					Usage:       "the input file",
					Destination: &filename,
				},
			},
			Action: func(c *cli.Context) error {
				if filename != "" {
					fileResult, error := os.Open(filename)
					defer fileResult.Close()

					// if file exists
					if stat, err := os.Stat(filename); err == nil {
						// if file is a directory
						if stat.IsDir() {
							fmt.Println("error: Expected File got directory " + filename)
						} else if error == nil {
							// Create new Scanner.
							scanner := bufio.NewScanner(fileResult)
							result := 0
							// Use Scan.
							for scanner.Scan() {
								result++
							}
							fmt.Println(result)
						} else {
							fmt.Println(error)
						}
						// if file does *not* exist
					} else if os.IsNotExist(err) {

						fmt.Println("error: No such file " + filename)
					} else {
						// Schrodinger: file may or may not exist. See err for details.

						// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

					}
				}

				return nil
			},
		},
		{
			Name:  "checksum",
			Usage: "Print checksum of file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file,f",
					Usage:       "the input file",
					Destination: &filename,
				},
				cli.BoolFlag{
					Name:        "md5",
					Usage:       "generate checksum md5",
					Destination: &md5Flag,
				},
				cli.BoolFlag{
					Name:        "sha1",
					Usage:       "generate checksum sha1",
					Destination: &sha1Flag,
				},
				cli.BoolFlag{
					Name:        "sha256",
					Usage:       "generate checksum sha256",
					Destination: &sha256Flag,
				},
			},
			Action: func(c *cli.Context) error {
				if filename != "" {

					data := []byte(filename)
					if c.Bool("md5") {
						fmt.Printf("%x", md5.Sum(data))
					}
					if c.Bool("sha1") {
						fmt.Printf("%x", sha1.Sum(data))
					}
					if c.Bool("sha256") {
						fmt.Printf("%x", sha256.Sum256(data))
					}
				}

				return nil
			},
		},
		{
			Name:  "version",
			Usage: "Show the version info",
			Action: func(c *cli.Context) error {
				fmt.Println(app.Version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
