package command

type Account struct {
	NickName   string `json:"nickname"`
	Sex        string `json:"sex"`
	Telephone  string `json:"telephone"`
	ProfileUrl string `json:"profileurl"`
	Password   string `json:"password"`
}
