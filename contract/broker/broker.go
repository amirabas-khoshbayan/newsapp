package broker

import "newsapp/entity"

type Publisher interface {
	Publish(event entity.Event, payload string)
}
