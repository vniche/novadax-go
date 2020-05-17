# novadax-go #

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/vniche/novadax-go?tab=doc)

novadax-go is a Go client library for accessing the [NovaDAX API v1](https://doc.novadax.com/).

## Usage ##

```go
import "github.com/vniche/novadax-go"
```

Construct a new NovaDAX client, then use the various services on the client to
access different parts of the NovaDAX API. For example:

```go
// LISTING SYMBOLS (PUBLIC ENDPOINT)

import (
    novadax "github.com/vniche/novadax-go"
)

client := novadax.Default()

// list all symbols available at NovaDAX
symbols, err := client.ListSymbols()
```

```go
// LISTING ORDERS (PRIVATE ENDPOINT, THUS REQUIRE ACCESS AND SECRET KEYS)

import (
    novadax "github.com/vniche/novadax-go"
)

/*
* This configuration is also possible via environment variables, eg.:
* NOVADAX_ACCESS_KEY="5388359-538583-5i9593-3596e0-6ca252484934aa4"
* NOVADAX_SECRET_KEY="nl3KVXiOp4JN74482h4nkahiu5jDKWkKhnMumMy"
*/

// novadax.New("ACCESS_KEY", "PRIVATE_KEY")
client := novadax.New("5388359-538583-5i9593-3596e0-6ca252484934aa4", "nl3KVXiOp4JN74482h4nkahiu5jDKWkKhnMumMy") // fake credentials here, just maintained a similar pattern to the actual data

// list all symbols available at NovaDAX
symbols, err := client.ListOreders()
```

### Rate Limiting ###

NovaDAX imposes a rate limit on all API clients. Public endpoints are
limited to 60 requests per second, while private endpoints can be invoked up to
20 requests per second. 

Learn more about NovaDAX rate limiting at
https://doc.novadax.com/pt-BR/#comunicacao-com-a-api.

### Integration Tests ###

TODO: implement tests

## Contributing ##
I would like to cover the entire NovaDAX API and contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward.

TODO: Contribution doc.

## License ##

This library is distributed under the GPLv3 license found in the [LICENSE](./LICENSE)
file.