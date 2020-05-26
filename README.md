# web-store go project

## RESTful API implementation

1. POST /product Create product
1. GET /product/:id Show product by ID
1. GET /product/all Show product list
1. PUT /product/:id Update product information (admin user only)

## DOC

see [doc](https://v-bus.github.io/web-store/)

## Run

```bash
./web-store
```

## build

```bash
ln -s $WORKSPACE/web-store $GOROOT/web-store
cd web-store
make static
```

## migration tool

You can update your current DB by [migration tool](web-store-db-migrate)

To build migration tool

```bash
cd web-store-db-migrate
make server
```

To run migration tool

```bash
cd web-store
./web-store-db-migrate-bin
```
