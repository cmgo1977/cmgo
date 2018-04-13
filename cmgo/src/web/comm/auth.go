package comm

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//create token
func CreateJwtToken() string {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = "a"                                   		//uid
	claims["mobile"] = "b"                                		//手机
	claims["iss"] = C.Auth.ISS                        		//签发主体
	claims["iat"] = time.Now()                            		//签发时间
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 		//1天
	//claims["exp"] = time.Now().Add(time.Microsecond).Unix() 	//立即过期

	//生成token并将其作为响应发送
	encodeToken, err := token.SignedString([]byte(C.Auth.SecretKey))
	if err != nil {
		log.Fatalf("redigo->RedigoPool->redis.Dial()初始化连接池时报错: %s\n", err)
	}

	return encodeToken
}

