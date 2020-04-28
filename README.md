# deepai
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/deepai)](https://goreportcard.com/report/github.com/syllabix/deepai)

an unofficial go client and cli for the deep ai api.

### client usage

You will need a valid api key issued from deepai to use the client. Checkout https://deepai.org/ for signup details

```
go get github.com/syllabix/deepai
```

```go
    ctx := context.Background()

    // expose your api key - perhaps through an env var
    var apiKey = os.Getenv("DEEPAI_KEY")

    client := deepai.NewClient(apiKey, option.Timeout(time.Second*10))

    // all client methods expect an io.Reader - so your image can come from anywhere
    // in this case - we load a local potentially scandalous file on disk...
    img, err := os.Open("path/to/nsfw/image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

    response, err := client.DetectNudity(ctx, img)
    if err != nil {
        if errors.Is(err, deepai.ErrUnauthorized) {
            // go get an api key
        }
    }

    // examine the response...
```

For a more complete example - check out cli in the `cmd/deepai` package.

### cli usage

```
go get -u github.com/syllabix/deepai/cmd/deepai
```

The cli is thinly wrapped interface around the client itself using the following flags:

```
  -api string
        an identifier for a deepai api (nsfw, moderation, face)
  -f string
        a valid path to an image file
  -key string
        a valid deep ai api key
```









