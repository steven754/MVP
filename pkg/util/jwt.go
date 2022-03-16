package util

import (
	"MVP/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret) //jwt密钥

/*在jwt中添加的自定义用户信息，有用户的ID和手机号,也可以加一些其它信息*/
type Claims struct {
	ID uint
	//Mobile string //这里没有使用手机号。因为token字符串可以被解析。。。。
	jwt.StandardClaims
}

/*生成Token*/
func GeterateToken(id uint, mobile string) (string, error) {
	nowTime := time.Now()                                                   //当前时间
	expireTime := nowTime.Add(setting.AppSetting.JwtExpireTime * time.Hour) //过期时间，为了测试这里是3小时后过期

	//设置自定义荷载
	claims := Claims{
		id,
		//mobile,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //该方法内部生成签名字符串，再用于获取完整、已签名的token

	return token, err
}

/*校验和解析token*/
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
