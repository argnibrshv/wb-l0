package subscribers

import (
	"sub/internals/app/cache"
	"sub/internals/app/db"
	"sub/internals/app/validation"

	stan "github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
)

// обертка для подключения к nats-streaming с параметрами сервера.
type Subscriber struct {
	sc          *stan.Conn
	storage     *db.OrdersStorage
	ordersCache *cache.Cache
}

func NewSubscriber(sc *stan.Conn, storage *db.OrdersStorage, ordersCache *cache.Cache) *Subscriber {
	subscriber := new(Subscriber)
	subscriber.storage = storage
	subscriber.sc = sc
	subscriber.ordersCache = ordersCache
	return subscriber
}

func (s *Subscriber) Subscribe() (stan.Subscription, error) {
	return (*s.sc).Subscribe("order-sub", s.messageHandlerFunc(), stan.DurableName("my-durable"), stan.DeliverAllAvailable(), stan.SetManualAckMode())
}

// Функция для обработки сообщений из nats-streaming с параметрами сервера.
func (s *Subscriber) messageHandlerFunc() stan.MsgHandler {
	return func(msg *stan.Msg) {
		//валидируем входящий json в соответствии с моделью
		id, data, err := validation.JSONValidation(msg.Data)
		if err != nil {
			log.Errorln(err)

			if err.Error() == "order id is empty" {
				if err := msg.Ack(); err != nil {
					return
				}
			}

			return
		}

		_, ok := s.ordersCache.GetOrderByID(id)
		if ok {
			log.Errorf("order with id: %s already in cache", id)

			if err := msg.Ack(); err != nil {
				return
			}
			return
		}

		err = s.storage.AddOrderToDB(id, data)
		if err != nil {
			if err.Error() == `ERROR: duplicate key value violates unique constraint "orders_pkey" (SQLSTATE 23505)` {
				log.Errorf("order with id: %s already in database", id)

				if err := msg.Ack(); err != nil {
					return
				}
			}

			return
		}

		s.ordersCache.AddOrderToCache(id, string(data))
		log.Printf("order with id: %s added in cache and database", id)

		if err := msg.Ack(); err != nil {
			return
		}
	}
}
