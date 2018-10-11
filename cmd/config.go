package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/inabajunmr/treview/config"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setting for default langage configration.",
	Run: func(cmd *cobra.Command, args []string) {

		allLangs := trending.FindLangs()

		var langs []string

		for {

			flangs, err := filterLang(allLangs)
			if err != nil {
				os.Exit(1)
			}

			if flangs == nil {
				break
			}

			flangs = append(flangs, ".exit")

			var result string
			prompt := &survey.Select{
				Message: "Choose a lang:",
				Options: flangs,
			}
			err = survey.AskOne(prompt, &result, nil)

			if err != nil {
				break
			}

			if result == ".exit" {
				break
			}

			langs = append(langs, result)
			// TODO remove result
			fmt.Printf("You choose %q\n", result)
		}

		usr, err := user.Current()
		if err != nil {
			os.Exit(1)
		}

		path := usr.HomeDir + "/.treview"
		cpath := path + "/.config"
		config.SetLangs(cpath, langs)
		fmt.Println("You choose ", langs)

	},
}

func filterLang(allLangs []string) ([]string, error) {
	prompt := &survey.Input{
		Message: "lang",
	}

	for {
		var l string
		err := survey.AskOne(prompt, &l, nil)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil, err
		}

		if l == "" {
			return nil, nil
		}

		var fLangs []string
		for _, lang := range allLangs {
			if strings.Contains(lang, l) {
				fLangs = append(fLangs, lang)
			}
		}

		if len(fLangs) == 0 {
			fmt.Println("No langage matched by " + l + ".")
			continue
		}

		return fLangs, nil
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
