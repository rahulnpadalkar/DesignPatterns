package publisher

import (
	sub "designpatterns/behavioural/observer/subscriber"
)

type Publisher struct {
	Subscribers []*sub.Subscriber
}

func (p *Publisher) Publish(number int) {

	for _, sub := range p.Subscribers {
		sub.GetValue(number)
	}

}

func (p *Publisher) AddSubscriber(s *sub.Subscriber) {

	p.Subscribers = append(p.Subscribers, s)
}
