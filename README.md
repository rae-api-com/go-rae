# rae-api.com Client

A Go client library for the rae-api.com, which provides programmatic access to word definitions, meanings, and verb conjugations from the Real Academia Española's (RAE) dictionary.

## Overview

Lightweight Go client for the [http://rae-api.com](RAE-API) service. It allows you to easily retrieve definitions, meanings, and verb conjugations from the RAE dictionary in a clean, structured JSON format.

> **Note**: This is **not** an official RAE client, and usage of the rae-api.com is subject to the API's terms and conditions.

## Installation

```bash
go get github.com/rae-api-com/go-rae
```

## Usage

### Basic Usage

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rae-api-com/go-rae"
)

func main() {
	// Create a new client
	c := client.New()
	
	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Fetch word data
	entry, err := c.Word(ctx, "hablar")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	
	// Print result
	data, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting output: %s\n", err)
		os.Exit(1)
	}
	
	fmt.Println(string(data))
}
```

### Command Line Usage

The repository also includes a simple command-line tool for quick lookups:

```bash
# Build the CLI tool
go build -o rae-cli

# Look up a word
./rae-cli hablar
```

## Response Structure

The API returns structured data in the following format:

```json
{
  "word": "hablar",
  "meanings": [
    {
      "origin": {
        "raw": "Del lat. comedĕre.",
        "type": "lat",
        "text": "comedĕre"
      },
      "senses": [
        {
          "meaning_number": 1,
          "raw": "1. tr. Masticar...",
          "category": "verb",
          "usage": "common",
          "description": "Masticar y deglutir un alimento sólido.",
          "synonyms": ["masticar", "etc."],
          "antonyms": []
        }
      ],
      "conjugations": {
        "non_personal": {
          "infinitive": "hablar",
          "gerund": "comiendo",
          "participle": "comido"
        },
        "indicative": {
          "present": {
            "singular_first_person": "como",
            // ... more conjugations
          }
        }
        // ... more verb tenses
      }
    }
  ]
}
```

## Features

- **Multiple Meanings** - Handles words that have more than one acepción (e.g., *hablar<sup>1</sup>* and *hablar<sup>2</sup>*).
- **Rich Definitions** - Access to synonyms, antonyms, usage labels (e.g., "colloquial", "desusado", etc.), and grammatical info.
- **Complete Verb Conjugations** - Get all tenses (Indicative, Subjunctive, Imperative) and non-personal forms (infinitive, gerund, participle).
- **Simple API** - Clean, easy-to-use interface for retrieving word data.

## License

[MIT License](LICENSE)

## Acknowledgements

This client uses the rae-api.com service, which is not affiliated with the Real Academia Española. All dictionary content belongs to the RAE and is subject to their terms and conditions.