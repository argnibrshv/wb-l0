package validation

import (
	"encoding/json"
	"errors"
	"sub/internals/app/models"
)

func JSONValidation(data []byte) (id string, validData []byte, err error) {

	var order models.Order

	err = json.Unmarshal(data, &order)
	if err != nil {
		return id, validData, err
	}

	id = order.ID
	if len(id) == 0 {
		err = errors.New("order id is empty")
		return id, validData, err
	}

	validData, err = json.Marshal(&order)
	if err != nil {
		return id, validData, err
	}

	return id, validData, nil
}
