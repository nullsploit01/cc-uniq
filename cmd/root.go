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
	Use:   "ccuniq [flags] [file]",
	Short: "Filter and process duplicate lines from input.",
	Long: `CC-Uniq is a command-line utility that enhances the functionality of traditional Unix 'uniq' by supporting direct output to a file. It reads input, processes duplicate lines, and outputs the results either to standard output or directly to a specified output file.
CC-Uniq can handle input from files or standard input streams, and it offers various options for managing duplicate lines:

Examples of using ccuniq:
  # Directly invoke ccuniq on a file, outputting to standard output
  ccuniq test.txt

  # Use ccuniq after a cat command, taking input from a pipe
  cat test.txt | ccuniq -

  # Direct output to a file using ccuniq's built-in output feature
  cat test.txt | ccuniq - out.txt

  # Count occurrences of each line and direct the output to a file
  ccuniq -c test.txt out.txt

  # Display only duplicate lines, outputting to a file
  ccuniq -d test.txt out.txt

  # Display only unique lines, directing the output to a specific file
  ccuniq -u test.txt out.txt

CC-Uniq supports several flags to customize its behavior:
  -c, --count        Count occurrences of each line and print them alongside the lines.
  -d, --repeated     Output only lines that appear more than once.
  -u, --unique       Output lines that appear exactly once.`,

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
