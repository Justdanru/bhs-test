package middleware

type RootMiddleware struct {
	Init *InitMiddleware
	Auth *AuthMiddleware
}

func NewRootMiddleware(
	init *InitMiddleware,
	auth *AuthMiddleware,
) *RootMiddleware {
	return &RootMiddleware{
		Init: init,
		Auth: auth,
	}
}
