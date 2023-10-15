package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// В случае ошибок при выполении запроса, в ответ добавляем сообщение ошибки, оборачиваем  ответ в json и добавляем http заголовки к ответу

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}
