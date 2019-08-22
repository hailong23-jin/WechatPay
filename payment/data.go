package payment

const (
	UnifiedOrderURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryURL =   "https://api.mch.weixin.qq.com/pay/orderquery"
	CloseOrderUrl =   "https://api.mch.weixin.qq.com/pay/closeorder"
)

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
	MchId string 			`xml:"mch_id"`				// 是	商户号
	DeviceInfo string 		`xml:"device_info"`
	NonceStr string 		`xml:"nonce_str"`			// 是	随机字符串
	Sign string 			`xml:"sign"`				// 是	签名
	SignType string 		`xml:"sign_type"`
	Body string 			`xml:"body"`				// 是	商品简单描述
	Detail string 			`xml:"detail"`
	Attach string 			`xml:"attach"`
	OutTradeNo string 		`xml:"out_trade_no"`		// 是	商户系统内部订单号
	FeeType string 			`xml:"fee_type"`
	TotalFee int	 		`xml:"total_fee"`			// 是	订单总金额 单位[分]
	SpbillCreateIp string 	`xml:"spbill_create_ip"`	// 是	终端IP 调用微信支付API的机器IP
	TimeStart string 		`xml:"time_start"`
	TimeExpire string 		`xml:"time_expire"`
	GoodsTag string 		`xml:"goods_tag"`
	NotifyUrl string 		`xml:"notify_url"`			// 是	回调地址 外网可访问
	TradeType string 		`xml:"trade_type"`			// 是	JSAPI
	ProductId string 		`xml:"product_id"`
	LimitPay string 		`xml:"limit_pay"`
	Openid string 			`xml:"openid"`
	Receipt string 			`xml:"receipt"`
	StoreInfo string 	 	`xml:"store_info"`
}

type StoreInfo struct {
	Id string				`json:"id"`
	Name string				`json:"name"`
	AreaCode string		`json:"area_code"`
	Address string			`json:"address"`
}

// 统一下单接口响应结构体
type ResponseUnifiedOrder struct {						// 是否必须
	ReturnCode	string		`xml:"return_code"`			// 是	返回状态码	SUCCESS/FAIL
	ReturnMsg	string		`xml:"return_msg"`

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId	string			`xml:"appid"`				// 是	小程序id
	MchId	string			`xml:"mch_id"`				// 是 	商户号
	DeviceInfo	string		`xml:"device_info"`
	NonceStr	string		`xml:"nonce_str"`			// 是	随机字符串
	Sign	string			`xml:"sign"`				// 是	签名
	ResultCode	string		`xml:"result_code"`			// 是	业务结果		SUCCESS/FAIL
	ErrCode	string			`xml:"err_code"`
	ErrCodeDes	string		`xml:"err_code_des"`

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType	string		`xml:"trade_type"`			// 是	交易类型		JSAPI
	PrepayId	string		`xml:"prepay_id"`			// 是	预支付交易会话标识
	CodeUrl	string			`xml:"code_url"`
}

// 订单查询请求结构体
type RequestOrderQuery struct {							// 是否必须
	AppId string			`xml:"appid"`				// 是
	MchId string			`xml:"mch_id"`				// 是
	TransactionId string	`xml:"transaction_id"`		//
	OutTradeNo string		`xml:"out_trade_no"`		//  与上面的二选一
	NonceStr string			`xml:"nonce_str"`			// 是
	Sign string				`xml:"sign"`				// 是
	SignType string			`xml:"sign_type"`
}

// 订单查询响应结构体
type ResponseOrderQuery struct {						// 是否必须
	ReturnCode string		`xml:"return_code"`			// 是
	ReturnMsg string		`xml:"return_msg"`			// 是

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId string			`xml:"appid"`				// 是
	MchId string			`xml:"mch_id"`				// 是
	NonceStr string			`xml:"nonce_str"`			// 是
	Sign string				`xml:"sign"`				// 是
	ResultCode string		`xml:"result_code"`			// 是
	ErrCode	string			`xml:"err_code"`
	ErrCodeDes	string		`xml:"err_code_des"`

	// 	以下字段在return_code 、result_code、trade_state都为SUCCESS时有返回
	// 	，如trade_state不为 SUCCESS，则只返回out_trade_no（必传）和attach（选传）。
	DeviceInfo string			`xml:"device_info"`
	OpenId string				`xml:"openid"`			// 是
	IsSubscribe string			`xml:"is_subscribe"`	// 是 是否关注公众账号 Y-关注  N-未关注
	TradeType string			`xml:"trade_type"`		// 是 JSAPI，NATIVE，APP，MICROPAY
	TradeState string			`xml:"trade_state"`		// 是
	/*
	SUCCESS—支付成功
	REFUND—转入退款
	NOTPAY—未支付
	CLOSED—已关闭
	REVOKED—已撤销（刷卡支付）
	USERPAYING--用户支付中
	PAYERROR--支付失败(其他原因，如银行返回失败)
	*/
	BankType string				`xml:"bank_type"`				// 是 付款银行
	TotalFee int				`xml:"total_fee"`				// 是 标价金额
	SettlementTotalFee int		`xml:"settlement_total_fee"`
	FeeType string				`xml:"fee_type"`
	CashFee int					`xml:"cash_fee"`				// 是 现金支付金额 【分】
	CashFeeType string			`xml:"cash_fee_type"`
	CouponFee int				`xml:"coupon_fee"`
	CouponCount int				`xml:"coupon_count"`
	CouponTypen string			`xml:"coupon_type_$n"`
	CouponIdn string			`xml:"coupon_id_$n"`
	CouponFeen int				`xml:"coupon_fee_$n"`
	TransactionId string		`xml:"transaction_id"`			// 是
	OutTradeNo string			`xml:"out_trade_no"`			// 是
	Attach string				`xml:"attach"`
	TimeEnd string				`xml:"time_end"`				// 是 支付完成时间
	TradeStateDesc string		`xml:"trade_state_desc"`		// 是 交易状态描述
}


// 关闭订单请求结构体
type RequestCloseOrder struct {
	AppId string			`xml:"appid"`				// 是
	MchId string			`xml:"mch_id"`				// 是
	OutTradeNo string		`xml:"out_trade_no"`		//  与上面的二选一
	NonceStr string			`xml:"nonce_str"`			// 是
	Sign string				`xml:"sign"`				// 是
	SignType string			`xml:"sign_type"`
}

// 关闭订单响应结构体
type ResponseCloseOrder struct {
	ReturnCode	string		`xml:"return_code"`			// 是	返回状态码	SUCCESS/FAIL
	ReturnMsg	string		`xml:"return_msg"`

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId	string			`xml:"appid"`				// 是	小程序id
	MchId	string			`xml:"mch_id"`				// 是 	商户号
	NonceStr	string		`xml:"nonce_str"`			// 是	随机字符串
	Sign	string			`xml:"sign"`				// 是	签名
	ResultCode	string		`xml:"result_code"`			// 是	业务结果		SUCCESS/FAIL
	ResultMsg string		`xml:"result_msg"`			// 是	业务结果描述
	ErrCode	string			`xml:"err_code"`
	ErrCodeDes	string		`xml:"err_code_des"`
}
























