// package main

// import (
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func main() {
// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name", Hello)

// 	http.ListenAndServe(":8080", router)
// }

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Write([]byte("Welcome!\n"))
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	name := ps.ByName("Onkar")
// 	w.Write([]byte("Hello, " + name + "!\n"))
// }


package auth

import (
	"fmt"
	"main/models"
	cons "main/utils"
	"net/http"

	jwt "github.com/golang-jwt/jwt"
)


func IsAuthorized(endpoint func(http.ResponseWriter,*http.Request))http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		if r.Header["Token"] != nil {

			token,err:=jwt.ParseWithClaims(r.Header["Token"][0],&models.Claims{},func(token *jwt.Token) (interface{},error){

				if _,ok:=token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil,fmt.Errorf("error")
				}
				return cons.Jwt_key,nil
			})


			if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
				
				if claims.Role=="admin"{

					endpoint(w,r)
				}
			} else {
				fmt.Println(err)
			}

	}else{

		fmt.Fprint(w,"NOt authorized")
	}
	
	})
}	