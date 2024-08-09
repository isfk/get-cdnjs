# get-cdnjs

> 从 [cdnjs](https://cdnjs.com/) 获取前端库，自动上传到自有七牛云。 
>
> 支持配置代理
>
> 支持查看七牛已上传资源库

## 使用

1. 下载对应系统的文件
2. 解压后得到 `get-cdnjs` 文件
3. 修改配置文件 `config.yaml`
4. 命令行运行 `./get-cdnjs --config config.yaml` 命令
5. 命令行运行 `./get-cdnjs list --config config.yaml` 命令, 可查看上传到七牛的资源库

### config.yaml 配置说明：

```yaml
access_key: "**************************"    # 七牛密钥
secret_key: "**************************"    # 七牛密钥
bucket: "**********"                        # 你的七牛空间bucket名
cdn_domain: "https://***.isfk.cn"           # 你的七牛云cdn域名
proxy: "http://127.0.0.1:7897"              # 代理, 没有可以留空
file_path: "cdnjs"                          # 存储到七牛的目录，可以是多级目录 如: cdnjs/libs
```

## 操作示例

```sh
# 上传jquery到七牛
> ./get-cdnjs --config config.yaml              
使用配置: config.yaml
当前配置不使用代理
请输入库名: jquery
正在查找版本号, 请稍等...
查找链接: https://api.cdnjs.com/libraries/jquery
最新版本号: 3.7.1, 可选版本号: 
[ 1.10.0 1.10.1 1.10.2 1.11.0 1.11.0-beta3 1.11.0-rc1 1.11.1 1.11.1-beta1 1.11.1-rc1 1.11.1-rc2 1.11.2 1.11.3 1.12.0 1.12.1 1.12.2 1.12.3 1.12.4 1.2.3 1.2.6 1.3.0 1.3.1 1.3.2 1.4.0 1.4.1 1.4.2 1.4.3 1.4.4 1.5.1 1.6.1 1.6.2 1.6.3 1.6.4 1.7 1.7.1 1.7.2 1.8.0 1.8.1 1.8.2 1.8.3 1.9.0 1.9.1 2.0.0 2.0.1 2.0.2 2.0.3 2.1.0 2.1.0-beta2 2.1.0-beta3 2.1.0-rc1 2.1.1 2.1.1-beta1 2.1.1-rc1 2.1.1-rc2 2.1.2 2.1.3 2.1.4 2.2.0 2.2.1 2.2.2 2.2.3 2.2.4 3.0.0 3.0.0-alpha1 3.0.0-beta1 3.0.0-rc1 3.1.0 3.1.1 3.2.0 3.2.1 3.3.0 3.3.1 3.4.0 3.4.1 3.5.0 3.5.1 3.6.0 3.6.1 3.6.2 3.6.3 3.6.4 3.7.0 3.7.1 4.0.0-beta 4.0.0-beta.2 ]
回车使用最新版本号, 否则请输入版本号: 
库名: jquery, 版本号: 3.7.1
正在查找文件, 请稍等...
查找链接: https://api.cdnjs.com/libraries/jquery/3.7.1
该版本存在以下文件: 
[ jquery.js jquery.min.js jquery.min.map jquery.slim.js jquery.slim.min.js jquery.slim.min.map ]
回车开始抓取, 结束请按键: <Ctrl + C> 

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.js
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.js

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.min.js

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.map
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.min.map

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.slim.js
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.js

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.slim.min.js
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.min.js

开始抓取文件: https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.slim.min.map
抓取成功: https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.min.map

jquery 所有文件已抓取到七牛:
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.min.js
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.min.map
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.js
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.min.js
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.min.map
https://pb.isfk.cn/cdnjs/jquery/3.7.1/jquery.slim.js

结束请按键: <Ctrl + C>


# 查看上传到七牛的资源库
> ./get-cdnjs list --config config.yaml     
使用配置: config.yaml
当前配置不使用代理
─ cdnjs:
  ├─ echarts:
  │  └─ 5.5.0
  ├─ jquery:
  │  └─ 1.10.0
  │  └─ 3.7.1
  ├─ layui:
  │  └─ 2.9.14
  ├─ swagger-ui:
  │  └─ 5.17.14
```