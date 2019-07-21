## Product simple CRUD backend

### Description
The backend app for running simple product CRUD app.

### Endpoints

| endpoint | method | params | output |
| --- | --- | --- | --- |
| /products | GET | | {result: [product], status: 1}
| /products | POST | {id: number, name: string} | {result: product, status: 1}
| /products | DELETE | {id: number} | {result: [product], status: 1}

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