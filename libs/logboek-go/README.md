# Logboek Go library

Softwarebibliotheek om eenvoudig dataverwerkingen te loggen.


## Installatie

```sh
go get github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go
```


## Gebruik

```go
package main

import (
    "context"
    "fmt"

    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
)

func main() {
    // Adres van de logboekserver
    logboekEndpoint := "localhost:9000"

    // Gebruik gRPC om de verwerkingshandelingen naar de logboekserver te versturen
    handler, _ := logboek.NewGRPCProcessingOperationHandler(context.Background(), logboekEndpoint)

    // Met een operator worden nieuwe verwerkingen gestart
    operator := logboek.NewProcessingOperator(handler)

    ctx := context.Background()
    // Start een nieuwe verwerking
    _, op := operator.StartProcessing(ctx, "een-verwerking")

    fmt.Println("Dataverwerking")

    // Stop de verwerking
    op.End()

    // De handler zal de verwerkingshandeling versturen naar de logboekserver
}
```

### Status

Een verwerkingshandeling heeft uitkomst welke als een status wordt opgegeven. Standaard is de status `logboek.StatusCodeUnknown`.

```go
_, op := operator.StartProcessing(ctx, "een-verwerking")

// De verwerking is successvol
op.SetStatus(logboek.StatusCodeOK)

// De verwerking is niet geslaagd
op.SetStatus(logboek.StatusCodeError)
```


### Attributen

Bij een verwerkingshandeling kunnen ook attributen opgegeven worden.

```go
import (
    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/attribute"
)

_, op := operator.StartProcessing(context.Background(), "een-verwerking")
op.SetAttributes(attribute.New("naam", "waarde"), attribute.New("dossier-nummer", "42"))
```

Om bij een verwerkingshandelingen een relatie te leggen naar een verwerkingensactiviteit uit een Register van de verwerkingsactiviteiten is een speciaal attribuut beschikbaar.

```go
import (
    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/attribute"
)

_, op := operator.StartProcessing(context.Background(), "een-verwerking")
op.SetAttributes(attribute.New(attribute.ProcessingActivityIDKey, "<id-uit-het-rva>"))
```


### Propagation

*Propagation* van de verwerkingshandeling naar een extern systeem. Dit maakt gebruik van de [traceparent](https://w3c.github.io/trace-context/#traceparent-header)-header.

```go
import (
    "context"
    "net/http"

    "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
    logboek_http "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/http"
)

func main() {
    ctx, _ := operator.StartProcessing(context.Background(), "een-verwerking")

    client := http.Client{
        // Gebruik een logboek transport om de `traceparent` header te vullen met de huidige verwerkingshandeling
        Transport: logboek_http.NewTransport(http.DefaultTransport),
    }

    req, _ := http.NewRequestWithContext(ctx, "GET", "https://example.com/", nil)
    client.Do(req)
}
```
