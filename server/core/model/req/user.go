package req

type CreateUserReq struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Country string `json:"country" binding:"required"`
}
