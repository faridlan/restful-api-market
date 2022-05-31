package helper

import (
	"database/sql"

	"github.com/faridlan/restful-api-market/model"
)

func NewNullString(s string) *model.NullString {
	if len(s) == 0 {
		return &model.NullString{}
	}

	return &model.NullString{
		NullString: sql.NullString{
			String: s,
			Valid:  true,
		},
	}
}
