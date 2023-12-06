package converter

import (
	"go_archetype/core/model/entity"
	"go_archetype/core/model/vo"
)

type UserConverter struct{}

var UserConv = &UserConverter{}

func (ins UserConverter) ToVO(item *entity.User) *vo.User {
	if item == nil {
		return nil
	}

	return &vo.User{
		Id:        item.Id,
		Name:      item.Name,
		Age:       item.Age,
		Country:   item.Country,
		CreatedAt: item.CreatedAt,
	}
}
