package util

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 随机生成5位数字
func Random5() string {
	rand.Seed(time.Now().Unix())
	var codes = make([]string, 5)
	var code string
	for i := 0; i < 5; i++ {
		codes[i] = strconv.Itoa(rand.Int() % 10)
	}
	code = strings.Join(codes, "")
	return code
}

// 使用正则表达式识别手机号码
func IsMobile(mobile string) error {
	result, _ := regexp.MatchString(`^(1[3|4|5|8|9][0-9]\d{4,8})$`, mobile)
	if result {
		return nil
	} else {
		return errors.New("不是一个正确的手机号码")
	}
}
