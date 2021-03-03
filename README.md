# Assignment 1

This is a pretty minimal solution to the assignment, nothing too special going on. Errors are reported over http or by logging them to stdout.

The server is very picky about how to format the url path (all lower case, dates must be yyyy-mm-dd). Error codes and messages will tell you if you did something wrong.

## Build, run and test

Build the project with:
`go build`

Run the assignment after building with:
`./assignment-1`

Run tests with:
`go test`

## Endpoints by example

### Diag

```
GET /exchange/v1/diag
```

```json
{
    "excheratesapi": 200,
    "restcountries": 200,
    "version": "v1",
    "uptime": 12
}
```

### History

```
GET /exchange/v1/exchangehistory/norway/2020-01-01-2020-01-10
```

```json
{
    "base": "EUR",
    "start_at": "2020-01-01",
    "end_at": "2020-01-10",
    "rates": {
        "2020-01-02": {
            "NOK": 9.8408
        },
        "2020-01-03": {
            "NOK": 9.8315
        },
        "2020-01-06": {
            "NOK": 9.8488
        },
        "2020-01-07": {
            "NOK": 9.8548
        },
        "2020-01-08": {
            "NOK": 9.8508
        },
        "2020-01-09": {
            "NOK": 9.8665
        },
        "2020-01-10": {
            "NOK": 9.8745
        }
    }
}
```

### Border

```
GET /exchange/v1/exchangeborder/norway?limit=2
```

```json
{
    "base": "NOK",
    "rates": {
        "Finland": {
            "currency": "EUR",
            "rate": 0.097508654
        },
        "Sweden": {
            "currency": "SEK",
            "rate": 0.98873776
        }
    }
}
```

## Third party

I chose to use a routing library called chi that simplifies parameter matching and using middlewares.

1. [chi](https://github.com/go-chi/chi)
