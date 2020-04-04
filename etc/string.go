// Author : Eric Kim
// Build Date : 23 Jul 2018  Last Update 02 Aug 2018
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package etc

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"reflect"
	"strings"
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

func ParentDir(params ...string) string {

	var workdir string
	if len(params) == 0 {
		workdir, _ = os.Getwd()
	} else {
		workdir = params[0]
	}

	sp := strings.Split(workdir, "/")
	parentdir := ""
	for i := 1; i < len(sp)-1; i++ {
		parentdir += "/" + sp[i]
	}
	return parentdir
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func TableName(m interface{}) string {
	s := reflect.TypeOf(m).Name()
	return SnakeString(s)
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data[:])
}
