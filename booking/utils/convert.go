package utils

import "database/sql"

type NullStr string

// String s > 0, Valid = true， 反之false
func (s NullStr) String() sql.NullString {
	if len(s) > 0 {
		return sql.NullString{String: string(s), Valid: true}
	}
	return sql.NullString{String: "", Valid: false}
}

type NullInt32 int32

// Int32 n > 0, Valid = true, 反之false
func (n NullInt32) Int32() sql.NullInt32 {
	if n > 0 {
		return sql.NullInt32{Int32: int32(n), Valid: true}
	}
	return sql.NullInt32{Int32: 0, Valid: true}
}
