package payment

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	appId = ""
	mch_id = ""
	apiKey = ""
	ip = ""
	notify_url = ""
	out_trade_no = ""
)

// 测试样例
func TestPayment() {
	// 创建实例并初始化
	o := &RequestUnifiedOrder{
		AppId:appId,
		Openid:"",		// 这个应当作为参数传过来
		MchId:mch_id,
		NonceStr:GetNonceStr(32),				// 获取32个随机字符串 数字 + 大写字母
		Body:"Body",
		OutTradeNo:GetOutTradeNo(),				// 获取32位商户号, 时间戳 + 随机数
		TotalFee:1000,
		SpbillCreateIp:ip,
		NotifyUrl: notify_url,
		TradeType:"JSAPI",
	}
	// 调用统一下单接口 获取prepay_id 和 nonce_str
	resp,err := UnifiedOrder(o)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 统一下单支付成功后, 要及时将订单存储到数据库中

	// 生成返回给客户端的信息
	res := CreateRequestPayment(resp.NonceStr, resp.PrepayId)
	fmt.Println(*res)
}

func TestOrderQuery() {
	o := &RequestOrderQuery{
		AppId:appId,
		MchId:mch_id,
		OutTradeNo:out_trade_no,
		NonceStr:GetNonceStr(32),
		SignType:"MD5",
	}

	resp,err := OrderQuery(o)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*resp)
}


// 统一下单接口
func UnifiedOrder(o *RequestUnifiedOrder) (*ResponseUnifiedOrder, error) {

	// 将结构体转换为map,
	// 键-对应 struct 中的tag xml值,
	// 值-对应 struct 中的字段值
	mp := Struct2Map(*o)
	// 根据map 获取签名
	o.Sign = Sign(mp, apiKey)

	// 将结构体转换为xml格式, 并调用微信小程序统一下单接口
	byteReq,err := xml.Marshal(o)
	if err != nil {
		return nil,err
	}
	strReq := strings.Replace(string(byteReq), "RequestUnifiedOrder", "xml", -1)
	req,err := http.NewRequest("POST", UnifiedOrderURL, bytes.NewReader([]byte(strReq)))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	client := http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	// 读取返回数据到响应结构体中
	resp_data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	res := ResponseUnifiedOrder{}
	if err = xml.Unmarshal(resp_data, &res); err != nil {
		return nil, err
	}
	// 判断返回值
	if res.ReturnCode == "FAIL" {
		return nil,errors.New("微信支付统一下单失败,原因:" + res.ReturnCode + " " + res.ReturnMsg)
	}
	if res.ResultCode == "FAIL" {
		return nil,errors.New("执行业务失败,原因:" + res.ResultCode + " " + res.ErrCode + " " + res.ErrCodeDes)
	}



	return &res,nil
}

// 查询订单接口
func OrderQuery(o *RequestOrderQuery) (*ResponseOrderQuery, error) {

	mp := Struct2Map(*o)
	o.Sign = Sign(mp, apiKey)

	byteReq,err := xml.Marshal(o)
	strReq := strings.Replace(string(byteReq), "RequestOrderQuery", "xml", -1)
	if err != nil {
		return nil,errors.WithMessagef(err, "OrderQuery marshal struct to xml err")
	}

	req,err := http.NewRequest("POST", OrderQueryURL, bytes.NewReader([]byte(strReq)))
	if err != nil {
		return nil,errors.WithMessagef(err, "OrderQuery create request err")
	}
	defer req.Body.Close()

	client := http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		return nil, errors.WithMessagef(err, "OrderQuery get response err")
	}

	byteResp,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessagef(err, "OrderQuery read response data err")
	}

	res := ResponseOrderQuery{}
	if err := xml.Unmarshal(byteResp, &res); err != nil {
		return nil, errors.WithMessagef(err, "OrderQuery unmarshal data err")
	}

	if res.ReturnCode == "FAIL" {
		return nil, errors.Errorf("Request Failed, return_code :%s, return_msg:%s", res.ReturnCode, res.ReturnMsg)
	}
	if res.ResultCode == "FAIL" {
		return nil, errors.Errorf("Request Failed, err_code :%s, err_code_des:%s", res.ErrCode, res.ErrCodeDes)
	}

	fmt.Println("查询订单成功: ")
	fmt.Println(res)

	return &res,nil


}

// 关闭订单接口
func CloseOrder(o *RequestCloseOrder) (*ResponseCloseOrder, error) {
	mp := Struct2Map(*o)
	o.Sign = Sign(mp, apiKey)

	byteReq,err := xml.Marshal(o)
	if err != nil {
		return nil, errors.WithMessagef(err, "CloseOrder marshal struct to byte err")
	}

	req,err := http.NewRequest("POST", CloseOrderUrl, bytes.NewReader(byteReq))
	if err != nil {
		return nil, errors.WithMessagef(err, "CloseOrder create request err")
	}
	defer req.Body.Close()
	client := http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		return nil,errors.WithMessagef(err, "CloseOrder get response err")
	}

	byteResp,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessagef(err, "CloseOrder read response err")
	}
	res := ResponseCloseOrder{}
	if err = xml.Unmarshal(byteResp, &res); err != nil {
		return nil, errors.WithMessage(err, "CloseOrder unmarshal response err")
	}

	if res.ReturnCode == "FAIL" {
		return nil, errors.Errorf("return_code: %s, return_msg:%s", res.ReturnCode, res.ReturnMsg)
	}
	if res.ResultCode == "FAIL" {
		return nil, errors.Errorf("result_code: %s, result_msg:%s", res.ResultCode, res.ResultMsg)
	}

	return &res, nil

}


// 获取签名
func Sign(mp map[string]interface{}, key string) string {
	// 获取键的集合
	keys := make([]string, 0)
	for k,_ := range mp {
		keys = append(keys, k)
	}
	// 排序
	sort.Strings(keys)
	// 组合字符串
	buff := strings.Builder{}
	for _,key := range keys {
		value := fmt.Sprintf("%v", mp[key])
		if value != "" {
			buff.WriteString(key + "=" + value + "&")
		}
	}
	if key != "" {
		buff.WriteString("key=" + key)
	}
	data := buff.String()

	// 加密
	encrypt := md5.New()
	encrypt.Write([]byte(data))
	upperSign := strings.ToUpper(hex.EncodeToString(encrypt.Sum(nil)))

	return upperSign

}

// 前端小程序调起接口所需要的参数
func CreateRequestPayment(nonce_str, prepay_id string) ( p*RequestPayment) {
	p  = &RequestPayment{}

	p.AppId = appId
	p.NonceStr = nonce_str
	p.Package = "prepay_id=" + prepay_id
	p.TimeStamp = fmt.Sprintf("%d", time.Now().Unix())
	p.SignType = "MD5"

	str := "appId=%s&nonceStr=%s&package=%s&signType=MD5&timeStamp=%s&key=%s"
	p.PaySign = GetMD5Encode(fmt.Sprintf(str, p.AppId, p.NonceStr, p.Package, p.TimeStamp, apiKey))
	return
}
