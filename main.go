package main

import (
	"fmt"
	"go-moco/kit/logger"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
	"golang.org/x/crypto/bcrypt"

	. "github.com/logrusorgru/aurora"
)

func main() {
	app := cli.NewApp()
	app.Name = "Bcrypt cli"
	app.Usage = `generate bcrypt hash from string`
	app.Version = "1.0.0"
	//  Global Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "cost",
			Value: "10",
			Usage: "Cost of bcrypt",
		},
		cli.StringFlag{
			Name:  "pass, p",
			Value: "",
			Usage: "Password to check",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() <= 0 {
			return fmt.Errorf("%v", Red("invalid argument"))
		}
		checkPass := c.GlobalString("pass")
		if checkPass != "" {
			hash := c.Args().Get(0)
			if ComparePasswords(hash, []byte(checkPass)) {
				fmt.Println(Green("MATCH"))
			} else {
				fmt.Println(Red("NOT MATCH"))
			}

		} else {
			cost := c.GlobalInt("cost")
			rawPWD := c.Args().Get(0)
			fmt.Print(HashAndSalt([]byte(rawPWD), cost))
		}
		return nil
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// HashAndSalt --
func HashAndSalt(pwd []byte, cost int) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		logger.McLog.Error(err)
		return ""
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords --
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
