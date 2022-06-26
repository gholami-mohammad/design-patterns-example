package main

import (
	"fmt"
	"reflect"
)

type NotificationObservable interface {
	OnMessage(message string)
}

type NotificationSubject struct {
	observers []NotificationObservable
}

func (sub *NotificationSubject) Subscribe(observable NotificationObservable) {
	sub.observers = append(sub.observers, observable)
}

func (sub *NotificationSubject) Unsubscribe(observable NotificationObservable) {
	observesLen := len(sub.observers)

	for k, ob := range sub.observers {
		if !reflect.DeepEqual(ob, observable) {
			continue
		}

		if observesLen == 1 {
			sub.observers = []NotificationObservable{}
			return
		}

		// observers count >= 2

		// k is the first index
		if k == 0 {
			sub.observers = sub.observers[1:]
			return
		}

		// k is the last index
		if k == observesLen-1 {
			sub.observers = sub.observers[:observesLen-1]
			return
		}

		sub.observers = append(sub.observers[:k], sub.observers[k+1:]...)
		return

	}
}

func (sub *NotificationSubject) Notify(message string) {
	for _, obs := range sub.observers {
		obs.OnMessage(message)
	}
}

type UserObserver struct {
	id int64
}

func (user UserObserver) OnMessage(message string) {
	fmt.Printf("OnMessage called from Object with ID: %d Message: %s\n", user.id, message)
}

func main() {
	observer1 := UserObserver{id: 1}
	observer2 := UserObserver{id: 2}
	observer3 := UserObserver{id: 3}

	subject := new(NotificationSubject)
	subject.Subscribe(observer1)
	subject.Subscribe(observer2)
	subject.Subscribe(observer3)
	subject.Notify("Mesage A")

	subject.Unsubscribe(observer2)
	subject.Notify("Mesage B")

	subject.Unsubscribe(observer3)
	subject.Notify("Mesage C")
}
