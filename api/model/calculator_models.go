package model

type BaseResult struct {
	First  int `json:"A"`
	Second int `json:"B"`
}
type AddResult struct {
	BaseResult
	Result int `json:"ADD"`
}

type MulResult struct {
	BaseResult
	Result int `json:"Mul"`
}
