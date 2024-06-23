# Extra HTTP File Server

Extra HTTP File Server基于Go HTTP File Server，附带额外功能。
它为简单静态网站提供了常用的功能。

![Extra HTTP File Server pages](doc/ehfs.gif)

# 与Go HTTP File Server的区别

## 代码库

基于Go HTTP File Server主分支，放弃了对旧版Go的支持。
这意味着不能使用旧的Go版本来编译较老系统的二进制文件，例如Windows XP。

## 行为变化
对于PKI验证URL`/.well-known/`，
即使指定了`--to-https`，也将跳过从http:到https:的重定向。

## 新增选项

```
--ip-allow <IP>|<network/prefix> ...
--ip-allow-file <file> ...
    只允许来自指定的IP或网络的客户端访问。
    不匹配的客户端IP会被拒绝访问。

--ip-deny <IP>|<network/prefix> ...
--ip-deny-file <file> ...
    只拒绝来自指定的IP或网络的客户端访问。
    不匹配的客户端IP会被允许访问。

--rewrite-host <分隔符><match><分隔符><replace>
    如果请求的host+URL（“host[:port]/request/path?param=value”的形式）匹配正则表达式`match`，
    将其重写为另一URL。

    重写的目标由`replace`指定。
    使用`$0`表示`match`的完整匹配。
    使用`$1`-`$9`来表示`match`中的子匹配。
--rewrite-host-post <分隔符><match><分隔符><replace>
    与--rewrite-host相似，但在重定向无匹配后执行。
--rewrite-host-end <分隔符><match><分隔符><replace>
    与--rewrite-host-post相似，但匹配后跳过后续处理流程。

--rewrite <分隔符><match><分隔符><replace>
    如果请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`，
    将其重写为另一种形式。

    重写的目标由`replace`指定。
    使用`$0`表示`match`的完整匹配。
    使用`$1`-`$9`来表示`match`中的子匹配。
--rewrite-post <分隔符><match><分隔符><replace>
    与--rewrite相似，但在重定向无匹配后执行。
--rewrite-end <分隔符><match><分隔符><replace>
    与--rewrite-post相似，但匹配后跳过后续处理流程。

--redirect <分隔符><match><分隔符><replace>[<separator><status-code>]
    当请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`时，
    执行HTTP重定向。

    重定向目标由`replace`指定。
    使用`$0`表示`match`的完整匹配。
    使用`$1`-`$9`来表示`match`中的子匹配。

    可选的`status_code`指定HTTP重定向代码。 默认为301。

--proxy <分隔符><match><分隔符><replace>
    如果请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`，
    将代理请求另一个目标。

    代理的目标由`replace`指定。
    使用`$0`表示`match`的完整匹配。
    使用`$1`-`$9`来表示`match`中的子匹配。

--return <分隔符><match><分隔符><status-code>
    当请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`时，
    立即返回状态码`status-code`并停止处理。
--to-status <分隔符><match><分隔符><status-code>
    与--return类似，但在ghfs内部处理流程完成后执行。

--status-page <分隔符><status-code><分隔符><fs-path>
    当响应状态码为`status-code`时，用文件`fs-path`的内容来响应。

--gzip-static
    当请求资源FILE时，如果客户端支持gzip解码，则尝试查找FILE.gz并输出为gzip压缩的内容。

--header-add <分隔符><match><分隔符><name><分隔符><value>
--header-set <分隔符><match><分隔符><name><分隔符><value>
    当请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`时，
    添加或设置响应头。
```

## 处理顺序

- 如果客户端IP不匹配`--ip-allow`或`--ip-allow-file`，返回403状态并停止处理
  - 如果状态码匹配，执行`--status-page`并停止处理。
- 如果客户端IP匹配`--ip-deny`或`--ip-deny-file`，返回403状态并停止处理
  - 如果状态码匹配，执行`--status-page`并停止处理。
- 如果URL匹配，执行`--rewrite-host`和`--rewrite`以转换URL。
- 如果URL匹配，执行`--redirect`并停止处理。
- 如果URL匹配，执行`--rewrite-host-post`和`--rewrite-post`以转换URL。
- 如果URL匹配，执行`--rewrite-host-end`和`--rewrite-end`以转换URL，跳过其余处理流程，例如`--rewrite[-host]-end`、`--proxy`、`--return`等。
- 如果URL匹配，执行`--proxy`并停止处理。
  - 如果URL匹配，执行`--header-add`和`--header-set`并停止处理。
- 如果URL匹配，执行`--return`并停止处理。
  - 如果URL匹配，执行`--header-add`和`--header-set`并停止处理。
  - 如果状态码匹配，执行`--status-page`并停止处理。
- ghfs内部处理流程
- 如果URL匹配，执行`--header-add`和`--header-set`。
- 如果URL匹配，执行`--to-status`并停止处理。
  - 如果状态码匹配，执行`--status-page`并停止处理。
- 如果状态码匹配，执行`--status-page`并停止处理。

## 举例

根据`redirect`参数执行重定向：

```sh
# 当请求 http://localhost:8080/redirect/www.example.com时，重定向到https://www.example.com
ehfs -l 8080 -r /path/to/share --redirect '#/redirect/(.*)#https://$1'
```

访问静态页面URL无须包含`.html`后缀：
- 将包含`.html`后缀的URL重定向到不包含的
- 重写不包含`.html`后缀的URL至带有后缀

```sh
ehfs -l 8080 -r /path/to/share --redirect '#(.*)\.html#$1' --rewrite-post '#^.*/[^/.]+$#$0.html'
```

指定404状态页文件：

```sh
ehfs -l 8080 -r /path/to/share --status-page '#404#/path/to/404/file'
```

拒绝显示关键性文件或目录，返回403状态：

```sh
ehfs -l 8080 -r /path/to/share --return '#.git|.htaccess#403'
```

## 编译
至少需要Go 1.20版本。
```sh
go build main.go
```
会在当前目录生成"main"可执行文件。
