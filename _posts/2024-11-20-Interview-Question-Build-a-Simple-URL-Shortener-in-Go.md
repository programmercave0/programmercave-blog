---
layout: post
title: "Interview Question: Build a Simple URL Shortener in Go"
description: "Learn how to create a simple and thread-safe URL shortener in Go using `sync.Mutex` for safe concurrent access to an in-memory map. Discover the implementation of generating unique short codes, storing mappings, and retrieving original URLs."
keywords: "URL shortener, Go programming, sync.Mutex, thread-safe, in-memory map, concurrent access, Go concurrency, random string generation, short code generation, simple URL shortener in Go, scalable URL shortener. interview question, software engineer"
header-img: ""
tags:  [Go, Interview-Questions]
date: 2024-11-20
toc: true
---

In this blog post, we will build a simple URL shortener in Go. This will demonstrate how to generate short codes for long URLs, store them in memory, and retrieve the original URL from the short code. We will also ensure thread-safety for concurrent access to the in-memory store using `sync.Mutex`.

_This question was asked to me by Radisys for the role of Golang Developer._

---

### **Step 1: Data Structure for Storing URLs**

First, we need a place to store the URLs. Since we want our URL shortener to be fast and simple, we will use an in-memory map to store the mapping between the short codes and the original URLs. We also need a way to safely access this map in a multi-threaded environment, which is where `sync.Mutex` comes in.

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// A map to store URLs in memory
var urlStore = make(map[string]string)
var mu sync.Mutex // Mutex to handle concurrent access
```

- **`urlStore`**: A map that stores the short code as the key and the corresponding long URL as the value.
- **`mu`**: A mutex to ensure thread-safe access to the `urlStore`.

---

### **Step 2: Generate a Short Code**

To shorten the URL, we need to generate a short code. The short code will be a random alphanumeric string. We will define a function to generate this random string and ensure that it is unique before storing it.

```go
// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random string of fixed length (short code)
func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
```

- **`generateShortCode`**: This function generates a random string of a fixed length (6 characters in this case). It uses a predefined charset consisting of lowercase letters, uppercase letters, and digits.

---

### **Step 3: Store the URL Mapping**

Once we have a short code, we need to store the mapping between the short code and the long URL. To ensure uniqueness, we will check if the generated short code already exists in the `urlStore` map. If it does, we generate a new code.

```go
// ShortenURL takes a long URL and returns a short code
func ShortenURL(longURL string) string {
	mu.Lock()
	defer mu.Unlock()

	// Generate a unique short code
	shortCode := generateShortCode(6)
	for urlStore[shortCode] != "" { // Ensure uniqueness
		shortCode = generateShortCode(6)
	}

	// Store the mapping in memory
	urlStore[shortCode] = longURL
	return shortCode
}
```

- **`ShortenURL`**: This function takes the long URL as input, generates a short code, and stores the mapping in `urlStore`. If the generated short code already exists, it generates a new one until a unique code is found.

---

### **Step 4: Retrieve the Original URL**

To retrieve the original URL from the short code, we will define a function that looks up the short code in the `urlStore` map. If the short code exists, it will return the corresponding long URL.

```go
// GetOriginalURL retrieves the original URL from a short code
func GetOriginalURL(shortCode string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	longURL, exists := urlStore[shortCode]
	return longURL, exists
}
```

- **`GetOriginalURL`**: This function takes the short code as input and looks it up in the `urlStore`. If the short code exists, it returns the original URL; otherwise, it returns `false`.

---

### **Step 5: Putting It All Together**

Now that we have all the necessary functions, let's put everything together in the `main` function. We will demonstrate how to shorten a URL and retrieve the original URL using the generated short code.

```go
// Main for demonstration
func main() {
	// Shorten a URL
	longURL := "https://example.com"
	shortCode := ShortenURL(longURL)
	fmt.Printf("Short code for %s: %s\n", longURL, shortCode)

	// Retrieve the original URL
	originalURL, exists := GetOriginalURL(shortCode)
	if exists {
		fmt.Printf("Original URL for short code %s: %s\n", shortCode, originalURL)
	} else {
		fmt.Println("Short code not found!")
	}
}
```

- The `main` function demonstrates the entire process:
  - A long URL is shortened using `ShortenURL`.
  - The original URL is retrieved using `GetOriginalURL`.

#### **Example Output:**
```plaintext
Short code for https://example.com: aB3cD1
Original URL for short code aB3cD1: https://example.com
```

---

### **Conclusion**

In this post, we have built a simple URL shortener in Go. The key features of this implementation include:
- **Thread-Safety**: Using `sync.Mutex` ensures that the URL store can be accessed concurrently without race conditions.
- **Short Code Generation**: We used a random string generation function to create unique short codes.
- **URL Mapping**: The short codes are mapped to long URLs and stored in memory for fast access.

### **Advantages**
- **Simple and Fast**: The use of an in-memory map ensures fast lookups and code generation.
- **Thread-Safe**: The implementation is safe to use in a multi-goroutine environment.

### **Limitations**
- **Ephemeral Storage**: Since the data is stored in memory, it is lost when the program exits.
- **Scalability**: This approach may not scale well for large datasets (e.g., millions of URLs) as it uses in-memory storage.

---

### **Possible Extensions**
- **Persistent Storage**: You can integrate a database like MySQL or PostgreSQL to store the URL mappings permanently.
- **REST API**: You can expose the URL shortening and retrieval functionality via a REST API, making it accessible over the web.

