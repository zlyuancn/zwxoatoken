/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/28
   Description :
-------------------------------------------------
*/

package zwxoatoken

import (
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/zlyuancn/zerrors"
)

const (
    AccessTokenUri = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
    JsapiTicketUri = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

const (
    GrantType = "client_credential"
)

type AssessToken struct {
    Errcode     int    `json:"errcode"`
    Errmsg      string `json:"errmsg"`
    AccessToken string `json:"access_token"`
    ExpiresIn   int    `json:"expires_in"`
}

type JsapiTicket struct {
    Errcode   int    `json:"errcode"`
    Errmsg    string `json:"errmsg"`
    Ticket    string `json:"ticket"`
    ExpiresIn int    `json:"expires_in"`
}

type Signature struct {
    Noncestr  string `json:"noncestr"`
    Signature string `json:"signature"`
    Timestamp int    `json:"timestamp"`
}

// 获取公众号accesst_token
func GetAccessToken(appid, secret string) (*AssessToken, error) {
    url := fmt.Sprintf(AccessTokenUri, GrantType, appid, secret)
    resp, err := DefaultClient.Get(url)
    if err != nil {
        return nil, zerrors.Wrap(err, "请求token时发生错误")
    }
    defer resp.Body.Close()

    out := new(AssessToken)
    if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
        return nil, zerrors.Wrap(err, "收到的数据无法解析")
    }

    if out.Errcode != 0 {
        return nil, zerrors.Errorf("收到错误代码: errcode: %d, %s", out.Errcode, out.Errmsg)
    }

    return out, nil
}

// 获取jsapi_ticket
func GetJsapiTicket(accesst_token string) (*JsapiTicket, error) {
    url := fmt.Sprintf(JsapiTicketUri, accesst_token)
    resp, err := DefaultClient.Get(url)
    if err != nil {
        return nil, zerrors.Wrap(err, "请求jsapi_ticket时发生错误")
    }
    defer resp.Body.Close()

    out := new(JsapiTicket)
    if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
        return nil, zerrors.Wrap(err, "收到的数据无法解析")
    }

    if out.Errcode != 0 {
        return nil, zerrors.Errorf("收到错误代码: errcode: %d, %s", out.Errcode, out.Errmsg)
    }

    return out, nil
}

// 签名
//
// 微信 JS 接口签名校验工具地址: http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign
func MakeSign(jsapi_ticket string, now_url string) *Signature {
    limit := strings.Index(now_url, "#")
    if limit > 0 {
        now_url = now_url[:limit]
    }

    time_stamp := int(time.Now().Unix())
    noncestr := strconv.Itoa(time_stamp)

    signature := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", jsapi_ticket, noncestr, time_stamp, now_url)

    out := &Signature{
        Noncestr:  noncestr,
        Timestamp: time_stamp,
        Signature: Sha1(signature),
    }
    return out
}
