# zwxoatoken
> 获取微信公众号token

## 获取
> `go get -u -v github.com/zlyuancn/zwxoatoken`

## 实例

```go
    // 获取access_token
    at, err := GetAccessToken("your_appid", "your_secret")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(at.AccessToken, at.ExpiresIn)

    // 获取jsapi_ticket
    jt, err := GetJsapiTicket(at.AccessToken)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(jt.Ticket, jt.ExpiresIn)

    // 获取签名数据
    url := "http://yourdomain/route?parame=value#123"
    out := MakeSign(jt.Ticket, url)
    log.Println(out.Noncestr, out.Signature, out.Timestamp)
```
