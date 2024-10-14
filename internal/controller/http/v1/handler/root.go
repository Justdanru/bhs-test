package handler

type RootHandler struct {
	User     *UserHandler
	Register *RegisterHandler
}

// TODO User: add, login, get, check_username

func NewRootHandler(
	user *UserHandler,
	register *RegisterHandler,
) *RootHandler {
	return &RootHandler{
		User:     user,
		Register: register,
	}
}
