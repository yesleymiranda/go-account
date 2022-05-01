package decoder

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func IDUint(r *http.Request) (uint, error) {
	v := mux.Vars(r)
	idStr := v["id"]
	if idStr == "" {
		return 0, errors.New("id is required")
	}

	ids, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || ids == 0 {
		return 0, errors.New("id is invalid")
	}

	return uint(ids), nil
}
