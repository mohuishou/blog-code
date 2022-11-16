package v1

import (
	"encoding/json"
	errors "errors"
)

// UnmarshalJSON sets *m to a copy of data.
func (m *ListArticleTagsResp) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}

	return json.Unmarshal(data, &m.Tags)
}

// MarshalJSON returns m as the JSON encoding of m.
func (m *ListArticleTagsResp) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}

	return json.Marshal(m.Tags)
}
