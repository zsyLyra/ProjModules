package encrypt

import (
	"ProjModules/utils/setting"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func EncodeMd5(value string) string {
	m := md5.New()
	v := fmt.Sprintf("%s%s%s", time.Now().String(), value, setting.AppSetting.SaltSecret)
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}