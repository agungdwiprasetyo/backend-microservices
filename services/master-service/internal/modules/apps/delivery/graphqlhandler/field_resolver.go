package graphqlhandler

import (
	"monorepo/services/master-service/internal/modules/apps/domain"
	shareddomain "monorepo/services/master-service/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/candishared"
)

// CommonFilter  basic filter model
type CommonFilter struct {
	Limit   *int    `json:"limit" validate:"omitempty,gte=1"`
	Page    *int    `json:"page" validate:"omitempty,gte=1"`
	Search  *string `json:"search"`
	Sort    *string `json:"sort" validate:"omitempty,oneof='desc' 'asc'"`
	ShowAll *bool   `json:"show_all"`
	OrderBy *string `json:"order_by"`
}

// ToSharedFilter method
func (f *CommonFilter) ToSharedFilter() (filter domain.FilterApps) {
	filter.Search = candihelper.PtrToString(f.Search)
	filter.OrderBy = candihelper.PtrToString(f.OrderBy)
	filter.Sort = candihelper.PtrToString(f.Sort)
	filter.ShowAll = candihelper.PtrToBool(f.ShowAll)

	if f.Limit == nil {
		filter.Limit = 10
	} else {
		filter.Limit = *f.Limit
	}
	if f.Page == nil {
		filter.Page = 1
	} else {
		filter.Page = *f.Page
	}

	return
}

// AppListResolver resolver
type AppListResolver struct {
	Meta candishared.Meta
	Data []shareddomain.Apps
}
