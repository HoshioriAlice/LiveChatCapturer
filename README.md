# Live Chat Capturer

捕获油管直播评论。

## Usage

```bash
git clone https://github.com/HoshioriAlice/LiveChatCapturer
go build
./LiveChatCapturer ${直播间URL}
```

需要在墙外服务器运行。

## 过程

先获取直播间网页，从中找到获取livechat需要的continuation信息，然后每隔一秒发送get_live_chat请求获取评论信息。