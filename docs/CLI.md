# CLI

This document describes how to run the API.

## Installing

```shell
go install github.com/shoriwe/rollee-test-2@latest
```

## Flags

| Flag         | Description                                                  | Example                                                      |
| ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `--database` | Use this command to specify the SQL backend to use. Available are: `sqlite://` and `postgres://`. If ignored runs the entire Database in memory | `postgres://host=127.0.0.1 user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable` |

## Example

```shell
api --database sqlite://... IP:PORT
```

