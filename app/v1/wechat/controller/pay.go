package controller

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"golang.org/x/net/context"
	"log"
)

func PayController(route *gin.RouterGroup) {

}

var mchID string
var mchCertificateSerialNumber string
var mchAPIv3Key string

func init() {
	_ready()
}

func _ready() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready()
	} else {
		value, err := cfg.GetSection("wechatpay")
		if err != nil {
			cfg.SetValue("wechatpay", "mchID", "")
			cfg.SetValue("wechatpay", "mchCertificateSerialNumber", "")
			cfg.SetValue("wechatpay", "mchAPIv3Key", "")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("wechatpay_ready")
			_ready()
		}
		mchID = value["mchID"]
		mchCertificateSerialNumber = value["mchCertificateSerialNumber"]
		mchAPIv3Key = value["mchAPIv3Key"]
	}
}

func pay_index(c *gin.Context) {

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
}
