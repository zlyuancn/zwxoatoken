/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/28
   Description :
-------------------------------------------------
*/

package zwxoatoken

import (
    "crypto/tls"
    "net/http"
)

var DefaultClient = &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }}
