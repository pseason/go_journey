package model

// 分页参数
type PageInfo struct {
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
	PageNum  int `form:"pageNum" json:"pageNum" binding:"required"`
}
