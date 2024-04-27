package common

type Page struct {
	PageNum  int `validate:"gte=1"`
	PageSize int `validate:"gte=0"`
	Offset   int
	KeyWord  interface{}
}
