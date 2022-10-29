package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"learngo2/crawal/engine"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}

		}
	}()
	return out
}

func save(item engine.Item) (err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	indexService := client.Index().Index("dating_profile").
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)
	return nil
}
