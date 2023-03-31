package service

import (
	"paper/internal/model"
)

type AdminUserService struct {
	AdminModel *model.Admin
}

func (a AdminUserService) AdminUserListPage(createTimeStart, createTimeEnd int64, name string, page, size int) (list *[]model.Admin, total int64, err error) {
	var search model.ListSearchParams
	search.Page = page
	search.Size = size
	search.Name = name
	search.CreateTimeStart = createTimeStart
	search.CreateTimeEnd = createTimeEnd
	list, total, err = model.AdminUserListSearch(search)
	return list, total, err
}

func (a AdminUserService) EditAdminUser(id int, name, password string) error {
	adminUser := model.Admin{
		Id:   id,
		Name: name,
	}

	if password != "" {
		adminUser.Password = password
	}

	return adminUser.EditAdmin()
}

func (a AdminUserService) DelAdminUser(id int) error {
	adminUser := model.Admin{
		Id: id,
	}
	return adminUser.DelAdmin()
}
