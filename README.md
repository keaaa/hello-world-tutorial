# Hello world

simple server demonstrating getting started with radix

## Docker

build `docker build -t hello-world .`

run `docker run -it --rm -p 8000:8000 hello-world` 

## k6

run `k6 run --vus 100 --duration 60s loadtest.js`