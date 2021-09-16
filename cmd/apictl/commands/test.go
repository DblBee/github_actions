package commands

import (
	"fmt"
	"os/exec"

	cli_utils "github.com/dblbee/github_actions/cmd/apictl/utils"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "The test command will run the unit testing suite against the codebase",
		Long:  `The test command will run the unit testing suite against the codebase and generate the code coverage report`,
		Run: func(cmd *cobra.Command, args []string) {
			verboseArg := ""

			if verboseFlag {
				verboseArg = " -v"
			}

			var reportCmd = cli_utils.Command{
				CommandName: "bash",
				CommandArgs: []string{
					"-c",
					"gocov-html ./testResults/coverage.json > ./testResults/coverage.html && open ./testResults/coverage.html",
				},
			}

			var cleanCmds = []cli_utils.Command{
				{
					CommandName: "bash",
					CommandArgs: []string{
						"-c",
						"rm -rf $HOME/go/bin/gocov-html",
					},
				},
				{
					CommandName: "bash",
					CommandArgs: []string{
						"-c",
						"rm -rf $HOME/go/bin/gocov",
					},
				},
				{
					CommandName: "bash",
					CommandArgs: []string{
						"-c",
						"rm -rf $HOME/go/bin/staticcheck",
					},
				},
			}

			var testCmds = []cli_utils.Command{
				{
					CommandName: "go",
					CommandArgs: []string{
						"install",
						"honnef.co/go/tools/cmd/staticcheck@latest",
					},
				},
				{
					CommandName: "go",
					CommandArgs: []string{
						"install",
						"github.com/matm/gocov-html@latest",
					},
				},
				{
					CommandName: "go",
					CommandArgs: []string{
						"install",
						"github.com/axw/gocov/gocov@latest",
					},
				},
				{
					CommandName: "staticcheck",
					CommandArgs: []string{
						".",
					},
				},
				{
					CommandName: "bash",
					CommandArgs: []string{
						"-c",
						fmt.Sprintf("gocov test ./...%s > ./testResults/coverage.json", verboseArg),
					},
				},
			}

			if reportFlag {
				testCmds = append(testCmds, reportCmd)
			}

			testCmds = append(testCmds, cleanCmds...)

			for _, c := range testCmds {
				cmd := exec.Command(c.CommandName, c.CommandArgs...)

				stderr, _ := cmd.StderrPipe()
				stdout, _ := cmd.StdoutPipe()

				cmd.Start()

				cli_utils.OutputLog(stderr)
				cli_utils.OutputLog(stdout)

				cmd.Wait()
			}
		},
	}
	verboseFlag bool
	reportFlag  bool
)

func init() {
	testCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "To generate verbose test results to the console")
	testCmd.Flags().BoolVarP(&reportFlag, "report", "r", false, "To generate a code coverage report")

	rootCmd.AddCommand(testCmd)
}
