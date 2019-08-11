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

### RUN

```bash
>>> make run
```