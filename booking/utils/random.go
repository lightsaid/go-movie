package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const chars = "qwertyuiopasdfghjklzxcvbnm1234567890"

// 初始化随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt 生成一个随机数字，范围（min~max）
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString 生成随机字符串, 长度 = n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(chars)
	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomLonLat() string {
	head := RandomInt(1, 200)
	last := RandomInt(1000, 999999)
	return fmt.Sprintf("%d.%d", head, last)
}

func RandomSeatNumber() string {
	head := RandomInt(1, 12)
	last := RandomInt(1, 20)
	return fmt.Sprintf("%d排-%d座", head, last)
}

func RandomLanguage() string {
	langs := []string{"国语", "粤语", "英语"}
	return langs[rand.Intn(len(langs))]
}

func RandomGenres() string {
	genres := []string{"戏剧", "动作", "搞笑", "推理", "武侠", "战争", "言情"}
	return genres[rand.Intn(len(genres))]
}
