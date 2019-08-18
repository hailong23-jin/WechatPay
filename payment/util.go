package payment

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i :=0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("xml")] = v.Field(i).Interface()
	}

	fmt.Println("Struct2Map: ",data)

	return data
}

// 获取随机字符串 n 位
func GetNonceStr(n int) string{
	rand.Seed(time.Now().Unix())
	arr := []string{"1","2","3","4","5","6","7","8","9","0",
					"A","B","C","D","E","F","G","H","I","J",
					"K","L","M","N","O","P","Q","R","S","T",
					"U","V","W","X","Y","Z"}
	var buff bytes.Buffer
	for i := 0; i < n; i++ {
		buff.WriteString(arr[rand.Intn(36)])
	}
	return buff.String()
}

func GetOutTradeNo() string {
	timeStamp := time.Now().Unix()
	str := strconv.FormatInt(timeStamp, 10)
	return fmt.Sprintf("%d%s",timeStamp, GetRandom(31-len(str)) )
}

func GetRandom(n int) string{
	rand.Seed(time.Now().Unix())
	arr := []string{"1","2","3","4","5","6","7","8","9","0"}
	var buff bytes.Buffer
	for i := 0; i < n; i++ {
		buff.WriteString(arr[rand.Int31n(10)])
	}
	return buff.String()
}

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}