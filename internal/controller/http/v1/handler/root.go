package handler

type RootHandler struct {
	User *UserHandler
}

// TODO User: add, login, get, check_username

func NewRootHandler(
	user *UserHandler,
) *RootHandler {
	return &RootHandler{
		User: user,
	}
}
