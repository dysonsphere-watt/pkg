package db

import "github.com/goravel/framework/contracts/database/orm"

// Helper function to create pagination queries

// query Is the existing query for the table, with .Model() or .Where() already called.
// dest will be a reference to the model output.
func Paginate(query orm.Query, page, pageSize int, dest any) (int32, int32, error) {
	var total int64

	err := query.Paginate(page, pageSize, dest, &total)
	if err != nil {
		return -1, -1, err
	}

	totalItems := int32(total)
	totalPages := int32((int(total) + pageSize - 1) / pageSize)

	return totalItems, totalPages, nil
}
