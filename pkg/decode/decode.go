package decode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func ToMD5(address string) string {
	hash := md5.Sum([]byte(address))
	fmt.Println(hex.EncodeToString(hash[:]))
	return hex.EncodeToString(hash[:])
}
