package controller

import (
	"crypto/rsa"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"golang.org/x/net/context"
	"log"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func PayController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("index", pay_index)
	route.Any("order", pay_order)

}

var mchID string
var appid string
var mchCertificateSerialNumber string
var mchAPIv3Key string

var mchPrivateKey *rsa.PrivateKey

func init() {
	_ready()
	_ready_key()
}

func _ready() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready()
	} else {
		value, err := cfg.GetSection("wechatpay")
		if err != nil {
			cfg.SetValue("wechatpay", "appid", "")
			cfg.SetValue("wechatpay", "mchID", "")
			cfg.SetValue("wechatpay", "mchCertificateSerialNumber", "")
			cfg.SetValue("wechatpay", "mchAPIv3Key", "")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("wechatpay_ready")
			_ready()
		}
		appid = value["appid"]
		mchID = value["mchID"]
		mchCertificateSerialNumber = value["mchCertificateSerialNumber"]
		mchAPIv3Key = value["mchAPIv3Key"]
	}
}

func _ready_key() {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	privatekey, err := utils.LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}
	mchPrivateKey = privatekey
	ctx = context.Background()
	// 使用商户私钥等初始化 clients，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	clients, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay clients err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: clients}
	resp, result, err := svc.DownloadCertificates(ctx)
	client = clients
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
}

var ctx context.Context

func pay_index(c *gin.Context) {
	_ready_key()
}

var client *core.Client

func pay_order(c *gin.Context) {
	orderid := Calc.GenerateOrderId()
	svc := jsapi.JsapiApiService{Client: client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(appid),
			Mchid:       core.String(mchID),
			Description: core.String("参加活动的邮费"),
			OutTradeNo:  core.String(orderid),
			//Attach:      core.String("自定义数据说明"),
			NotifyUrl: core.String("https://api.gallery.familyeducation.org.cn/v1/wechat/api/notify"),
			Amount: &jsapi.Amount{
				Total: core.Int64(100),
			},
			Payer: &jsapi.Payer{
				Openid: core.String("oRrdQt76e1d17oYlUF3Bf-PaQlBU"),
			},
		},
	)
	if err == nil {
		RET.Success(c, 0, resp, nil)
	} else {
		RET.Fail(c, 500, err, err.Error())
	}
}
