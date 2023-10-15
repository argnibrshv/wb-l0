package processors

type OrdersStorageInterface interface {
	GetOrderByID(id string) (result string, err error)
}

type OrdersCacheInterface interface {
	AddOrderToCache(id string, data string)
	GetOrderByID(id string) (string, bool)
}

type OrdersProcessor struct {
	storage     OrdersStorageInterface
	ordersCache OrdersCacheInterface
}

func NewOrdersProcessor(storage OrdersStorageInterface, ordersCache OrdersCacheInterface) *OrdersProcessor {
	processor := new(OrdersProcessor)
	processor.storage = storage
	processor.ordersCache = ordersCache
	return processor
}

func (o *OrdersProcessor) GetOrderByID(id string) (string, error) {
	data, ok := o.ordersCache.GetOrderByID(id)
	if !ok {
		data, err := o.storage.GetOrderByID(id)
		if err != nil {
			return data, err
		}
		o.ordersCache.AddOrderToCache(id, data)
		//если принципиально нужно из кэша вернуть данные
		data, _ = o.ordersCache.GetOrderByID(id)
		return data, nil
	}
	return data, nil
}
