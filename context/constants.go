package context

type key int

// Keys for middlewares
const (
	KeyUserloaderMiddleware key = iota
	KeyCurrentUser
	KeyAppConfig
	KeyLogger
)
