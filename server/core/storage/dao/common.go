package dao

import (
	r "go_archetype/core/storage"
)

func queryOne[T any](dest *T, sql string, args ...interface{}) *T {
	result := r.DB.Raw(sql, args...).Scan(dest)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return dest
}

func queryRows[T any](dest *[]T, sql string, args ...interface{}) *[]T {
	result := r.DB.Raw(sql, args...).Scan(dest)
	if result.Error != nil {
		panic(result.Error)
	}
	return dest
}
