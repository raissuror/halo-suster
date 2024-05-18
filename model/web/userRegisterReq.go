package web

type UserRegisterReq struct {
	Nip      string `validate:"required" json:"nip"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
}
