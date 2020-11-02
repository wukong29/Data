package controllers

import (
	"BeegoDemo/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

/**
 * 证书详细信息查看页面控制器
 */
type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get() {
	//0、获取前端页面get请求时携带的cert_id数据
	certId := c.GetString("cert_id")
	fmt.Println("要查询的认证ID：", certId)
	//1、准备数据：根据cert_id到区块链上查询具体的信息，获得到区块信息
	block, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		c.Ctx.WriteString("链上数据查询错误！")
		return
	}
	//查询未遇到错误，有两情况：查到和未查到
	if block == nil {
		c.Ctx.WriteString("抱歉，未查询到链上数据，请重试！")
		return
	}
	//certId = hex.EncodeToString(block.Data)
	c.Data["CertId"] = strings.ToUpper(string(block.Data))
	//2、跳转页面
	c.TplName = "cert_detail.html" //显示并跳转到详情页面
}
