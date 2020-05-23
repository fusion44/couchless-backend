package context

type key int

// Keys for middlewares
const (
	KeyUserloaderMiddleware key = iota
	KeyCurrentUser
	KeyAppConfig
	KeyLogger
)

const (
	// DefaultFITFileDir the directory where fit files will be stored
	DefaultFITFileDir = "fit"
	// DefaultImageFileDir the directory where image files will be stored
	DefaultImageFileDir = "images"
)
