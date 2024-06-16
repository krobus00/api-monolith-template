package request

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=30,unique_db=users:username"`
	Email    string `json:"email" binding:"required,email,unique_db=users:email"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}
