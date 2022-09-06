package middleware

import (
	"errors"
	"mjpclab.dev/ehfs/src/param"
	"mjpclab.dev/ghfs/src/middleware"
	baseParam "mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/serverError"
)

var errInvalidParamValue = errors.New("invalid param value")
var errParamCountNotMatch = errors.New("base-param count is not equal to param count")

func ParamToMiddlewares(baseParam *baseParam.Param, param *param.Param) (preMids, postMids []middleware.Middleware, errs []error) {
	preMids = make([]middleware.Middleware, 0, len(param.Rewrites)+
		len(param.Redirects)+
		len(param.Proxies)+
		len(param.Returns),
	)
	postMids = make([]middleware.Middleware, 0, len(param.StatusPages))
	statusPageMids := make([]middleware.Middleware, 0, len(param.StatusPages))

	// status pages
	for i := range param.StatusPages {
		mid, err := getStatusPageMiddleware(param.StatusPages[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			statusPageMids = append(statusPageMids, mid)
			postMids = append(postMids, mid)
		}
	}

	// rewrites
	for i := range param.Rewrites {
		mid, err := getRewriteMiddleware(param.Rewrites[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			preMids = append(preMids, mid)
		}
	}

	// redirects
	for i := range param.Redirects {
		mid, err := getRedirectMiddleware(param.Redirects[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			preMids = append(preMids, mid)
		}
	}

	// proxies
	for i := range param.Proxies {
		mid, err := getProxyMiddleware(param.Proxies[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			preMids = append(preMids, mid)
		}
	}

	// returns
	for i := range param.Returns {
		mid, err := getReturnStatusMiddleware(param.Returns[i], statusPageMids)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			preMids = append(preMids, mid)
		}
	}

	return
}

func ApplyMiddlewares(baseParams []*baseParam.Param, params []*param.Param) (errs []error) {
	if len(baseParams) != len(params) {
		return []error{errParamCountNotMatch}
	}

	for i := range baseParams {
		var es []error
		baseParams[i].PreMiddlewares, baseParams[i].PostMiddlewares, es = ParamToMiddlewares(baseParams[i], params[i])
		errs = append(errs, es...)
	}

	return
}
