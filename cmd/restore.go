/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"kubecx/app"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{""},
	Short:   "恢复 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx()
		err := kc.Restore()
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
