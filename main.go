package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func analyse(args []string) {
  for _, arg := range args {
    fmt.Println("Analysing...",arg)
  }

}
func main() {
  app := cli.NewApp()
  app.Name = "Behind"
  app.Usage = "know how behind your packages are"
  app.Action = func(c *cli.Context) error {
    fmt.Println("Hello Devs")
    if !c.Args().Present() {
        return cli.NewExitError("No input files to analyse", 1)
    }

    args := c.Args()
    analyse(args)
    return nil
  }

  app.Run(os.Args)
}
