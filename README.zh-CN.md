# Extra HTTP File Server

Extra HTTP File Server基于Go HTTP File Server，附带额外功能。
它为简单静态网站提供了常用的功能。

# 与Go HTTP File Server的区别

## 代码库

基于Go HTTP File Server主分支，放弃了对旧版Go的支持。
这意味着不能使用旧的Go版本来编译较老系统的二进制文件，例如Windows XP。

## 新增选项

```
--redirect <分隔符><match><separator><分隔符>[<separator><status_code>]
    当请求的URL（“/request/path?param=value”的形式）匹配正则表达式`match`时，
    执行HTTP重定向。

    重定向目标由`replace`指定。
    使用`$0`表示`match`的完整匹配。
    使用`$1`-`$9`来表示`match`中的子匹配。

    可选的`status_code`指定HTTP重定向代码。 默认为301。
```

## 举例

根据`redirect`参数执行重定向：

```sh
# 当请求 http://localhost:8080/redirect/www.example.com时，重定向到https://www.example.com
ehfs -l 8080 -r /path/to/share --redirect '#/redirect/(.*)#https://$1'
```
