package handler

type RootHandler struct {
	User          *UserHandler
	CheckUsername *CheckUsernameHandler
	Register      *RegisterHandler
}

// TODO User: add, login, get, check_username

func NewRootHandler(
	user *UserHandler,
	checkUsername *CheckUsernameHandler,
	register *RegisterHandler,
) *RootHandler {
	return &RootHandler{
		User:          user,
		CheckUsername: checkUsername,
		Register:      register,
	}
}
