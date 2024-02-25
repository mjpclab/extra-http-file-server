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

	err = options.AddFlagValues("ipallows", "--ip-allow", "", nil, "specify allowed client IP, rests are denied")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("ipallowfiles", "--ip-allow-file", "", nil, "specify allowed client IP from files, rests are denied")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("ipdenies", "--ip-deny", "", nil, "specify denied client IP, rests are allowed if no allow list")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("ipdenyfiles", "--ip-deny-file", "", nil, "specify denied client IP from files, rests are allowed if no allow list")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("rewritehosts", "--rewrite-host", "", nil, "add rule to replace request URL by host+request_URL, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("rewritehostspost", "--rewrite-host-post", "", nil, "add rule to replace request URL by host+request_URL after redirects, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("rewritehostsend", "--rewrite-host-end", "", nil, "add rule to replace request URL by host+request_URL, and skip further actions, format <sep><match><sep><replace>")
	serverError.CheckFatal(err)

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

	err = options.AddFlagValues("headeradds", "--header-add", "", nil, "add response header, format <sep><match><sep><name><sep><value>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("headersets", "--header-set", "", nil, "set response header, format <sep><match><sep><name><sep><value>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("tostatuses", "--to-status", "", nil, "add rule to move to status code after ghfs internal process, format <sep><match><sep><code>")
	serverError.CheckFatal(err)

	err = options.AddFlagValues("statuspages", "--status-page", "", nil, "set page file for specific http status code, format <sep><status><sep><fs-path>")
	serverError.CheckFatal(err)

	err = options.AddFlag("gzipstatic", "--gzip-static", "EHFS_GZIP_STATIC", "look for request-file.gz on file system to output compressed content")
	serverError.CheckFatal(err)

	return cmd
}

func CmdResultsToParams(results []*goNixArgParser.ParseResult) (params []*Param, errs []error) {
	params = make([]*Param, 0, len(results))

	for _, result := range results {
		param := &Param{}

		// IP allows/denies
		param.IPAllows, _ = result.GetStrings("ipallows")
		param.IPAllowFiles, _ = result.GetStrings("ipallowfiles")
		param.IPDenies, _ = result.GetStrings("ipdenies")
		param.IPDenyFiles, _ = result.GetStrings("ipdenyfiles")

		// rewrite hosts
		rewriteHosts, _ := result.GetStrings("rewritehosts")
		param.RewriteHosts = baseParam.SplitAllKeyValue(rewriteHosts)

		// rewrite hosts post
		rewritesHostsPost, _ := result.GetStrings("rewritehostspost")
		param.RewriteHostsPost = baseParam.SplitAllKeyValue(rewritesHostsPost)

		// rewrite hosts end
		rewriteHostsEnd, _ := result.GetStrings("rewritehostsend")
		param.RewriteHostsEnd = baseParam.SplitAllKeyValue(rewriteHostsEnd)

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
		redirects, _ := result.GetStrings("redirects")
		param.Redirects = toString3s(redirects)

		// proxies
		proxies, _ := result.GetStrings("proxies")
		param.Proxies = baseParam.SplitAllKeyValue(proxies)

		// returns
		returns, _ := result.GetStrings("returns")
		param.Returns = baseParam.SplitAllKeyValue(returns)

		// headers
		headerAdds, _ := result.GetStrings("headeradds")
		param.HeaderAdds = toString3s(headerAdds)
		headerSets, _ := result.GetStrings("headersets")
		param.HeaderSets = toString3s(headerSets)

		// to statuses
		toStatuses, _ := result.GetStrings("tostatuses")
		param.ToStatuses = baseParam.SplitAllKeyValue(toStatuses)

		// status pages
		statusPages, _ := result.GetStrings("statuspages")
		param.StatusPages = baseParam.SplitAllKeyValue(statusPages)

		// gzip statics
		param.GzipStatic = result.HasKey("gzipstatic")

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
