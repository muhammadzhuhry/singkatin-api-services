package request

type Url struct {
	LongUrl string `validate:"required,url" json:"url"`
}
