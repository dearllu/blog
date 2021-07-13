package chitchat

import "net/http"

func anthenticate(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.p
}
