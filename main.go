package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "getQiitaInfo"
	app.Usage = "get various info through qiita api"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "myinfo",
			Usage: "get user info",
			Action: func(c *cli.Context) error {
				accessToken := os.Getenv("QIITA_ACCESS_TOKEN")
				if accessToken == "" {
					fmt.Println("Please set $QIITA_ACCESS_TOKEN")
				}

				user := os.Getenv("QIITA_USER")
				if user == "" {
					fmt.Println("Please set $QIITA_USER")
				}

				info, err := cmd.FetchQiitaData(accessToken, user)
				if err != nil {
					log.Fatal(err)
				}

				cmd.OutputQiitaInfo(info)
				return nil
			},
		},
	}

	app.RunAndExitOnError()
}
