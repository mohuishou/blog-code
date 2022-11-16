package domain

// Tag 标签数据
type Tag struct {
	Model

	Key   string `json:"key"`
	Value string `json:"value"`
}
