gocore_server
-------------

Simple REST server with 2 endpoints:

```
POST /gocore
GET /gocore
```

The POST endpoint is called from the gocore library (if the advertisingURL setting is configured) and expects to receive an arbitrary JSON payload.  This JSON should contain host, port and serviceName elements.

The JSON will be stored in memory overwriting any previous payload that has the same host, port and serviceName.

