## Product simple CRUD backend

### Description
The backend app for running simple product CRUD GRPC app.

### Endpoints

All endpoints are available as proto files in proto folder.

### Env variables

The env vars can be obtained from global environment, or via .env file (which should be placed in working dir).

| variable | type | default | description |
| --- | --- | --- | --- |
| PORT | number | 8080 |the web app port
| LOG_LEVEL | number, 1-6 | 6 |the logging level

### Build

```bash
>>> make build
```

### Available commands

#### Server

```calendar grpc_server```

#### Client

create event: ```calendar grpc_client create <text> <timestamp>```
update event: ```calendar grpc_client update <id> <text> <timestamp>```
get event: ```calendar grpc_client get <id>```
delete event: ```calendar grpc_client delete <id>```
list events: ```calendar grpc_client list```
