package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/inabajunmr/treview/config"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setting for default langage configration.",
	Run: func(cmd *cobra.Command, args []string) {

		allLangs, err := trending.FindLangs()
		if err != nil {
			fmt.Println("Can not get langs from GitHub.")
			os.Exit(1)
		}
		var langs []string

		for {

			flangs, err := filterLang(allLangs)

			prompt := promptui.Select{
				Label: "Select lang",
				Items: flangs,
			}

			_, result, err := prompt.Run()

			if err != nil {
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

	prompt := promptui.Prompt{
		Label: "lang",
	}

	for {

		filter, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil, err
		}

		var fLangs []string
		for _, lang := range allLangs {
			if strings.Index(lang, filter) != -1 {
				fLangs = append(fLangs, lang)
			}
		}

		if len(fLangs) == 0 {
			fmt.Println("No langage matched by " + filter + ".")
			continue
		}

		return fLangs, nil
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
