package param

import (
	"mjpclab.dev/ghfs/src/goNixArgParser"
	baseParam "mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/serverError"
	"os"
)

var cliCmd = NewCliCmd()

func NewCliCmd() *goNixArgParser.Command {
	cmd := baseParam.NewCliCmd()
	options := cmd.Options()

	// define option
	var err error

	err = options.AddFlagValues("redirects", "--redirect", "", nil, "add rule for http redirect, format <sep><match><sep><replace>[<sep><code>]")
	serverError.CheckFatal(err)

	return cmd
}

func CmdResultsToParams(results []*goNixArgParser.ParseResult) (params []*Param, errs []error) {
	params = make([]*Param, 0, len(results))

	for _, result := range results {
		param := &Param{}

		// redirects
		strRedirects, _ := result.GetStrings("redirects")
		redirects := baseParam.SplitAllKeyValues(strRedirects)
		param.Redirects = make([][3]string, len(redirects))
		for i := range redirects {
			copy(param.Redirects[i][:], redirects[i])
		}

		param.normalize()
		params = append(params, param)
	}

	return
}

func ParseFromCli() (baseParams []*baseParam.Param, params []*Param, printVersion, printHelp bool, errs []error) {
	var es []error
	var cmdResults []*goNixArgParser.ParseResult

	cmdResults, printVersion, printHelp, errs = baseParam.ArgsToCmdResults(cliCmd, os.Args)
	if printVersion || printHelp || len(errs) > 0 {
		return
	}

	baseParams, es = baseParam.CmdResultsToParams(cmdResults)
	errs = append(errs, es...)

	params, es = CmdResultsToParams(cmdResults)
	errs = append(errs, es...)

	return
}

func PrintHelp() {
	cliCmd.OutputHelp(os.Stdout)
}
