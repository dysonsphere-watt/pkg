package pkg

type ctxkey string

const (
	UserIDHeaderKey string = "user-id"
	KeyUserID       ctxkey = "userID"
)
