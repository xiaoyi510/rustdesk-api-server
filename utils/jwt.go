package utils

import "github.com/dgrijalva/jwt-go"

// 指定加密密钥
var jwtSecret = []byte("2d9a0da267bee9c14d8e7aaedeca907c")

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserId      int    `json:"id"`
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	jwt.StandardClaims
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
