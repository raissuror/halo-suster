package web

type UserLoginReq struct {
	Nip      string `validate:"required" json:"nip"`
	Password string `validate:"required" json:"password"`
}
