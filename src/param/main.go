package param

import (
	baseParam "mjpclab.dev/ghfs/src/param"
	"strconv"
)

type Param struct {
	// value: [match, replace]
	Rewrites [][2]string
	// value: [match, replace, code?]
	Redirects [][3]string
}

func (param *Param) normalize() {
	// redirects
	const defaultRedirectCode = "301"
	redirects := make([][3]string, 0, len(param.Redirects))
	for i := range param.Redirects {
		if len(param.Redirects[i][0]) == 0 || len(param.Redirects[i][1]) == 0 {
			continue
		}

		code, err := strconv.Atoi(param.Redirects[i][2])
		if err != nil {
			param.Redirects[i][2] = defaultRedirectCode
		} else {
			param.Redirects[i][2] = strconv.Itoa(baseParam.NormalizeRedirectCode(code))
		}

		redirects = append(redirects, param.Redirects[i])
	}
	param.Redirects = redirects
}
