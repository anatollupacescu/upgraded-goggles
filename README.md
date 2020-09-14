## Exchange app

build the docker container with

```sh
docker build -t exchange .
```

to run it locally

```sh
docker run --rm -it -p8080:8080 exchange:latest
```

to provide a different upstream exchange endpoint

```sh
docker run --rm -it -p8080:8080 -e EXCHANGE_RATES_URL=http://your.new.url?USD exchange:latest
```