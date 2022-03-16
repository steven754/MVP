package v1

import (
	"MVP/models"
	"MVP/pkg/e"
	"MVP/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//登录，注册接口
func UserStore(c *gin.Context) {
	mobile := c.PostForm("mobile") //取出请求参数手机号mobile
	vCode := c.PostForm("code")    //取出请求参数验证码code

	//对请求对参数进行验证
	validate := validation.Validation{}
	validate.Required(mobile, "Mobile").Message("手机号有误") //因为测试环境，所以使用了Required方法，正式下使用Mobile方法，做手机号校验。
	validate.Length(vCode, 4, "Code").Message("验证码格式不正确")
	//校验错误，有错误返回json
	if isOk := checkValidation(&validate, c); isOk == false { //校验不通过
		return
	}

	//数据库根据手机号查询用户信息
	user, err := models.FindUserByMobile(mobile)
	if gorm.IsRecordNotFoundError(err) { //如果数据库没有查询到没有用户信息,代表要注册，新创建用户信息
		user, err = models.CreateUser(mobile)
		if err != nil {
			util.ResponseWithJson(e.ERROR, "数据库操作错误", c)
			return
		}
	} else {
		if err != nil {
			util.ResponseWithJson(e.ERROR, "数据库操作错误", c)
			return
		}
	}
	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User": user,
	}, c)
	token, err := util.GeterateToken(user.ID, user.Mobile)
	if err != nil {
		util.ResponseWithJson(e.ERROR, "创建token失败", c)
		return
	}
	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User":  user,
		"Token": token,
	}, c)
}

/*检查请求参数是否有错误，如果有的话返回false*/
func checkValidation(vali *validation.Validation, c *gin.Context) bool {
	if vali.HasErrors() { //请求的参数有误
		var errs []string                 //创建一个保存错误信息的数组
		for _, err := range vali.Errors { //遍历错误信息数组，把错误信息添加到数组当中
			errs = append(errs, err.Message)
		}
		util.ResponseWithJson(e.INVALID_PARAMS, errs, c) //返回客户端错误信息
		return false
	}
	return true
}
