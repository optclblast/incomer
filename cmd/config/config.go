/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"encoding/json"
	"incomer/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var CustomDBPath string

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   RunConfig,
}

func RunConfig(cmd *cobra.Command, args []string) {
	statefiledir, err := os.Getwd()
	if err != nil {
		panic("error can't parse statfile dir path: " + err.Error())
	}

	log.Println(statefiledir)

	stateFileData, err := os.ReadFile(statefiledir + "/state.json")
	if err != nil {
		panic("error can't read from statfile.json: " + err.Error())
	}

	log.Println(stateFileData)

	var stateFile config.StateJSON
	err = json.Unmarshal(stateFileData, &stateFile)
	if err != nil {
		panic("error can't parde from statfile.json data: " + err.Error())
	}

	log.Println(statefiledir + CustomDBPath)

	if CustomDBPath != "" {
		stateFile.CustomDBpath = statefiledir + CustomDBPath
		log.Println(statefiledir + CustomDBPath)
		newStateFileData, err := json.Marshal(stateFile)
		log.Println(string(newStateFileData))
		if err != nil {
			panic("error can't prepare new dbpath to writing to state file: " + err.Error())
		}

		err = os.WriteFile(statefiledir+"/state.json", newStateFileData, 0644)
		if err != nil {
			panic("error can't write new dbpath to state file: " + err.Error())
		}

	}
}

func init() {
	ConfigCmd.Flags().StringVarP(
		&CustomDBPath,
		"dbpath",
		"d",
		"",
		"If provided, sets a path to already existing database, that will be used",
	)
}
