## copy file util (like dd)

### RUN

```bash
>>> go get -u github.com/ega-forever/otus_go@task9_dd
>>> otus_go -from="testfile" -to="testfile_copy" -limit=2 -offset=23
```

### TEST
```bash
>>> go test
```