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
	list, total, err = model.ListSearch(search)
	return list, total, err
}
