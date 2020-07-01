package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"machine-geek.cn/recruit-server/common"
	"machine-geek.cn/recruit-server/database"
	"machine-geek.cn/recruit-server/model"
	"machine-geek.cn/recruit-server/service"
	"time"
)

const (
	SECRET = "QQ794763733"
	EXPIRE = time.Hour * 24
)

func LoginAdmin(c *gin.Context) {
	var input model.Admin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(200, common.Fail("Json格式错误！"))
		return
	}

	admin := service.GetAdminByName(input.Name)
	if admin != nil {
		if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)) == nil {
			token := getAdminToken(admin)
			addAdminToRedis(admin)
			c.JSON(200, common.Success(token))
			return
		}
	}
	c.JSON(200, common.Fail("账户密码不正确！"))
}

func LogoutAdmin(c *gin.Context) {
}

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
