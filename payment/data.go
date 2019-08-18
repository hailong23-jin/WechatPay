package payment

// 定义小程序调用 所需要使用到的结构体

// 小程序调起支付数据签名字段列表
type RequestPayment struct {
	AppId string			`xml:"appid"`
	TimeStamp string		`xml:"timeStamp"`
	NonceStr string			`xml:"nonceStr"`
	Package string			`xml:"package"`
	PaySign string			`xml:"paySign"`
	SignType string			`xml:"signType"`
}

// 统一下单接口请求结构体
type RequestUnifiedOrder struct {						// 是否必须
	AppId string 			`xml:"appid"`				// 是	小程序id
	Mch_id string 			`xml:"mch_id"`				// 是	商户号
	Device_info string 		`xml:"device_info"`
	Nonce_str string 		`xml:"nonce_str"`			// 是	随机字符串
	Sign string 			`xml:"sign"`				// 是	签名
	Sign_type string 		`xml:"sign_type"`
	Body string 			`xml:"body"`				// 是	商品简单描述
	Detail string 			`xml:"detail"`
	Attach string 			`xml:"attach"`
	Out_trade_no string 	`xml:"out_trade_no"`		// 是	商户系统内部订单号
	Fee_type string 		`xml:"fee_type"`
	Total_fee int	 		`xml:"total_fee"`			// 是	订单总金额 单位[分]
	Spbill_create_ip string `xml:"spbill_create_ip"`	// 是	终端IP 调用微信支付API的机器IP
	Time_start string 		`xml:"time_start"`
	Time_expire string 		`xml:"time_expire"`
	Goods_tag string 		`xml:"goods_tag"`
	Notify_url string 		`xml:"notify_url"`			// 是	回调地址 外网可访问
	Trade_type string 		`xml:"trade_type"`			// 是	JSAPI
	Product_id string 		`xml:"product_id"`
	Limit_pay string 		`xml:"limit_pay"`
	Openid string 			`xml:"openid"`
	Receipt string 			`xml:"receipt"`
	Store_info string 	 	`xml:"store_info"`
}

type StoreInfo struct {
	Id string				`json:"id"`
	Name string				`json:"name"`
	Area_code string		`json:"area_code"`
	Address string			`json:"address"`
}

// 统一下单接口响应结构体
type ResponseUnifiedOrder struct {						// 是否必须
	Return_code	string		`xml:"return_code"`		// 是	返回状态码	SUCCESS/FAIL
	Return_msg	string		`xml:"return_msg"`

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId	string			`xml:"appid"`				// 是	小程序id
	Mch_id	string			`xml:"mch_id"`				// 是 	商户号
	Device_info	string		`xml:"device_info"`
	Nonce_str	string		`xml:"nonce_str"`			// 是	随机字符串
	Sign	string			`xml:"sign"`				// 是	签名
	Result_code	string		`xml:"result_code"`		// 是	业务结果		SUCCESS/FAIL
	Err_code	string		`xml:"err_code"`
	Err_code_des	string	`xml:"err_code_des"`

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Trade_type	string		`xml:"trade_type"`			// 是	交易类型		JSAPI
	Prepay_id	string		`xml:"prepay_id"`			// 是	预支付交易会话标识
	Code_url	string		`xml:"code_url"`
}

