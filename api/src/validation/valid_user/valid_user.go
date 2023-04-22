package validation

import (
	"net/http"
)

func ValidUser(r *http.Request, w http.ResponseWriter, ID uint64) bool {
	//userIdToken, err := auth.GetIdToken(r)
	//fmt.Printf(strconv.FormatUint(userIdToken, 10), "token")
	//fmt.Println(err, "this is err")
	//if err != nil {
	//	response.Err(w, http.StatusUnauthorized, err)
	//	return false
	//}
	//
	//if userIdToken != ID {
	//	response.Err(w, http.StatusForbidden, err)
	//	return false
	//
	//}
	return true
}
