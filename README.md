# get-cdnjs

- 从 [jsDelivr](https://www.jsdelivr.com/) 获取 npm 包，自动上传到自有七牛云。 
- 支持配置代理
- 支持查看七牛已上传资源库
- 支持删除七牛已存在的资源库

## 使用

1. 下载对应系统的文件
2. 解压后得到 `get-cdnjs` 文件
3. 修改配置文件 `config.yaml`
4. 命令行运行 `./get-cdnjs` 命令，或自定义配置 `./get-cdnjs --config config.yaml`
5. 命令行运行 `./get-cdnjs list` 命令, 可查看上传到七牛的资源库
5. 命令行运行 `./get-cdnjs delete` 命令, 可删除七牛已存在的资源库

### config.yaml 配置说明：

> 默认位置 `~/.get-cdnjs.yaml`
>
> 运行命令式可以指定配置 `--config your-config.yaml`

```yaml
access_key: "**************************"    # 七牛密钥
secret_key: "**************************"    # 七牛密钥
bucket: "**********"                        # 你的七牛空间bucket名
cdn_domain: "https://***.isfk.cn"           # 你的七牛云cdn域名
proxy: "http://127.0.0.1:7897"              # 代理, 没有可以留空
file_path: "npm"                            # 存储到七牛的目录，可以是多级目录 如: npm/libs
```

## 操作示例

```sh
# 上传jquery到七牛
> ./get-cdnjs

# 查看上传到七牛的资源库
> ./get-cdnjs list


# 删除七牛的资源库
> ./get-cdnjs delete
```