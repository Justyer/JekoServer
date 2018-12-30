package weapon

type Weapon struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Attack   int    `json:"attack"`
	CurBlood int    `json:"cur_blood"`
	MaxBlood int    `json:"max_blood"`
}
