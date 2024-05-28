# Logboek Python library


## Installatie

```sh
pip install "logboek @ https://github.com/MinBZK/logboek-dataverwerkingen-logboek/archive/main.zip#subdirectory=libs/logboek-python"
```


## Gebruik

```python
from logboek import init_processing_operator
from logboek.handlers import GrpcProcessingOperationHandler

# Adres van de logboekserver
logboek_endpoint = "localhost:9000"

# Gebruik gRPC om de verwerkingshandelingen naar de logboekserver te versturen
handler = GrpcProcessingOperationHandler(logboek_endpoint)

# Configureer eenmalig de operator
init_processing_operator(handler)

# Met een operator worden nieuwe verwerkingen gestart
operator = get_processing_operator()

# Start een nieuwe verwerking, als een *context manager*
with operator.start_proccessing_as_current("een-verwerking") as op:
    pass
```

Verwerkingenhandelingen kunnen ook gelaagd zijn. In onderstaand voorbeeld wordt automatisch de binnenste handeling onderdeel van de buitenste.

```python
with operator.start_proccessing_as_current("buiten"):
    with operator.start_proccessing_as_current("binnen"):
        pass
    pass
```

Het `parent_operation_id` van de binnenste verwerkingenhandeling verwijst naar het `operation_id` van de buitenste verwerkingenhandeling.


### Status

Een verwerkingshandeling heeft uitkomst welke als een status wordt opgegeven. Standaard is de status `StatusCode.UNKNOWN`.

```python
from logboek import StatusCode

with operator.start_proccessing_as_current("een-verwerking") as op:
    # De verwerking is successvol
    op.set_status(StatusCode.OK)

    # De verwerking is niet geslaagd
    op.set_status(StatusCode.ERROR)

# Bij een *exception* wordt de status automatisch op StatusCode.ERROR gezet.
with operator.start_proccessing_as_current("een-verwerking"):
    raise Exception()
```


### Attributen

Bij een verwerkingshandeling kunnen ook attributen opgegeven worden.

```python
with operator.start_proccessing_as_current("een-verwerking") as op:
    op.set_attributes({"naam": "waarde", "dossier-nummer": "42"})
```

Om bij een verwerkingshandelingen een relatie te leggen naar een verwerkingensactiviteit uit een Register van de verwerkingsactiviteiten is een speciaal attribuut beschikbaar.

```python
from logboek.attributes import Core

with operator.start_proccessing_as_current("een-verwerking") as op:
    op.set_attribute(Core.PROCESSING_ACTIVITY_ID, "<id-uit-het-rva>")
```


### Propagation

*Propagation* van de verwerkingshandeling naar een extern systeem. Dit maakt gebruik van de [traceparent](https://w3c.github.io/trace-context/#traceparent-header)-header.

```python
from logboek.propagators import TraceContextPropegator
import requests

propegator = TraceContextPropegator()

with operator.start_proccessing_as_current("een-verwerking") as op:
    headers={}
    propegator.inject(headers)

    requests.get("https://example.com", headers=headers)
```
