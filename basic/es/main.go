package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("conn es success")
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Id("1").
		BodyJson(tweet).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("insert success")
}
