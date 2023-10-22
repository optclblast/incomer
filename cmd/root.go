/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"incomer/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "incomer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//fmt.Println(viper.)

	historyFile := viper.GetString(config.HISTORY_FILE_PATH_ARG)
	if historyFile == "" {
		panic("fatal error: can't get history file path")
	}

	//rootCmd.PersistentFlags().StringVar()
}

func initConfig() {
	usrHome, err := os.UserHomeDir()
	if err != nil {
		panic("fatal error: " + err.Error())
	}

	config.USER_HOME = usrHome

	if _, err := os.Open(usrHome + "/" + config.CONFIG_FILE); errors.Is(err, os.ErrNotExist) {
		err = config.CreateDefaultConfigFile()
		if err != nil {
			panic("fatal error: " + err.Error())
		}
	}

	viper.SetConfigFile(config.CONFIG_FILE)
}
