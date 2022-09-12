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
	// status pages
	statusPageMids := make([]middleware.Middleware, 0, len(param.StatusPages))
	for i := range param.StatusPages {
		mid, err := getStatusPageMiddleware(param.StatusPages[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			statusPageMids = append(statusPageMids, mid)
		}
	}

	// rewrites
	rewriteMids := make([]middleware.Middleware, 0, len(param.Rewrites))
	for i := range param.Rewrites {
		mid, err := getRewriteMiddleware(param.Rewrites[i], middleware.GoNext)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteMids = append(rewriteMids, mid)
		}
	}

	// rewrites end
	rewriteEndMids := make([]middleware.Middleware, 0, len(param.RewritesEnd))
	for i := range param.RewritesEnd {
		mid, err := getRewriteMiddleware(param.RewritesEnd[i], middleware.SkipRests)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteEndMids = append(rewriteEndMids, mid)
		}
	}

	// redirects
	redirectMids := make([]middleware.Middleware, 0, len(param.Redirects))
	for i := range param.Redirects {
		mid, err := getRedirectMiddleware(param.Redirects[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			redirectMids = append(redirectMids, mid)
		}
	}

	// proxies
	proxyMids := make([]middleware.Middleware, 0, len(param.Proxies))
	for i := range param.Proxies {
		mid, err := getProxyMiddleware(param.Proxies[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			proxyMids = append(proxyMids, mid)
		}
	}

	// returns
	returnMids := make([]middleware.Middleware, 0, len(param.Returns))
	for i := range param.Returns {
		mid, err := getReturnStatusMiddleware(param.Returns[i], statusPageMids)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			returnMids = append(returnMids, mid)
		}
	}

	// combine all mids
	preMids = make([]middleware.Middleware, 0, len(rewriteMids)+
		len(rewriteEndMids)+
		len(redirectMids)+
		len(proxyMids)+
		len(returnMids),
	)
	preMids = append(preMids, rewriteMids...)
	preMids = append(preMids, rewriteEndMids...)
	preMids = append(preMids, redirectMids...)
	preMids = append(preMids, proxyMids...)
	preMids = append(preMids, returnMids...)

	postMids = make([]middleware.Middleware, 0, len(param.StatusPages))
	postMids = append(postMids, statusPageMids...)

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
