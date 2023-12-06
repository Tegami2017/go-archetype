package dao

import (
	"go_archetype/core/model/entity"
	"go_archetype/core/storage"
)

type UserDao struct {
}

var User = &UserDao{}

func (r *UserDao) SelectById(id *int) *entity.User {
	sql := `select id, name, age, country, created_at 
			from user
			where id = ?`
	return queryOne(&entity.User{}, sql, *id)

}

func (r *UserDao) ListByIds(ids []int) *[]entity.User {
	sql := `select id, name, age, country, created_at
			from user
			where id IN ?`
	return queryRows(&[]entity.User{}, sql, ids)

}

func (r *UserDao) Create(user *entity.User) {
	storage.DB.Omit("Id", "CreatedAt", "UpdatedAt").Create(user)
}

func (r *UserDao) BatchCreate(users *[]entity.User) {
	storage.DB.Select("Name", "Age", "Country").Create(users)
}
