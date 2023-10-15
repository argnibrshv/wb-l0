package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"sub/internals/app/processors"

	"github.com/gorilla/mux"
)

type OrdersHandler struct {
	processor *processors.OrdersProcessor
}

func NewOrdersHandler(processor *processors.OrdersProcessor) *OrdersHandler {
	handler := new(OrdersHandler)
	handler.processor = processor
	return handler
}

func (o *OrdersHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	//получаем параметры пути.
	vars := mux.Vars(r)
	if vars["id"] == "" {
		//оборачиваем обшибку в json и отправляем ответ
		WrapError(w, errors.New("missing id"))
		return
	}

	order, err := o.processor.GetOrderByID(vars["id"])
	if err != nil {
		//оборачиваем обшибку в json и отправляем ответ
		WrapError(w, errors.New("order not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, order)
}
