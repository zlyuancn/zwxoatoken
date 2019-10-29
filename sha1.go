/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/28
   Description :
-------------------------------------------------
*/

package zwxoatoken

import (
    "crypto/sha1"
    "encoding/hex"
)

func Sha1(text string) string {
    c := sha1.New();
    c.Write([]byte(text))
    return hex.EncodeToString(c.Sum(nil))
}
