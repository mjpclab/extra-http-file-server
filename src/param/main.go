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
	// value: [match, replace]
	Proxies [][2]string
	// value: [match, code, file?]
	Returns [][3]string

	// value: [code, file]
	StatusPages [][2]string
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

	// returns
	returns := make([][3]string, 0, len(param.Returns))
	for i := range param.Returns {
		if len(param.Returns[i][0]) == 0 || len(param.Returns[i][1]) == 0 {
			continue
		}
		code, err := strconv.Atoi(param.Returns[i][1])
		if err != nil {
			continue
		}
		param.Returns[i][1] = strconv.Itoa(code)

		returns = append(returns, param.Returns[i])
	}
	param.Returns = returns
}
