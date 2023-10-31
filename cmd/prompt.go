/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kubecx/app"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:     "prompt",
	Aliases: []string{},
	Short:   "获取当前 [kubectl context] 的 prompt 信息,用于配置 shell 提示",
	Long: `配置方式为在shell配置文件增加: 
PROMPT='$(kubecx prompt)'$PROMPT`,
	Run: func(cmd *cobra.Command, args []string) {
		kc := kubectx.NewKubeCtx()
		err := kc.Prompt(promptCmdOptions)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	},
}

var promptCmdOptions = new(kubectx.PromptOptions)

func init() {
	rootCmd.AddCommand(promptCmd)
	promptCmd.Flags().StringVarP(&promptCmdOptions.Prefix, "prefix", "", "[", "输出前缀")
	promptCmd.Flags().StringVarP(&promptCmdOptions.Icon, "icon", "", "⎈", "标识")
	promptCmd.Flags().StringVarP(&promptCmdOptions.Separator, "separator", "", "|", "分离符号")
	promptCmd.Flags().StringVarP(&promptCmdOptions.Divider, "divider", "", ":", "内容间隔符号")
	promptCmd.Flags().StringVarP(&promptCmdOptions.Suffix, "suffix", "", "]", "输出后缀")
}
