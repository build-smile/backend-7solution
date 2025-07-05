package middleware

type Middleware struct {
	secretKey string
}

func NewMiddleware(secretKey string) *Middleware {
	return &Middleware{
		secretKey: secretKey,
	}
}
