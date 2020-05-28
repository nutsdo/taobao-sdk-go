package helpers

import (
	"errors"
	"fmt"
	"regexp"
)

//识别淘口令，并返回
func IdentifyTpwd(code string) (tkl string, err error) {

	//识别货币符号｜数字符号
	pattern := `([\p{Sc}])\w{8,12}(\p{Sc})`

	matched,_ := regexp.MatchString(pattern, code)
	reg,_ := regexp.Compile(pattern)
	result := reg.FindAllStringSubmatch(code,-1)

	fmt.Printf("%v",tkl)

	if !matched {
		return "", errors.New("未识别淘口令")
	}

	fmt.Printf("匹配淘口令的结果:%#v\n",result)

	return result[0][0],nil
}