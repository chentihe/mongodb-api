package types

type MongoPaginate struct {
	Page  int64 `form:"page" binding:"min=1" json:"page"`
	Limit int64 `form:"limit" json:"limit"`
}

func NewMongoPaginate(page int, limit int) *MongoPaginate {
	return &MongoPaginate{
		Page:  int64(page),
		Limit: int64(limit),
	}
}
