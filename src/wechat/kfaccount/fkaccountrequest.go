package kfaccount

type Kfaccountrequest struct {
	Kf_account string `json:"kf_account"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`

	RequestJSON string `json:"-"`
}
