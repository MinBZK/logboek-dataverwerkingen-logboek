# Logboek dataverwerkingen referentie-implementatie

Een referentie-implementatie van de [Logboek dataverwerkingen standaard](https://minbzk.github.io/logboek-dataverwerkingen/).


## Componenten

Deze repository bevat de volgende componenten:

- **[`proto`](./proto)** bevat de [Protocol Buffers](https://protobuf.dev/)-definities;
- **[`server`](./server)** de logboekserver, ontvangt verwerkingshandelingen (logs) van componenten welke data verwerken;
- **[`libs`](./libs)** softwarebibliotheken (Go en Python) om te gebruiken in dataverwerkenden componenten. Maak het eenvoudig om verwerkingshandelingen op te sturen naar de logboekserver.


### Server

De logboekserver implementeerd de [gRPC](https://grpc.io/)-service zoals beschreven in [`logboek.proto`](./proto/logboek/v1/logboek.proto).

Verwerkingshandelingen worden opgeslagen in een zogenaamde *storage backend*. De volgende worden ondersteund:

- [SQLite](https://www.sqlite.org/): handig voor tijdens het ontwikkelen, niet geschikt voor productie. Er is geen databaseserver nodig, de logboekserver kan zelfstandig de verwerkingshandelingen in een bestand opslaan;
- [Cassandra](https://cassandra.apache.org/): een gedistribueerde database, geschikt voor het opslaan van grote hoeveelheden data.

Databaseschema's voor SQLite en Cassandra zijn te vinden in [`server/pkg/storage/schema`](./server/pkg/storage/schema).


## Gebruik

De logboekserver is beschikbaar als *Docker image*: `ghcr.io/minbzk/logboek-dataverwerkingen-logboek`.

```sh
docker run --rm --publish 9000:9000 --volume ./instance:/var/lib/logboek ghcr.io/minbzk/logboek-dataverwerkingen-logboek
```

In de `./instance` directory is het databasebestand te vinden met de verwerkingshandelingen.


### Opties

De volgende opties zijn beschikbaar, als *environment variables*:

- `LISTEN_ADDRESS`: adres waar de gRPC service op beschikbaar is. Naar dit adres sturen de dataverwerkenden componenten de verwerkingshandelingen. Standaard is dit `127.0.0.1:9000`;
- `STORAGE_TYPE`: welke *storage backend* er gebuikt moet worden. Mogelijke waarden zijn: `sqlite` en `cassandra`. Standaard is dat `sqlite`;
- `STORAGE_SQLITE_PATH`: waar het SQLite-databasebestand opgslagen moet worden. Standaard is dit `logboek.db`;
- `STORAGE_CASSANDRA_SERVERS`: een lijst van Cassandra servers. Standaard is dit `127.0.0.1:9042`.


### Clients

Eenvoudig voorbeeld van hoe de Python-softwarebibliotheek gebruikt kan worden in een dataverwerkende applicatie.

```python
from logboek import get_processing_operator

logboek_operator = get_processing_operator()

with logboek_operator.start_proccessing_as_current("handeling"):
    # verwerking van gegevens
    ....
```

Kijk voor uitgebreidere voorbeelden in `README.md`'s in de afzonderlijke directory's: [`libs/logboek-go`](./libs/logboek-go/README.md) en [`libs/logboek-python`](./libs/logboek-python/README.md).


## Ontwikkeling

### Vereisten

- [Go](https://go.dev/doc/install)
- [Python 3](https://www.python.org/)
- [Docker](https://docs.docker.com/get-docker/)


### Protobuf

```sh
# Installeer de benodigdheden
make deps

# Genereer code voor de Go- en Python-softwarebibliotheken
make proto
```

### Docker image

```sh
make server-image
```


## License

Licensed under EUPL v1.2
