package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type PageInfo struct {
	Page     int64
	PageSize int64
}

const dbCachePrefix = "WATT_DB_CACHE_"

func GetPageInfo(reqPageInfo interface{}) (*PageInfo, error) {
	val := reflect.ValueOf(reqPageInfo)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return &PageInfo{Page: 1, PageSize: 10}, nil
		}

		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, received a %s", val.Kind().String())
	}

	pageField := val.FieldByName("Page")
	if !pageField.IsValid() || pageField.Type().Kind() != reflect.Int64 {
		return nil, errors.New("struct does not contain an int64 'Page' field")
	}

	pageSizeField := val.FieldByName("PageSize")
	if !pageSizeField.IsValid() || pageSizeField.Type().Kind() != reflect.Int64 {
		return nil, errors.New("struct does not contain an int64 'Page Size' field")
	}

	pageInfo := PageInfo{
		Page:     Ternary(pageField.Int() == 0, 1, pageField.Int()),
		PageSize: Ternary(pageSizeField.Int() == 0, 10, pageSizeField.Int()),
	}
	return &pageInfo, nil
}

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
