package cmd

import (
	"os"

	"github.com/nullsploit01/cc-uniq/internal"
	"github.com/spf13/cobra"
)

var withCount bool
var onlyRepeated bool
var onlyUnique bool

var rootCmd = &cobra.Command{
	Use:   "cc-uniq",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var inputFile *os.File
		var outputFileName string

		if len(args) < 1 {
			cmd.Usage()
			return
		}

		if args[0] == "-" {
			fileInfo, _ := os.Stdin.Stat()
			if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
				inputFile = os.Stdin
			}
		} else {
			file, err := os.Open(args[0])
			if err != nil {
				cmd.ErrOrStderr().Write([]byte(err.Error()))
				os.Exit(1)
			}

			inputFile = file
		}
		defer inputFile.Close()

		if len(args) >= 2 {
			outputFileName = args[1]
		}

		u := internal.NewUniq(cmd)
		if err := u.PrintUniqueLinesFromFile(inputFile, outputFileName, withCount, onlyRepeated, onlyUnique); err != nil {
			cmd.ErrOrStderr().Write([]byte(err.Error()))
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&withCount, "count", "c", false, "Count of number of lines")
	rootCmd.Flags().BoolVarP(&onlyRepeated, "repeated", "d", false, "Print only repeated lines")
	rootCmd.Flags().BoolVarP(&onlyUnique, "unique", "u", false, "Print only unique lines")
}
