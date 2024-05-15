package pkg

import (
	"fmt"

	"github.com/goravel/framework/contracts/database/orm"
)

type PageInfo struct {
	Page     uint64
	PageSize uint64
}

type QueryRequest interface {
	GetKeyWord() string
	GetPageInfo() *PageInfo
}

// Helper function to allow users to write query functions faster
func QueryAllPaginator[T QueryRequest](query orm.Query, req T, keyFields []string, results *any, totalPtr *int64) error {
	keyWord, pageInfo := req.GetKeyWord(), req.GetPageInfo()

	if keyWord != "" {
		where := false
		wildcardKW := "%" + keyWord + "%"
		for _, field := range keyFields {
			if where {
				query = query.OrWhere(fmt.Sprintf("%s LIKE ?", field), wildcardKW)
			} else {
				query = query.Where(fmt.Sprintf("%s LIKE ?", field), wildcardKW)
			}
		}
	}

	if pageInfo == nil {
		pageInfo = &PageInfo{Page: 1, PageSize: 10}
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	err := query.Paginate(int(pageInfo.Page), int(pageInfo.PageSize), results, totalPtr)
	if err != nil {
		return err
	}

	return nil
}
