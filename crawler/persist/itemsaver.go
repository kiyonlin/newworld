package persist

import "log"

// ItemSaver returns a channel to save and print items
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-out
			log.Printf("Item saver got Item: #%d: %v", itemCount, item)
		}
	}()
	return out
}
