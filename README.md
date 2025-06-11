# Go Wrapper for [Fortnite-API.com](https://fortnite-api.com)

[![Go Reference](https://pkg.go.dev/badge/github.com/BurakYs/go-fortnite-api-com.svg)](https://pkg.go.dev/github.com/BurakYs/go-fortnite-api-com)
[![Release](https://img.shields.io/github/v/release/BurakYs/go-fortnite-api-com)](https://github.com/BurakYs/go-fortnite-api-com/releases)
[![License](https://img.shields.io/github/license/BurakYs/go-fortnite-api-com)](LICENSE)
[![Discord](https://img.shields.io/discord/621452110558527502?label=Discord&logo=discord)](https://discord.gg/eysmvFT2rV)

A simple Go wrapper for [Fortnite-API.com](https://fortnite-api.com), with support for all available endpoints with context-based requests.

## Installation

```
go get github.com/BurakYs/go-fortnite-api-com
```

## Quick Example

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	fortniteapi "github.com/BurakYs/go-fortnite-api-com"
)

func main() {
	client := fortniteapi.NewClient(fortniteapi.LanguageEnglish, "your-api-key")

	flags := fortniteapi.CombineFlags(
		fortniteapi.FlagIncludePaths,
		fortniteapi.FlagIncludeGameplayTags,
	)

	searchParams := fortniteapi.BRCosmeticSearchParams{
		Name:          "Peely",
		ResponseFlags: flags,
	}

	cosmetics, err := client.SearchBRCosmetic(context.TODO(), searchParams)
	if err != nil {
		log.Fatalln(err)
	}

	cosmeticsJSON, _ := json.MarshalIndent(cosmetics, "", "  ")
	fmt.Println(string(cosmeticsJSON))
}
```

## Links
- [API Documentation](https://dash.fortnite-api.com)
- [pkg.go.dev](https://pkg.go.dev/github.com/BurakYs/go-fortnite-api-com)
- [Discord](https://discord.gg/eysmvFT2rV)

## License
This project is licensed under the [MIT License](LICENSE).