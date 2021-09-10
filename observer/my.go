package main

import "fmt"

type subject interface {
	register(id string)
	deregister(id string)
	notifyAll()
}

type item struct {
	notifyFuncs map[string]func(string)
	name        string
	inStock     bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(id string) {
	if i.notifyFuncs == nil {
		i.notifyFuncs = map[string]func(string){}
	}
	i.notifyFuncs[id] = customerNotificationFunc(id)
}

func (i *item) deregister(id string) {
	delete(i.notifyFuncs, id)
}

func (i *item) notifyAll() {
	for _, notifyFunc := range i.notifyFuncs {
		notifyFunc(i.name)
	}
}

func customerNotificationFuncs(ids []string) map[string]func(string) {
	funcs := make(map[string]func(string))
	for _, id := range ids {
		funcs[id] = customerNotificationFunc(id)
	}
	return funcs
}

func customerNotificationFunc(id string) func(string) {
	return func(itemName string) {
		fmt.Printf("Sending email to customer %s for item %s\n", id, itemName)
	}
}

func main() {
	shirtItem := newItem("Nike Shirt")
	shirtItem.register("Tom")
	shirtItem.register("Jerry")
	shirtItem.updateAvailability()
}
