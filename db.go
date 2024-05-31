package pkg

import (
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type PageInfo struct {
	Page     int64
	PageSize int64
}

const dbCachePrefix = "WATT_DB_CACHE_"

// Helper function to allow users to write query functions faster
func QueryAllPaginator(query orm.Query, keyWord string, keyFields []string, pageInfo *PageInfo, resultPtr any, totalPtr *int64) error {
	if keyWord != "" {
		wildcardKW := "%" + keyWord + "%"
		for _, field := range keyFields {
			query = query.OrWhere(fmt.Sprintf("%s LIKE ?", field), wildcardKW)
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

func CachedOrmQuerySingle(id string, query orm.Query, result any) error {
	key := dbCachePrefix + id

	if facades.Cache().Has(key) {
		res := facades.Cache().Get(key)
		result = &res
		return nil
	}

	err := query.First(result)
	if err != nil {
		return err
	}

	facades.Cache().Put(key, result, time.Hour*6)
	return nil
}

func CachedOrmQueryMulti(id string, query orm.Query, result any) error {
	key := dbCachePrefix + id

	if facades.Cache().Has(key) {
		res := facades.Cache().Get(key)
		result = &res
		return nil
	}

	err := query.Find(result)
	if err != nil {
		return err
	}

	facades.Cache().Put(key, result, time.Hour*6)
	return nil
}
