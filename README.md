# novadax-go #

novadax-go is a Go client library for accessing the [NovaDAX API v1](https://doc.novadax.com/).

## Usage ##

```go
import "github.com/artemis-tech/novadax-go/client"
```

Construct a new GitHub client, then use the various services on the client to
access different parts of the GitHub API. For example:

```go
import (
    novadax "github.com/artemis-tech/novadax-go/client"
)

client := novadax.NewClient()

// list all symbols available at NOVADAX
symbols, err := client.ListSymbols()
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