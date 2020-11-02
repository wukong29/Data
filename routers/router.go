package routers

import (
	"BeegoDemo/controllers"
	"github.com/astaxie/beego"
)

/**
 * router.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求，用于匹配请求
 */
func init() {
	//注册页面
	beego.Router("/", &controllers.MainController{})
	//用户注册的接口请求
	beego.Router("/user_register", &controllers.RegisterController{})
	//直接登录的页面请求接口
	beego.Router("/login.html", &controllers.LoginController{})
	//用户登录请求接口
	beego.Router("/user_login",&controllers.LoginController{})
	//用户修改密码请求接口
	beego.Router("/pwd_password",&controllers.CgangController{})
	//登入页面请求注册
	beego.Router("/register_register",&controllers.RegisterController{})
	//文件上传接口
	beego.Router("/upload", &controllers.UploadFileController{})
	//查看认证数据列表
	beego.Router("/upload_file.html",&controllers.UploadFileController{})
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})
}
