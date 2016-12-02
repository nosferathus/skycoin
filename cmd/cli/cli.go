package main

import (
	"fmt"
	"os"

	"strings"

	skycli "github.com/skycoin/skycoin/src/api/cli"
	"github.com/skycoin/skycoin/src/util"
	"github.com/urfave/cli"
)

var (
	commandHelpTemplate = `USAGE:
		{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{if .Category}}
		
CATEGORY:
		{{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
		{{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:
		{{range .VisibleFlags}}{{.}}
		{{end}}{{end}}
	`
)

func main() {
	// get rpc address from env
	rpcAddr := os.Getenv("RPC_ADDR")
	if rpcAddr == "" {
		rpcAddr = "127.0.0.1:6422"
	}

	// get wallet dir from env
	wltDir := os.Getenv("WALLET_DIR")
	if wltDir == "" {
		home := util.UserHome()
		wltDir = home + "/.skycoin/wallets"
	}

	wltName := os.Getenv("WALLET_NAME")
	if wltName == "" {
		wltName = "skycoin_cli.wlt"
	} else {
		if !strings.HasSuffix(wltName, ".wlt") {
			fmt.Println("value of 'WALLET_NAME' env is not correct, must has .wlt extenstion")
			return
		}
	}

	// init the skycli
	skycli.Init(skycli.RPCAddr(rpcAddr),
		skycli.WalletDir(wltDir),
		skycli.DefaultWltName(wltName))

	cli.SubcommandHelpTemplate = commandHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help,h",
		Usage: "show help, can also be used to show subcommand help",
	}

	app := cli.NewApp()
	app.Usage = "the skycoin command line interface"
	app.Version = "0.1"
	app.Commands = skycli.Commands
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
