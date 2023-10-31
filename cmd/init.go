/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"kubecx/app"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{""},
	Short:   "初始化 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx()
		err := kc.InitMainConfig()
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
