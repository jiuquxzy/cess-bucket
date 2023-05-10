/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package console

import (
	"os"

	"github.com/CESSProject/cess-bucket/configs"
	"github.com/CESSProject/cess-bucket/node"
	sdkgo "github.com/CESSProject/sdk-go"
	"github.com/CESSProject/sdk-go/core/client"
	"github.com/spf13/cobra"
)

// Withdraw the staking
func Command_Withdraw_Runfunc(cmd *cobra.Command, args []string) {
	var (
		ok  bool
		err error
		n   = node.New()
	)

	// Build profile instances
	n.Cfg, err = buildConfigFile(cmd, "", 0)
	if err != nil {
		logERR(err.Error())
		os.Exit(1)
	}

	//Build client
	cli, err := sdkgo.New(
		configs.Name,
		sdkgo.ConnectRpcAddrs(n.Cfg.GetRpcAddr()),
		sdkgo.ListenPort(n.Cfg.GetServicePort()),
		sdkgo.Workspace(n.Cfg.GetWorkspace()),
		sdkgo.Mnemonic(n.Cfg.GetMnemonic()),
		sdkgo.TransactionTimeout(configs.TimeToWaitEvent),
	)
	if err != nil {
		logERR(err.Error())
		os.Exit(1)
	}

	n.Cli, ok = cli.(*client.Cli)
	if !ok {
		logERR("Invalid client type")
		os.Exit(1)
	}

	txhash, err := n.Cli.Withdraw()
	if err != nil {
		if txhash == "" {
			logERR(err.Error())
			os.Exit(1)
		}
		logWARN(txhash)
		os.Exit(0)
	}

	logOK(txhash)
	os.Exit(0)
}
