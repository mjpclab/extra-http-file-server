# Extra HTTP File Server

Extra HTTP File Server is based on Go HTTP File Server, with extra features.
It provides frequently used features for a simple static website.

# Different to Go HTTP File Server

## Code base

Based on Go HTTP File Server's main branch, dropped support for legacy Go version.
This means it is impossible to use legacy Go version to compile binaries for legacy systems, e.g. Windows XP.

## New options

```
--ip-allow <IP>|<network/prefix> ...
--ip-allow-file <file> ...
    Only allow client access from specific IP or network.
    Unmatched client IP will be denied.

--ip-deny <IP>|<network/prefix> ...
--ip-deny-file <file> ...
    Only denly client access from specific IP or network.
    Unmatched client IP will be allowed to access.

--rewrite <separator><match><separator><replace>
    Transform a request URL (in the form of "/request/path?param=value")
    into another one if it is matched by regular expression `match`.

    The rewrite target is specified by `replace`.
    Use `$0` to represent the whole match in `match`.
    use `$1` - `$9` to represent sub matches in `match`.
--rewrite-post <sep><match><sep><replace>
    Similar to --rewrite, but executes after redirects has no match.
--rewrite-end <sep><match><sep><replace>
    Similar to --rewrite-post, but skip rest process if matched.

--redirect <separator><match><separator><replace>[<separator><status-code>]
    Perform an HTTP redirect when request URL (in the form of "/request/path?param=value")
    is matched by regular expression `match`.

    The redirect target is specified by `replace`.
    Use `$0` to represent the whole match in `match`.
    use `$1` - `$9` to represent sub matches in `match`.

    Optional `status_code` specifies HTTP redirect code. defaults to 301.

--proxy <separator><match><separator><replace>
    Proxy a request URL (in the form of "/request/path?param=value")
    to target if it is matched by regular expression `match`.

    The proxy target is specified by `replace`.
    Use `$0` to represent the whole match in `match`.
    use `$1` - `$9` to represent sub matches in `match`.

--return <separator><match><separator><status-code>
    When request URL (in the form of "/request/path?param=value")
    is matched by `match`, return the status code `status-code`
    immediately and stop processing.
--to-status <separator><match><separator><status-code>
    Similar to --return, but process after ghfs internal process finished.

--status-page <separator><status-code><separator><fs-path>
    When response status is `status-code`, respond with the file content from `fs-path`.

--header-add <separator><match><separator><name><separator><value>
--header-set <separator><match><separator><name><separator><value>
    Add or set response header if URL(in the form of "/request/path?param=value")
    matches `match`.
```

## Processing order

- if client IP not match `--ip-allow` or `--ip-allow-file`, return status 403, and stop processing
  - `--status-page` executed if status code matched, and stop processing.
- if client IP match `--ip-deny` or `--ip-deny-file`, return status 403, and stop processing
  - `--status-page` executed if status code matched, and stop processing.
- `--rewrite` executed to transform the URL if matched.
- `--redirect` executed if URL matched, and stop processing.
- `--rewrite-post` executed to transform the URL if matched.
- `--rewrite-end` executed to transform the URL if matched, and skip rest of `--rewrite-end`, `--redirect`, `--proxy` and `--return`.
- `--proxy` executed if URL matched, and stop processing.
  - `--header-add` and `--header-set` executed if URL matched, and stop processing.
- `--return` executed if URL matched, and stop processing.
  - `--header-add` and `--header-set` executed if URL matched, and stop processing.
  - `--status-page` executed if status code matched, and stop processing.
- ghfs internal process
- `--header-add` and `--header-set` executed if URL matched.
- `--to-status` executed if URL matched, and stop processing.
  - `--status-page` executed if status code matched, and stop processing.
- `--status-page` executed if status code matched, and stop processing.

## Examples

Perform redirect according to `redirect` param:

```sh
# when requesting http://localhost:8080/redirect/www.example.com, redirect to https://www.example.com
ehfs -l 8080 -r /path/to/share --redirect '#/redirect/(.*)#https://$1'
```

Serve static page without `.html` suffix in URL:
- redirect URL contains `.html` suffix to no suffix
- rewrite URL without suffix to with `.html` suffix 

```sh
ehfs -l 8080 -r /path/to/share --redirect '#(.*)\.html#$1' --rewrite-post '#^.*/[^/.]+$#$0.html'
```

Specify page for 404 status:

```sh
ehfs -l 8080 -r /path/to/share --status-page '#404#/path/to/404/file'
```

Refuse to serve for critical files or directories, returns 403 status:

```sh
ehfs -l 8080 -r /path/to/share --return '#.git|.htaccess#403'
```
