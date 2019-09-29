## Event notification services

### Description
The backend services for scanning and reminding about upcoming events

### Env variables

The env vars can be obtained from global environment, or via .env file (which should be placed in working dir).

| variable | type | default | description |
| --- | --- | --- | --- |
| LOG_LEVEL | number, 1-6 | 6 |the logging level
| DB_URI | string | postgres://user:123@localhost:5432/otus | the postgres connection string
| QUEUE_URI | string | amqp://guest:guest@localhost:5672 | the rabbitmq connection string

### Build scan service

```bash
>>> make build_scan
```

### Build notification service

```bash
>>> make build_notification
```

### Available commands

#### scan_service

run scan service and fill db with fake events ```scan_service fill <amount_of_events>```
run scan service ```scan_service scan <scan_period_in_seconds> <earliest_timestamp>```

#### notification_service

listen to events on rmq: ```service_notification```
