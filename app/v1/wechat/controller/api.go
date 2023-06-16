package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"main.go/app/v1/enroll/model/EnrollModel"
	"main.go/app/v1/wechat/model/WechatOrderModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func ApiController(route *gin.RouterGroup) {

	route.Any("notify", api_notify)
}

func api_notify(c *gin.Context) {
	err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mchCertificateSerialNumber, mchID, mchAPIv3Key)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return
	}
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return
	}
	transaction := new(payments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(context.Background(), c.Request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return
	}
	// 处理通知内容
	fmt.Println(notifyReq.Summary)
	order_id := transaction.OutTradeNo
	//fmt.Println(transaction.TransactionId)
	order := WechatOrderModel.Api_find_orderId(order_id)
	data := EnrollModel.Api_find_byOrderId(order_id)
	if len(order) > 1 {
		data = EnrollModel.Api_find(order["relative_id"])
	}
	if len(data) < 1 {
		c.JSON(200, map[string]any{
			"code":    "FAIL",
			"message": "未找到订单号",
		})
		return
	}
	if *transaction.TradeState == "SUCCESS" {
		var enroll EnrollModel.Interface
		enroll.Db = tuuz.Db()
		//enroll.Api_update_isPayed(order_id, 1)
		enroll.Api_update_isPayed_byId(data["id"], 1)
		EnrollModel.Api_update_orderId(data["id"], order_id)
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
	} else if *transaction.TradeState == "NOTPAY" {
		EnrollModel.Api_update_orderId(data["id"], "")
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
	} else if *transaction.TradeState == "CLOSED" {
		var enroll EnrollModel.Interface
		enroll.Db = tuuz.Db()
		enroll.Api_update_isPayed(order_id, -1)
		EnrollModel.Api_update_orderId(data["id"], "")
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
	} else if *transaction.TradeState == "REVOKED" {
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
		EnrollModel.Api_update_orderId(data["id"], "")
	} else if *transaction.TradeState == "USERPAYING" {
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
	} else if *transaction.TradeState == "PAYERROR" {
		EnrollModel.Api_update_orderId(data["id"], "")
		WechatOrderModel.Api_update_status(order_id, -1, *transaction.TradeState)
	}
	c.JSON(200, map[string]any{
		"code":    "SUCCESS",
		"message": "成功",
	})
}
