package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/syllabix/deepai"

	"github.com/syllabix/deepai/option"
)

type method func(context.Context, io.Reader) (interface{}, error)
type registry map[string]method

func register(c *deepai.Client) map[string]method {
	return registry{
		"nsfw": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.DetectNudity(ctx, r)
		},
		"moderation": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.Moderate(ctx, r)
		},
		"face": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.RecognizeFaces(ctx, r)
		},
		"demographics": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.GetFaceDemographics(ctx, r)
		},
		"summarize": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.SummarizeText(ctx, r)
		},
		"sentiment": func(ctx context.Context, r io.Reader) (interface{}, error) {
			return c.AnalyzeSentiment(ctx, r)
		},
	}
}

func main() {
	apiKey := flag.String("key", "", "a valid deep ai api key")
	filename := flag.String("f", "", "a valid path to an image file")
	api := flag.String("api", "", "an identifier for a deepai api (nsfw, moderation, face, etc...)")
	flag.Parse()

	ctx := context.Background()

	client := deepai.NewClient(*apiKey, option.Timeout(time.Second*10))
	registry := register(client)

	method, registered := registry[strings.ToLower(*api)]
	if !registered {
		fmt.Println("")
		fmt.Printf("Unsupported API '%s' provided to client\n", *api)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Please use one of the following supported deep api aliases:")
		fmt.Println()
		for a := range registry {
			fmt.Println("#", a)
		}
		fmt.Println()
		os.Exit(1)
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	response, err := method(ctx, file)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(payload))

}
