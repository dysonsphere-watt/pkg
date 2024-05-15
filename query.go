package pkg

import (
	"fmt"

	"github.com/goravel/framework/contracts/database/orm"
)

type PageInfo struct {
	Page     uint64
	PageSize uint64
}

// Helper function to allow users to write query functions faster
func QueryAllPaginator(query orm.Query, keyWord string, keyFields []string, pageInfo *PageInfo, resultPtr any, totalPtr *int64) error {
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

	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	err := query.Paginate(int(pageInfo.Page), int(pageInfo.PageSize), resultPtr, totalPtr)
	if err != nil {
		return err
	}

	return nil
}
