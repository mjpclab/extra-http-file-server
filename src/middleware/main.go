package middleware

import (
	"errors"
	"mjpclab.dev/ehfs/src/param"
	"mjpclab.dev/ehfs/src/util"
	"mjpclab.dev/ghfs/src/middleware"
	baseParam "mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/serverError"
)

var errInvalidParamValue = errors.New("invalid param value")
var errParamCountNotMatch = errors.New("base-param count is not equal to param count")

func ParamToMiddlewares(baseParam *baseParam.Param, param *param.Param) (preMids, inMids, postMids []middleware.Middleware, errs []error) {
	var mid middleware.Middleware
	var err error
	var es []error

	// headers
	headerMids := make([]middleware.Middleware, 0, len(param.HeaderAdds)+len(param.HeaderSets))
	for i := range param.HeaderAdds {
		mid, err = getHeaderAddMiddleware(param.HeaderAdds[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			headerMids = append(headerMids, mid)
		}
	}
	for i := range param.HeaderSets {
		mid, err = getHeaderSetMiddleware(param.HeaderSets[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			headerMids = append(headerMids, mid)
		}
	}

	// dependent: gzip static
	gzipStaticMids := make([]middleware.Middleware, 0, 1)
	if param.GzipStatic {
		mid, err = getGzipStaticMiddleware()
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			gzipStaticMids = append(gzipStaticMids, mid)
		}
	}

	// dependent: status pages
	statusPageMids := make([]middleware.Middleware, 0, len(param.StatusPages))
	for i := range param.StatusPages {
		mid, err = getStatusPageMiddleware(param.StatusPages[i], param.GzipStatic)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			statusPageMids = append(statusPageMids, mid)
		}
	}

	// ip allows
	ipAllowMids := make([]middleware.Middleware, 0, 1)
	mid, es = getIPAllowMiddleware(param.IPAllows, param.IPAllowFiles, statusPageMids)
	errs = append(errs, es...)
	if mid != nil {
		ipAllowMids = append(ipAllowMids, mid)
	}

	// ip denies
	ipDenyMids := make([]middleware.Middleware, 0, 1)
	mid, es = getIPDenyMiddleware(param.IPDenies, param.IPDenyFiles, statusPageMids)
	errs = append(errs, es...)
	if mid != nil {
		ipDenyMids = append(ipDenyMids, mid)
	}

	// rewrite hosts
	rewriteHostMids := make([]middleware.Middleware, 0, len(param.RewriteHosts))
	for i := range param.RewriteHosts {
		mid, err = getRewriteHostMiddleware(param.RewriteHosts[i], middleware.GoNext)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteHostMids = append(rewriteHostMids, mid)
		}
	}

	// rewrite hosts post
	rewriteHostPostMids := make([]middleware.Middleware, 0, len(param.RewriteHostsPost))
	for i := range param.RewriteHostsPost {
		mid, err = getRewriteHostMiddleware(param.RewriteHostsPost[i], middleware.GoNext)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteHostPostMids = append(rewriteHostPostMids, mid)
		}
	}

	// rewrite hosts end
	rewriteHostEndMids := make([]middleware.Middleware, 0, len(param.RewriteHostsEnd))
	for i := range param.RewriteHostsEnd {
		mid, err = getRewriteHostMiddleware(param.RewriteHostsEnd[i], middleware.SkipRests)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteHostEndMids = append(rewriteHostEndMids, mid)
		}
	}

	// rewrites
	rewriteMids := make([]middleware.Middleware, 0, len(param.Rewrites))
	for i := range param.Rewrites {
		mid, err = getRewriteMiddleware(param.Rewrites[i], middleware.GoNext)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteMids = append(rewriteMids, mid)
		}
	}

	// rewrites post
	rewritePostMids := make([]middleware.Middleware, 0, len(param.RewritesPost))
	for i := range param.RewritesPost {
		mid, err = getRewriteMiddleware(param.RewritesPost[i], middleware.GoNext)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewritePostMids = append(rewritePostMids, mid)
		}
	}

	// rewrites end
	rewriteEndMids := make([]middleware.Middleware, 0, len(param.RewritesEnd))
	for i := range param.RewritesEnd {
		mid, err = getRewriteMiddleware(param.RewritesEnd[i], middleware.SkipRests)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			rewriteEndMids = append(rewriteEndMids, mid)
		}
	}

	// redirects
	redirectMids := make([]middleware.Middleware, 0, len(param.Redirects))
	for i := range param.Redirects {
		mid, err = getRedirectMiddleware(param.Redirects[i])
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			redirectMids = append(redirectMids, mid)
		}
	}

	// proxies
	proxyMids := make([]middleware.Middleware, 0, len(param.Proxies))
	for i := range param.Proxies {
		mid, err = getProxyMiddleware(param.Proxies[i], headerMids)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			proxyMids = append(proxyMids, mid)
		}
	}

	// returns
	returnMids := make([]middleware.Middleware, 0, len(param.Returns))
	for i := range param.Returns {
		mid, err = getReturnStatusMiddleware(param.Returns[i], util.Concat(headerMids, statusPageMids))
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			returnMids = append(returnMids, mid)
		}
	}

	// headers (moved to dependent)

	// to statuses
	toStatusMids := make([]middleware.Middleware, 0, len(param.ToStatuses))
	for i := range param.ToStatuses {
		mid, err = getReturnStatusMiddleware(param.ToStatuses[i], statusPageMids)
		errs = serverError.AppendError(errs, err)
		if mid != nil {
			toStatusMids = append(toStatusMids, mid)
		}
	}

	// status pages (moved to dependent)

	// gzip static (moved to dependent)

	// pki validation skip to-https
	pkiValidationSkipToHttpsMids := make([]middleware.Middleware, 0, 1)
	if baseParam.ToHttps {
		mid = getPkiValidationSkipToHttpsMiddleware()
		pkiValidationSkipToHttpsMids = append(pkiValidationSkipToHttpsMids, mid)
	}

	// perms
	permsUrlsMids := make([]middleware.Middleware, 0, 1)
	mid, es = getPermsUrlsMiddleware(param.PermsUrls)
	errs = append(errs, es...)
	if mid != nil {
		permsUrlsMids = append(permsUrlsMids, mid)
	}

	permsDirsMids := make([]middleware.Middleware, 0, 1)
	mid, es = getPermsDirsMiddleware(param.PermsDirs)
	errs = append(errs, es...)
	if mid != nil {
		permsDirsMids = append(permsDirsMids, mid)
	}

	// combine mids
	preMids = util.Concat(
		ipAllowMids,
		ipDenyMids,
		rewriteHostMids,
		rewriteMids,
		redirectMids,
		rewriteHostPostMids,
		rewritePostMids,
		rewriteHostEndMids,
		rewriteEndMids,
		proxyMids,
		returnMids,
		pkiValidationSkipToHttpsMids,
	)

	postMids = util.Concat(
		headerMids,
		toStatusMids,
		statusPageMids,
		gzipStaticMids,
	)

	inMids = util.Concat(
		permsUrlsMids,
		permsDirsMids,
	)

	return
}

func ApplyMiddlewares(baseParams []*baseParam.Param, params []*param.Param) (errs []error) {
	if len(baseParams) != len(params) {
		return []error{errParamCountNotMatch}
	}

	for i := range baseParams {
		var es []error
		baseParams[i].PreMiddlewares, baseParams[i].InMiddlewares, baseParams[i].PostMiddlewares, es = ParamToMiddlewares(baseParams[i], params[i])
		errs = append(errs, es...)
	}

	return
}
