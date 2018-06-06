package utils

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	TendermintRpcAddrFlag = cli.StringFlag{
		Name:  "tendermint_addr",
		Value: "tcp://localhost:46657",
		Usage: "This is the address that minter will use to connect to the tendermint core node. Please provide a port.",
	}
)

func GetMinterHome() string {
	home := os.Getenv("MINTERHOME")

	if home == "" {
		home = "~/.minter"
	}

	return home
}
