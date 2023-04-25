/*
Copyright © 2023 prae-api
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/prae-api/stringer/pkg"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stringer",
	Short: "stringer: a simple CLI to transform and inspect string",
	Long: `stringer is a super fancy CLI (jk)

	One can use stringer to modify or inspect strings from the terminal	
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverse a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := pkg.Reverse(args[0])
		fmt.Printf("arg[0] is %v\n", args[0])
		fmt.Println(res)
	},
}

var onlyDigits bool //NoOptDefVal = true by default (having option is somthing like this -d=false), refer to doc https://github.com/spf13/pflag/blob/v1.0.5/bool.go#L54

var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		fmt.Printf("onlyDigit is: %t\n", onlyDigits)
		res, kind := pkg.Inspect(i, onlyDigits)

		plurals := "s"
		if res == 1 {
			plurals = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, plurals)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while excuting your CLI '%s", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stringer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(reverseCmd)
	rootCmd.AddCommand(inspectCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
