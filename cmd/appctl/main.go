package main

import (
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// cfg := config.LoadConfig()
	// kafka := kafka.NewKafkaPublisher(cfg)

	// event := event.NewEvent(kafka)

	// event.

	src, err := imaging.Open("testdata/test_image.jpeg")

	if err != nil {
		log.Fatalf("failed to get image: %v", err)
	}

	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	err = imaging.Save(src, "testdata/output.jpeg")

	if err != nil {
		log.Fatalf("failed to process image: %v", err)
	}
}
