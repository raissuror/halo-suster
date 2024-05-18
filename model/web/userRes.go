package web

type UserRes struct {
	UserId string `json:"id"`
	Nip    string `json:"nip"`
	Name   string `json:"name"`
	Token  string `json:"accessToken"`
}
