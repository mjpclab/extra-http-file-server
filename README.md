# Extra HTTP File Server

Extra HTTP File Server is based on Go HTTP File Server, with extra features.
It provides frequently used features for a simple static website.

# Different to Go HTTP File Server

## Code base

Based on Go HTTP File Server's main branch, dropped support for legacy Go version.
This means it is impossible to use legacy Go version to compile binaries for legacy systems, e.g. Windows XP.

## New options

```
--rewrite <separator><match><separator><replace>
    Transform a request URL (in the form of "/request/path?param=value")
    into another one if it is matched by regular expression `match`.

    The rewrite target is specified by `replace`.
    Use `$0` to represent the whole match in `match`.
    use `$1` - `$9` to represent sub matches in `match`.

    Rewrite will not evaluate recursively.

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

--return <separator><match><separator><status-code>[<separator><fs-path>]
    When request URL (in the form of "/request/path?param=value")
    is matched by `match`, return the status code `status-code`
    immediately and stop processing.
    Optional response page content is specified by `fs-path`.

--status-page <separator><status-code><separator><fs-path>
    When response status is `status-code`, respond with the file content from `fs-path`.
```

## Option processing order

- `--rewrite` executed to transform the URL.
- `--redirect` executed if URL matched, and stop processing.
- `--proxy` executed if URL matched, and stop processing.
- `--return` executed if URL matched, and stop processing.

## Examples

Perform redirect according to `redirect` param:

```sh
# when requesting http://localhost:8080/redirect/www.example.com, redirect to https://www.example.com
ehfs -l 8080 -r /path/to/share --redirect '#/redirect/(.*)#https://$1'
```
