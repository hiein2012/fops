package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

//CheckFileExists : check file exists or not
func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckFileIsBinary : if given file is binary data , return true.
// Otherwise(vaild MIME type) , return false
func CheckFileIsBinary(filename string) bool {

	fileContent, _ := ioutil.ReadFile(filename)
	byteFileContent := []byte(fileContent)

	// DetectContentType always returns a valid MIME type
	// If it is binary data , http.DetectContentType() returns "application/octet-stream"
	contentType := http.DetectContentType(byteFileContent)

	if contentType == "application/octet-stream" {
		return true
	}
	return false
}

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
				// check given file exists or not
				// If file does not exists ,  return
				if CheckFileExists(filename) == false {
					fmt.Println("error: No such file " + filename)
					return nil
				}

				// Already checked file given file exists
				// check the content of given file is binary type or not
				// if it is binary type , return
				if CheckFileIsBinary(filename) == true {
					fmt.Println("error: Cannot do linecount for binary file " + filename)
					return nil
				}

				// Already checked file given file exists , and not a binary file type
				// use FileMode to check file/dir
				fi, _ := os.Lstat(filename)

				switch mode := fi.Mode(); {
				// if given file is directory
				case mode.IsDir():
					fmt.Println("error: Expected File got directory " + filename)

				// if given file is regular file
				case mode.IsRegular():
					// start to line count
					fileResult, error := os.Open(filename)
					if error != nil {
						log.Fatal(error)
					}
					defer fileResult.Close()
					scanner := bufio.NewScanner(fileResult)
					lineCountNumber := 0
					for scanner.Scan() {
						lineCountNumber++
					}
					fmt.Println(lineCountNumber)

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
				// check given file exists or not
				// If file does not exists ,  return
				if CheckFileExists(filename) == false {
					fmt.Println("error: No such file " + filename)
					return nil
				}

				// Already checked file given file exists
				// therefore read the file content and turn to byte type
				fileContent, _ := ioutil.ReadFile(filename)
				byteFileContent := []byte(fileContent)

				// print specific checksum depends on given flag
				if c.Bool("md5") {
					fmt.Printf("%x", md5.Sum(byteFileContent))
				}
				if c.Bool("sha1") {
					fmt.Printf("%x", sha1.Sum(byteFileContent))
				}
				if c.Bool("sha256") {
					fmt.Printf("%x", sha256.Sum256(byteFileContent))
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
