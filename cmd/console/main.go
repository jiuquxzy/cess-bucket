/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package console

import (
	"fmt"
	"os"

	"github.com/CESSProject/cess-bucket/configs"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   configs.Name,
	Short: configs.Description,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("\x1b[%dm[err]\x1b[0m %v\n", 41, err)
		os.Exit(1)
	}
}

// init
func init() {
	rootCmd.AddCommand(
		Command_Version(),
		Command_State(),
		Command_Run(),
		Command_Withdraw(),
	)
	rootCmd.PersistentFlags().StringP("config", "c", "conf.yaml", "custom configuration file")
	rootCmd.PersistentFlags().StringP("rpc", "", "ws://1127.0.0.1:9948/", "rpc endpoint")
	rootCmd.PersistentFlags().StringP("ws", "", "/", "workspace")
	rootCmd.PersistentFlags().StringP("ip", "", "0.0.0.0", "listening ip address")
	rootCmd.PersistentFlags().StringP("income", "", "", "income account")
	rootCmd.PersistentFlags().IntP("port", "", 15000, "listening port")
	rootCmd.PersistentFlags().Uint64P("space", "", 1000, "maximum space used (GiB)")
}

func Command_Version() *cobra.Command {
	cc := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(configs.Name + " " + configs.Version)
			os.Exit(0)
		},
		DisableFlagsInUseLine: true,
	}
	return cc
}

func Command_State() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "state",
		Short:                 "Query storage miner information",
		Run:                   Command_State_Runfunc,
		DisableFlagsInUseLine: true,
	}
	return cc
}

func Command_Run() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "run",
		Short:                 "Automatically register and run",
		Run:                   runCmd,
		DisableFlagsInUseLine: true,
	}
	return cc
}

func Command_Withdraw() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "withdraw",
		Short:                 "withdraw stakes",
		Run:                   Command_Withdraw_Runfunc,
		DisableFlagsInUseLine: true,
	}
	return cc
}