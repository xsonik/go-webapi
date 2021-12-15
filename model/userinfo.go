package model

type UserInfo struct {
	UserCode   string `json:"user_code"`
	UserName   string `json:"user_name"`
	Filiale    string `json:"filiale"`
	ModuleFlag string `json:"module_flag"`
	OperHost   string `json:"oper_host"`
	OperIP     string `json:"oper_ip"`
}
