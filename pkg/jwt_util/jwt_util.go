/**
 * @Author pibing
 * @create 2020/12/27 12:14 PM
 */

package jwt_util

import (
	"github.com/dgrijalva/jwt-go"
)
const SECRET = "pibing"

//把对象生成token
func MakeToken(obj map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(obj))
	tokenString, err := token.SignedString([]byte(SECRET))
	return tokenString, err
}
//解析token为对象
func ParseToken(tokenStr string) map[string]interface{} {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil
	}
	finToken := token.Claims.(jwt.MapClaims)
	return finToken
}
