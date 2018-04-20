package persist

import (
	"context"
	"errors"
	"log"

	"github.com/kiyonlin/newworld/crawler/engine"
	elastic "gopkg.in/olivere/elastic.v5"
)

// ItemSaver returns a channel to save and print items
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in doker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-out
			log.Printf("Item saver got Item: #%d: %v", itemCount, item)
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item saver :error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.ID != "" {
		indexService.Id(item.ID)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
