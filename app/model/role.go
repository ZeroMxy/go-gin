package model

type Role struct {
	PKIncrModel
	Name      string     `json:"name"`
	Mark      string     `json:"mark"`
	Sort      int        `json:"sort"`
	Remark    string     `json:"remark"`
	Status    int        `json:"status"`
}
