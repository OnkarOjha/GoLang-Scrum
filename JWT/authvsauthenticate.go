// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im9ua2FyIiwiZXhwIjoxNjc3NDc5NzYzfQ.5qfBB4gzZ9cjrZ0d3zvW9DWwUzn_wfEtGk2iIwtwxWIfteCaszutNaBG7_dylbhjRdEiVItzbeOd-5_YTjiJYQLDPx-7dXzFgUbZXe0OVJRyWXtV_9k0_86h1Xe9ZyAqRosWMSsYMGe2Md5hZ_AyrWnrOwpJoG94kil_ink33HgW6H9RBhIV6zrCgiPkxj0KCU
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func main() {
	// Example JWT token
	//^ token for authentication
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im9ua2FyIiwiZXhwIjoxNjc3NDc5NzYzfQ.5qfBB4gzZ9cjrZ0d3zvW9DWwUzn_wfEtGk2iIwtwxWIfteCaszutNaBG7_dylbhjRdEiVItzbeOd-5_YTjiJYQLDPx-7dXzFgUbZXe0OVJRyWXtV_9k0_86h1Xe9ZyAqRosWMSsYMGe2Md5hZ_AyrWnrOwpJoG94kil_ink33HgW6H9RBhIV6zrCgiPkxj0KCU"

	// ^ token for authorization
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXSwiaWF0IjoxNTE2MjM5MDIyfQ.cw-CyRjQz0shZ5G5dYbf5Zr02MsqX1tEpv1tFdhWwRw"

	// ^ token is signed or unsigned

	// Parse the token without verifying the signature
	token, _, _ := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if token.Valid {
		if token.Signature == "" {
			// The token is unsigned
			fmt.Println("Token is unsigned")
		} else {
			// The token is signed
			fmt.Println("Token is signed")
		}
	} else {
		// The token is invalid
		fmt.Println("Token is invalid")
	}

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	fmt.Println("token: ", token) // token infromation hai

	if err != nil {
		fmt.Println("Failed to parse JWT token:", err)

	}

	// Check if token contains "sub" (subject) claim
	if sub, ok := token.Claims.(jwt.MapClaims)["username"].(string); ok {
		// If "sub" claim is present, the token is likely used for authentication
		fmt.Println("Token is used for authentication, name:", sub)
	} else {
		// If "sub" claim is not present, check if token contains "roles" claim
		if roles, ok := token.Claims.(jwt.MapClaims)["roles"].([]interface{}); ok {
			// If "roles" claim is present, the token is likely used for authorization
			strRoles := make([]string, len(roles))
			for i, r := range roles {
				strRoles[i] = r.(string)
			}
			fmt.Println("Token is used for authorization, roles:", strings.Join(strRoles, ","))
		} else {
			// If neither "sub" nor "roles" claim is present, the token is invalid
			fmt.Println("Invalid JWT token")
		}
	}
}
