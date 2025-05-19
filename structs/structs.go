package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	id    string
	name  string
	stock int
	price float32
}

type Order struct {
	id        string
	item      Item
	value     float32
	status    string
	createdAt time.Time
}

// (o Order) => receiver type that binds this function to Order struct
func newOrder(id string, items *[]Item) *Order {
	intId, _ := strconv.Atoi(id)
	order := Order{
		id:        id,
		item:      (*items)[intId-1],
		value:     (*items)[intId-1].price,
		createdAt: time.Now(),
	}
	(*items)[intId-1].stock--
	return &order // always return a pointer to the struct to maintain reference to it.
}

func newItem(id string, name string, stock int, price float32) *Item {
	item := Item{
		id:    id,
		name:  name,
		stock: stock,
		price: price,
	}
	return &item
}

func getItemsFromInventory() []string {
	return []string{"Shampoo", "Tea Powder", "Phone Stand", "Earphone", "Milk", "Shoe"}
}

func createItems(items []string) *[]Item {
	itemInventory := make([]Item, len(items))
	for i, item := range items {
		itemObject := newItem(strconv.Itoa(i+1), item, 10, 100.00)
		itemInventory[i] = *itemObject
	}
	return &itemInventory
}

func displayItems(items []Item) {
	fmt.Println("ITEMS AVAILABLE")
	for _, item := range items {
		fmt.Printf("%v. %v -> stock(%v), price(%v)\n", item.id, item.name, item.stock, item.price)
	}
}

func buyItem(items *[]Item, id int) *Order {
	// create new order
	item := (*items)[id-1]
	order := newOrder(item.id, items)
	return order
}

func buyInloop(items *[]Item) bool {
	displayItems(*items)
	fmt.Println("\nEnter Number to buy or type 0 to exit")
	reader := bufio.NewReader(os.Stdin)
	id, _ := reader.ReadString('\n')
	if strings.TrimSpace(id) == "0" {
		return false
	}
	intId, _ := strconv.Atoi(strings.TrimSpace(id))
	order := buyItem(items, intId)
	fmt.Println("Order Created ->", *order)
	return true
}

func main() {
	// creating initial list of items
	items := createItems(getItemsFromInventory())
	for {
		bought := buyInloop(items)
		if !bought {
			break
		}
	}
}
