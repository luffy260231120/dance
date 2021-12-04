package model

type UserAuth struct {
	UserID int    `json:"user_id" binding:"required,max=128"`
	Token  string `json:"token" binding:"required,max=256"`
}

type UserAuthWithoutToken struct {
	UserID int    `json:"user_id" binding:"required,max=128"`
	Token  string `json:"token"`
}

type Pagination struct {
	Page int `json:"page" binding:"required"`
	Size int `json:"size" binding:"required,min=1,max=10"`
}

type PaginationAcc struct {
	Page int `json:"page" binding:"required"`
	Size int `json:"size" binding:"required,min=1,max=50"`
}
