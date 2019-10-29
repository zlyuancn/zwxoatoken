/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/28
   Description :
-------------------------------------------------
*/

package zwxoatoken

import "testing"

func TestGetAccessToken(t *testing.T) {
    out, err := GetAccessToken("your_appid", "your_secret")
    if err != nil {
        t.Fatal(err)
    }
    t.Log(out.AccessToken, out.ExpiresIn)
}

func TestGetJsapiTicket(t *testing.T) {
    accesst_token := "your_accesst_token"
    out, err := GetJsapiTicket(accesst_token)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(out.Ticket, out.ExpiresIn)
}

func TestMakeSign(t *testing.T) {
    jsapi_ticket := "your_jsapi_ticket"
    url := "http://yourdomain/route?parame=value#123"
    out := MakeSign(jsapi_ticket, url)
    t.Log(out.Noncestr, out.Signature, out.Timestamp)
}
