# Live Chat Capturer

捕获油管直播评论。

## Usage

Server:

需要go 1.11.4,其他版本能否运行未知。

```bash
git clone https://github.com/HoshioriAlice/LiveChatCapturer
cd server
go build
./LiveChatCapturer ${服务器端口}
```

需要在墙外服务器运行。

Client：

需要.NET Framework 4.6.1。

编辑client文件下的config.txt并保存:
```
直播间地址（目前是夸哥的）
服务器IP（目前是自己的服务器，在哪台服务器上部署了server端就填哪台服务器的IP）
服务器端口（目前是自己的服务器开放的端口，填上面那个${服务器端口}
屏蔽正则表达式1
屏蔽正则表达式2
屏蔽正则表达式3
......
```

然后运行LiveChatCapture.exe。


## 过程

先获取直播间网页，从中找到获取livechat需要的continuation信息，然后每隔一秒发送get_live_chat请求获取评论信息。

Client与服务器通信：

```bash
Client:
# 发送直播间地址
GET / HTTP/1.1
Action: Connect
Live-Page: ${live_url}

Server:
# 忽略HTTP响应头，只显示正文
# 成功时
{
  "status": "Success",
  "continuation": "${continuation}"
}
# 失败时
{
  "status": "Failed",
  "continuation": ""
}

Client:
# 更新直播评论
GET / HTTP/1.1
Action: Update
Continuation: ${continuation}

Server:
# 忽略HTTP响应头，只显示正文
# 成功时
{
  "continuation": ${new_continuation},
  "messages": [
    {
      "sender": ${sender},
      "message": ${message},
      # 非SC时该字段为空字符串
      "purchase": ${purchase}
    },
    {
      "sender": ${sender2},
      "message": ${message2},
      # 非SC时该字段为空字符串
      "purchase": ${purchase2}
    }
    ...
  ]
}

# 失败时
HTTP/1.1 404 Not Found

```

## TODO

我真的不会写C#，Bug一堆是肯定的，界面是我乱糊的，求求有能天狗改下。

服务器端也是瞎糊的，好久没碰过go，如果跑出bug来也请改下。
