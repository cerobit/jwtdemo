package  main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func parseToken(){
hmacSampleSecret := []byte("your-256-bit-secret")
tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return hmacSampleSecret, nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["name"], claims["iat"])
} else {
    fmt.Println(err)
}
}

func  srvlogin(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a seeerver response")
}

func main(){
	parseToken()
    http.HandleFunc("/login", srvlogin)
    http.ListenAndServe(":8080", nil)
}