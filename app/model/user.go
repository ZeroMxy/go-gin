package model

type User struct {
	PKIncrModel `xorm:"extends"`
	Type 		int	   `json:"type"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Nickname    string `json:"nickname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Gender      int    `json:"gender"`
	Age         int    `json:"age"`
	Avatar      string `json:"avatar"`
	Remark      string `json:"remark"`
	LastIp		string `json:"lastIp"`
	Status      int    `json:"status"`
}

type UserRole struct {
	User    `xorm:"extends"`
	RoleId   int    `json:"roleId"`
	RoleName string `json:"roleName"`
}
