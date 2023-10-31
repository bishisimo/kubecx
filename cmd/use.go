/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"kubecx/app"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:     "use",
	Aliases: []string{"set"},
	Short:   "使用指定 [kubectl context] 资源",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && useOptions.Target == "" {
			_ = cmd.Help()
			return
		}
		if len(args) > 0 && useOptions.Target == "" {
			useOptions.Target = args[0]
		}
		kc := kubectx.NewKubeCtx()
		err := kc.Use(useOptions)
		if err != nil {
			log.Warn().Err(err).Send()
			return
		}
	},
}

var useOptions = new(kubectx.UseOptions)

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().StringVarP(&useOptions.Target, "target", "t", "", "使用指定的context资源")
}
