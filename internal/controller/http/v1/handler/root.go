package handler

type RootHandler struct {
	User          *UserHandler
	CheckUsername *CheckUsernameHandler
	Register      *RegisterHandler
	Login         *LoginHandler
}

func NewRootHandler(
	user *UserHandler,
	checkUsername *CheckUsernameHandler,
	register *RegisterHandler,
	login *LoginHandler,
) *RootHandler {
	return &RootHandler{
		User:          user,
		CheckUsername: checkUsername,
		Register:      register,
		Login:         login,
	}
}
