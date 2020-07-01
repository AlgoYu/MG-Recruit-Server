package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"machine-geek.cn/recruit-server/database"
	"machine-geek.cn/recruit-server/model"
	"time"
)

const (
	SECRET = "QQ794763733"
	EXPIRE = time.Hour * 24
)

// 加入Token信息到Redis
func addAdminToRedis(v *model.Admin) {
	database.RDB.HSet("admin"+string(v.Id), "id", v.Id)
	database.RDB.HSet("admin"+string(v.Id), "name", v.Name)
	database.RDB.HSet("admin"+string(v.Id), "picture", v.Picture)
	database.RDB.Expire("admin"+string(v.Id), EXPIRE)
}

// 获取Admin Token
func getAdminToken(v *model.Admin) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   v.Id,
		"name": v.Name,
	})
	tokenString, err := token.SignedString([]byte(SECRET))
	if err == nil {
		return tokenString
	} else {
		//beego.Error(err.Error())
	}
	return ""
}

// JWT解析函数
func parse(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(SECRET), nil
}
