package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func LoadConfig(section string) (map[string]string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return map[string]string{}, errors.New("为找到配置文件")
		} else {
			return map[string]string{}, errors.New("读取配置文件失败")
		}
	}
	if len(viper.GetStringMapString(section)) == 0 {
		return map[string]string{}, errors.New("无法获取键值")
	} else {
		return viper.GetStringMapString(section), nil
	}
}

func LoadDbConfig(section string) (string, error) {
	conf, err := LoadConfig(section)
	if err != nil {
		return "", err
	}
	if conf["username"] == "" || conf["password"] == "" || conf["hostname"] == "" || conf["port"] == "" || conf["database"] == "" || conf["charset"] == "" {
		return "", errors.New("获取配置失败")
	}
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf["username"], conf["password"], conf["hostname"], conf["port"], conf["database"], conf["charset"])
	return str, nil
}

//int 转 string
func IntToStr(t int) string {
	str := strconv.Itoa(t)
	return str
}

//string 转 int
func StrToInt(str string) int {
	t, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return t
}

//string 转 int64
func StrToInt64(s string) int64 {
	t, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return t
}

// string 转 float64
func StrToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// float64 转 string
func Float64ToStr(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	return s
}

// float64 转 int
func Float64ToInt(f float64) int {
	return int(f)
}

//int64 转 string
func Int64ToStr(t int64) string {
	str := strconv.FormatInt(int64(t), 10)
	return str
}

//int64 转 int
func Int64ToInt(t int64) int {
	str := strconv.FormatInt(int64(t), 10)
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

//验证密码
func CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil {
		return true
	} else {
		return false
	}
}
