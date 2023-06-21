package middleware

import (
	"mjpclab.dev/ghfs/src/middleware"
	ghfsUtils "mjpclab.dev/ghfs/src/util"
	"net/http"
	"strings"
)

type permsInfo struct {
	path  string
	perms struct {
		upload  bool
		mkdir   bool
		delete  bool
		archive bool
	}
	users []string
}

type matchPath func(expectedPath string, context *middleware.Context) bool

func matchUrlPath(expectedPath string, context *middleware.Context) bool {
	return ghfsUtils.HasUrlPrefixDir(context.VhostReqPath, expectedPath)
}

func matchFsPath(expectedPath string, context *middleware.Context) bool {
	return ghfsUtils.HasFsPrefixDir(context.AliasFsPath, expectedPath)
}

func getPermsInfos(paramsList [][]string) (infos []permsInfo) {
	const pathIndex = 0
	const permsIndex = 1
	const userStartIndex = 2
	infos = make([]permsInfo, 0, len(paramsList))

	for _, params := range paramsList {
		info := permsInfo{
			path:  params[pathIndex],
			users: params[userStartIndex:],
		}

		hasPerm := false
		perms := strings.Split(strings.ToLower(params[permsIndex]), ",")
		for _, perm := range perms {
			perm = strings.TrimSpace(perm)
			switch perm {
			case "upload":
				info.perms.upload = true
				hasPerm = true
			case "mkdir":
				info.perms.mkdir = true
				hasPerm = true
			case "delete":
				info.perms.delete = true
				hasPerm = true
			case "archive":
				info.perms.archive = true
				hasPerm = true
			}
		}
		if !hasPerm {
			continue
		}

		infos = append(infos, info)
	}

	return
}

func updatePerms(matchPath matchPath, infos []permsInfo, context *middleware.Context) {
	if !context.AuthSuccess {
		return
	}

	for _, info := range infos {
		if !matchPath(info.path, context) {
			continue
		}

		authUserIndex := context.Users.FindIndex(context.AuthUserName)
		if authUserIndex < 0 {
			continue
		}

		for _, user := range info.users {
			if context.Users.FindIndex(user) != authUserIndex {
				continue
			}

			if info.perms.upload {
				*context.CanUpload = true
			}
			if info.perms.mkdir {
				*context.CanMkdir = true
			}
			if info.perms.delete {
				*context.CanDelete = true
			}
			if info.perms.archive {
				*context.CanArchive = true
			}

			break
		}
	}
}

func getPermsUrlsMiddleware(paramsList [][]string) (middleware.Middleware, []error) {
	permsInfos := getPermsInfos(paramsList)
	if len(permsInfos) == 0 {
		return nil, nil
	}
	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		updatePerms(matchUrlPath, permsInfos, context)
		return middleware.GoNext
	}, nil
}

func getPermsDirsMiddleware(paramsList [][]string) (middleware.Middleware, []error) {
	permsInfos := getPermsInfos(paramsList)
	if len(permsInfos) == 0 {
		return nil, nil
	}
	return func(w http.ResponseWriter, r *http.Request, context *middleware.Context) (result middleware.ProcessResult) {
		updatePerms(matchFsPath, permsInfos, context)
		return middleware.GoNext
	}, nil
}
