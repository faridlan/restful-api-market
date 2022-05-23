package model

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (v *NullString) MarshalJSONString() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.String)
}
