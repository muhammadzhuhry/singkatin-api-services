package request

type Url struct {
	LongUrl string `validate:"required" json:"url"`
}
