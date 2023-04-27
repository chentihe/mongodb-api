package types

type MongoPaginate struct {
	Page  int64
	Limit int64
}

func NewMongoPaginate(page int, limit int) *MongoPaginate {
	return &MongoPaginate{
		Page:  int64(page),
		Limit: int64(limit),
	}
}
