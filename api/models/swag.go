package models

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type ConfigInput struct {
	Key   string      `json:"key"`
	Type  string      `json:"type"`
	Value interface{} ``
}

type ConfigUpdate struct {
	Type  string      `json:"type"`
	Value interface{} ``
}

type GetConfigsByKey struct {
	Keys [] string `json:"keys" example:"co:test,co:array,co:json"`
}

type AuthToken struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ExampleGetConfig struct {
	Success bool `json:"success"`
	Config ConfigFormatted 
}

type ExampleGetConfigs struct {
	Success bool `json:"success"`
	Config []ConfigFormatted 
}

type ExampleSuccessConfig struct {
	Success bool `json:"success"`
}


