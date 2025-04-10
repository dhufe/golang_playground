package main

func main() {
	productItem := newItem("Dummy Shirt")

	observerFirst := &Customer{id: "test_customer@acme.org"}
	observerSecond := &Customer{id: "example@gmail.com"}

	productItem.register(observerFirst)
	productItem.register(observerSecond)

	productItem.updateAvailability()
}
