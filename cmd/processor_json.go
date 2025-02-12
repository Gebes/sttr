// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var json_flag_i bool

func init() {	
	jsonCmd.Flags().BoolVarP(&json_flag_i, "indent", "i", false, "Indent the output (prettyprint)")
	rootCmd.AddCommand(jsonCmd)
}

var jsonCmd = &cobra.Command{
	Use:     "json",
	Short:   "Format your text as JSON",
	Aliases: []string{},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in, err = ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := ioutil.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.FormatJSON{}
		flags = append(flags, processors.Flag{Short: "i", Value: json_flag_i})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
