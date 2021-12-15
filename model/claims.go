package model

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	UserCode   string `json:"user_code"`
	UserName   string `json:"user_name"`
	Filiale    string `json:"filiale"`
	ModuleFlag string `json:"module_flag"`
	OperHost   string `json:"oper_host"`
	OperIP     string `json:"oper_ip"`
	jwt.StandardClaims
}
