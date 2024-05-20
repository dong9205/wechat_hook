# 通过微信发送信息

## 运行服务

### 默认配置运行

```bash
docker run --rm -p 9200:9200 wechat_hook serve
```
运行该命令后，会将二维码打印到屏幕上，然后通过微信扫描二维码即可登录

### 指定配置文件

```bash
docker run --log-opt  max-size=100m --log-opt max-file=3 -d --name wechat_hook -v /mnt/data/wechat_hook:/mnt/data/wechat_hook -p 9200:9200  dong9205/wechat_hook:v0.0.3 serve -c /mnt/data/wechat_hook/config.yaml
```

## 调用服务接口

### 向其他联系人发送消息

```bash
curl --location --request POST 'http://localhost:9200/api/push/msg' --header 'Content-Type: application/json' --data-raw '{
    "dest": "联系人名称",
    "dest_type": "friend",
    "msg": "Hello Friend!"
}'
```

### 向群里发送消息

> 在登陆成功后，需要其他人员先再群里发送一条消息

```bash
curl --location --request POST 'http://localhost:9200/api/push/msg' --header 'Content-Type: application/json' --data-raw '{
    "dest": "哒哒哒",
    "dest_type": "group",
    "msg": "Hello Group!"
}'
```
