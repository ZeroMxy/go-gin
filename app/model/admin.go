package model

type Admin struct {
	PKIncrModel
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Nickname  string     `json:"nickname"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Gender    int        `json:"gender"`
	Age       int        `json:"age"`
	Avatar    string     `json:"avatar"`
	Remark    string     `json:"remark"`
	Status    int        `json:"status"`
}

type AdminRole struct {
	Admin    `xorm:"extends"`
	RoleId   int    `json:"roleId"`
	RoleName string `json:"roleName"`
}
