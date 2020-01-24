package subscriber

import "fmt"

type Subscriber struct {
	Name string
}

func (s *Subscriber) GetValue(number int) {

	fmt.Printf("Subscriber: %v got value %v\n", s.Name, number)

}
