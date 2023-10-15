package processors

import (
	"errors"
	"sub/internals/app/cache"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFindOrderByIDSuccessFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCache := cache.NewCache()
	storage := NewMockOrdersStorageInterface(ctrl)
	ordersProcessor := NewOrdersProcessor(storage, testCache)

	expect := `{"testJson":"testJson"}`

	storage.EXPECT().GetOrderByID("b563feb7b2b84b6test").Return(expect, nil)

	res, err := ordersProcessor.GetOrderByID("b563feb7b2b84b6test")

	if err != nil {
		t.Errorf("fail to get order: %s", err)
		return
	}

	if res != expect {
		t.Errorf("result not match, want: %s, got %s", expect, res)
		return
	}
}

func TestFindOrderByIDSuccessFromCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCache := cache.NewCache()
	storage := NewMockOrdersStorageInterface(ctrl)
	ordersProcessor := NewOrdersProcessor(storage, testCache)

	expect := `{"testJson":"testJson"}`

	ordersProcessor.ordersCache.AddOrderToCache("b563feb7b2b84b6test", expect)

	res, err := ordersProcessor.GetOrderByID("b563feb7b2b84b6test")

	if err != nil {
		t.Errorf("fail to get order: %s", err)
		return
	}

	if res != expect {
		t.Errorf("result not match, want: %s, got %s", expect, res)
		return
	}
}

func TestFindOrderByIDFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCache := cache.NewCache()
	storage := NewMockOrdersStorageInterface(ctrl)
	ordersProcessor := NewOrdersProcessor(storage, testCache)

	expect := ""

	storage.EXPECT().GetOrderByID("b563feb7b2b84b6test").Return(expect, errors.New("sql: no rows in result set"))

	res, err := ordersProcessor.GetOrderByID("b563feb7b2b84b6test")

	if err == nil {
		t.Errorf("must return error: sql: no rows in result set")
	}

	if res != expect {
		t.Errorf("result not match, want: %s, got %s", expect, res)
		return
	}
}
