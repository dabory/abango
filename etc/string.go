// Author : Eric Kim
// Build Date : 23 Jul 2018  Last Update 02 Aug 2018
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package etc

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func RandString(i int) string {
	b := make([]byte, i)
	rand.Read(b)
	return (base64.URLEncoding.EncodeToString(b))[0:i]
}

func GetAskName() string {

	i := len(os.Args)
	if i < 2 {
		MyErr("ZMXCDKDALKSJD", errors.New("command arguments are less then 2"), true)
	} else {

		return os.Args[i-1]
	}
	return ""
}

// func (t *EnvConf) StrToStruct(str string) {
// 	if err := json.Unmarshal([]byte(str), &t); err != nil {
// 		MyErr("sjdfljsf", err, true)
// 	}
// 	return
// }

func randString(i int) string {
	b := make([]byte, i)
	rand.Read(b)
	return (base64.URLEncoding.EncodeToString(b))[0:i]
}

// func randNumber(len int) string { // 나중에 코드 반드시 리팩토링 할 것
// 	a := make([]int, len)
// 	for i := 0; i <= len-1; i++ {
// 		a[i] = rand.Intn(len)
// 		return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")
// 	}
// }

func randBytes(i int) []byte {
	return []byte(randString(i))
}

func myHash(data []byte, leng int) []byte {
	hash := sha256.New() //SHA-3 규격임.
	hash.Write(data)

	mdStr := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	rtn := ""
	if leng == 0 {
		rtn = mdStr
	} else {
		rtn = mdStr[10 : 10+leng]
	}
	return []byte(rtn)
}

func myToken(leng int) string {
	b := make([]byte, leng)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func reverseString(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

func reverseBytes(s []byte) []byte {
	cs := make([]byte, len(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return cs
}

func getCnt(s []byte, cnt int) []byte {
	// ret := ""
	var ret []byte
	if len(s) > cnt {
		ret = s[0:cnt]
	} else if len(s) < cnt {
		ret = append(s, strings.Repeat("=", cnt-len(s))...)
	} else {
		ret = s
	}
	return ret
}

// func structToMap(in interface{}, tag string) (map[string]interface{}, string) {
// 	out := make(map[string]interface{})

// 	v := reflect.ValueOf(in)
// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}

// 	// we only accept structs
// 	if v.Kind() != reflect.Struct {
// 		fmt.Errorf("ToMap only accepts structs; got %T", v)
// 		return nil, MyErr("only accepts structs", nil, false)
// 	}

// 	typ := v.Type()
// 	for i := 0; i < v.NumField(); i++ {
// 		// gets us a StructField
// 		fi := typ.Field(i)
// 		if tagv := fi.Tag.Get(tag); tagv != "" {
// 			out[tagv] = v.Field(i).Interface()
// 		}
// 	}
// 	return out, ""
// }

func parentDir() string { // Copy시메모리 소모 없슴.
	workDir, _ := os.Getwd()
	sp := strings.Split(workDir, "/")
	parentDir := ""
	for i := 1; i < len(sp)-1; i++ {
		parentDir += "/" + sp[i]
	}
	return parentDir
}

func getOTp(n int) []byte {
	const letters = "0123456789"
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		MyErr("rand.Read", err, false)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return bytes
}
