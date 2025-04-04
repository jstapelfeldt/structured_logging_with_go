# Domain specific structured logging example with zap logger

## Setup

Services are instantiated in the main file.

Every domain service has its own zap logger as a field which is created in the constructor with its domain name.

The logging library ensures that all domain specific logs have a domain specific tag in the structured log, e.g. `"domain":"sales_service"`.

## Dynamic log level adjustment
The logging library offers an endpoint that takes as arguments the **domain** name with a desired **log level**. Using this endpoint one can change the log level per domain at runtime. 


## Example 


Starting the service
```bash
➜  go run main.go
```

```json
{"level":"info","timestamp":"2025-04-04T19:50:57.002+0200","msg":"application starting","domain":"global"}
{"level":"info","timestamp":"2025-04-04T19:50:57.002+0200","msg":"server running on :8080","domain":"global"}
```

Creating an order using the sales service
```bash
➜  curl "http://localhost:8080/process-order?orderID=44"
```

```json
{"level":"info","timestamp":"2025-04-04T18:36:03.513+0200","msg":"processing order","domain":"sales_service","order_id":"44"}
```


Building a house using the infra service
```bash
➜  curl "http://localhost:8080/build-infra?type=house"   
```
```json
{"level":"info","timestamp":"2025-04-04T18:36:29.780+0200","msg":"building object","domain":"infra_service","object_type":"house"}
```

Setting the log level to DEBUG for the sales service domain
```bash
➜  curl "http://localhost:8080/set-log-level?domain=sales_service&level=debug"
```
```json
{"level":"info","timestamp":"2025-04-04T22:37:42.058+0200","msg":"log level updated","domain":"sales_service","new_level":"debug"}
```

Creating an order (now logs on DEBUG)
```json
{"level":"debug","timestamp":"2025-04-04T18:37:26.005+0200","msg":"processing order started","domain":"sales_service"}
{"level":"info","timestamp":"2025-04-04T18:37:26.005+0200","msg":"processing order","domain":"sales_service","order_id":"13"}
{"level":"debug","timestamp":"2025-04-04T18:37:26.005+0200","msg":"processing order finished","domain":"sales_service"}
```

Building more infrastructure still logs on level INFO
```json
{"level":"info","timestamp":"2025-04-04T18:37:41.659+0200","msg":"building object","domain":"infra_service","object_type":"house"}
```

## Library

- https://github.com/uber-go/zap
- https://pkg.go.dev/go.uber.org/zap


# Happy logging!
