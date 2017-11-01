package facade

import (
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/reldemo"
)

const (
	//Name is applicatin name
	Name = "godump"
	//Version is version for applicatin
	Version = "v0.1.1"
)

//ExitCode is OS exit code enumeration class
type ExitCode int

const (
	//Normal is OS exit code "normal"
	Normal ExitCode = iota
	//Abnormal is OS exit code "abnormal"
	Abnormal
)

//Int convert integer value
func (c ExitCode) Int() int {
	return int(c)
}

//Stringer method
func (c ExitCode) String() string {
	switch c {
	case Normal:
		return "normal end"
	case Abnormal:
		return "abnormal end"
	default:
		return "unknown"
	}
}

var (
	cui = gocli.NewUI()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: Name + " [flags] [binary file]",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		reader := cui.Reader()
		if len(args) > 0 {
			file, err2 := os.Open(args[0]) //args[0] is maybe file path
			if err2 != nil {
				return err
			}
			defer file.Close()
			reader = file
		}
		result, err := reldemo.DumpBytes(reader, name)
		if err != nil {
			return err
		}
		cui.WriteFrom(result)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ui *gocli.UI) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = Abnormal
		}
	}()

	//execution
	exit = Normal
	cui = ui
	if err := rootCmd.Execute(); err != nil {
		exit = Abnormal
	}
	return
}

func init() {
	rootCmd.Flags().StringP("name", "n", "dumpList", "value name")
}
