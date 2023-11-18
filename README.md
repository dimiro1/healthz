# Healthz

Healthz is a straightforward Go package for adding health check HTTP handlers to your web apps. It's ideal for
Kubernetes-like environments to ensure your service is running smoothly.

# Features

Just two functions - simple and to the point.

# Installation

Run:

```shell
go get github.com/dimiro1/healthz
```

# Usage

## Custom Health Check

```go
package main

import (
	"net/http"
	"github.com/dimiro1/healthz"
)

func main() {
	healthCheckLogic := func() bool {
		// Define your health check logic here
		return true // return true if healthy, false otherwise
	}

	http.Handle("/healthz", healthz.Check(healthCheckLogic))
	_ = http.ListenAndServe(":8080", nil)
}

```

## Always Healthy

For a basic always-healthy endpoint:

```go
package main

import (
	"net/http"
	"github.com/dimiro1/healthz"
)

func main() {
	http.HandleFunc("/healthz", healthz.AlwaysUp)
	_ = http.ListenAndServe(":8080", nil)
}

```

# LICENSE

This project is licensed under the MIT License.
