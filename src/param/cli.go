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

	err = options.AddFlagValues("rewrites", "--rewrite", "", nil, "add rule to replace request URL, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("rewritespost", "--rewrite-post", "", nil, "add rule to replace request URL after redirects, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("rewritesend", "--rewrite-end", "", nil, "add rule to replace request URL, and skip further actions, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("redirects", "--redirect", "", nil, "add rule for http redirect, format <sep><match><sep><replace>[<sep><code>]")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("proxies", "--proxy", "", nil, "add rule to proxy request URL, format <sep><match><sep><target>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("returns", "--return", "", nil, "add rule to return status code, format <sep><match><sep><code>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("statuspages", "--status-page", "", nil, "set page file for specific http status code, format <sep><status><sep><fs-path>")
	serverError.CheckFatal(err)

	return cmd
}

func CmdResultsToParams(results []*goNixArgParser.ParseResult) (params []*Param, errs []error) {
	params = make([]*Param, 0, len(results))

	for _, result := range results {
		param := &Param{}

		// rewrites
		rewrites, _ := result.GetStrings("rewrites")
		param.Rewrites = baseParam.SplitAllKeyValue(rewrites)

		// rewrites post
		rewritesPost, _ := result.GetStrings("rewritespost")
		param.RewritesPost = baseParam.SplitAllKeyValue(rewritesPost)

		// rewrites end
		rewritesEnd, _ := result.GetStrings("rewritesend")
		param.RewritesEnd = baseParam.SplitAllKeyValue(rewritesEnd)

		// redirects
		strRedirects, _ := result.GetStrings("redirects")
		redirects := baseParam.SplitAllKeyValues(strRedirects)
		param.Redirects = make([][3]string, len(redirects))
		for i := range redirects {
			copy(param.Redirects[i][:], redirects[i])
		}

		// proxies
		proxies, _ := result.GetStrings("proxies")
		param.Proxies = baseParam.SplitAllKeyValue(proxies)

		// returns
		returns, _ := result.GetStrings("returns")
		param.Returns = baseParam.SplitAllKeyValue(returns)

		// status pages
		statusPages, _ := result.GetStrings("statuspages")
		param.StatusPages = baseParam.SplitAllKeyValue(statusPages)

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
