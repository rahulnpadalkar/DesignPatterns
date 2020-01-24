package main

import (
	pub "designpatterns/behavioural/observer/publisher"
	sub "designpatterns/behavioural/observer/subscriber"
	"fmt"
)

func main() {

	p := pub.Publisher{
		Subscribers: []*sub.Subscriber{},
	}

	sub1 := sub.Subscriber{
		Name: "sub1",
	}
	sub2 := sub.Subscriber{
		Name: "sub2",
	}
	sub3 := sub.Subscriber{
		Name: "sub3",
	}

	p.AddSubscriber(&sub1)
	p.AddSubscriber(&sub2)
	p.AddSubscriber(&sub3)

	p.Publish(1)
	p.Publish(2)
	p.Publish(3)

	fmt.Println("Done publishing!")

}
