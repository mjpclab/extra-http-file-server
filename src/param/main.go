package param

import (
	"mjpclab.dev/ehfs/src/util"
	baseParam "mjpclab.dev/ghfs/src/param"
	"strconv"
)

type Param struct {
	IPAllows     []string
	IPAllowFiles []string
	IPDenies     []string
	IPDenyFiles  []string
	// value: [match, replace]
	RewriteHosts     [][2]string
	RewriteHostsPost [][2]string
	RewriteHostsEnd  [][2]string
	Rewrites         [][2]string
	RewritesPost     [][2]string
	RewritesEnd      [][2]string
	// value: [match, replace, code?]
	Redirects [][3]string
	// value: [match, replace]
	Proxies [][2]string
	// value: [match, code]
	Returns [][2]string

	// value: [match, name, value]
	HeaderAdds [][3]string
	// value: [match, name, value]
	HeaderSets [][3]string
	// value: [match, code]
	ToStatuses [][2]string
	// value: [code, file]
	StatusPages [][2]string
}

func (param *Param) normalize() {
	param.IPAllows = util.Filter(param.IPAllows, nonEmptyString)
	param.IPAllowFiles = util.Filter(param.IPAllowFiles, nonEmptyString)
	param.IPDenies = util.Filter(param.IPDenies, nonEmptyString)
	param.IPDenyFiles = util.Filter(param.IPDenyFiles, nonEmptyString)

	param.Rewrites = util.Filter(param.Rewrites, nonEmptyKeyString2)
	param.RewritesPost = util.Filter(param.RewritesPost, nonEmptyKeyString2)
	param.RewritesEnd = util.Filter(param.RewritesEnd, nonEmptyKeyString2)
	param.Redirects = util.Filter(param.Redirects, nonEmptyKeyString3)
	const defaultRedirectCode = "301"
	redirects := make([][3]string, 0, len(param.Redirects))
	for i := range param.Redirects {
		code, err := strconv.Atoi(param.Redirects[i][2])
		if err != nil {
			param.Redirects[i][2] = defaultRedirectCode
		} else {
			param.Redirects[i][2] = strconv.Itoa(baseParam.NormalizeRedirectCode(code))
		}

		redirects = append(redirects, param.Redirects[i])
	}
	param.Redirects = redirects

	param.Proxies = util.Filter(param.Proxies, nonEmptyKeyString2)
	param.Returns = util.Filter(param.Returns, nonEmptyKeyString2)

	param.HeaderAdds = util.Filter(param.HeaderAdds, nonEmptyString3)
	param.HeaderSets = util.Filter(param.HeaderSets, nonEmptyString3)
	param.ToStatuses = util.Filter(param.ToStatuses, nonEmptyKeyString2)
	param.StatusPages = util.Filter(param.StatusPages, nonEmptyKeyString2)

}
