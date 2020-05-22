# web-store go project

## RESTful API implementation

1. POST /product Create product
1. GET /product/:id Show product by ID
1. GET /product/all Show product list
1. PUT /product/:id Update product information (admin user only)

## Possible errors

1. 403 Unauthorized
1. 200 "No such Product ID"
1. 403 "No user specified"
1. 403 "Only admin can update product"

## DOC

see [doc](https://v-bus.github.io/web-store/docs/)

## Run

```bash
./web-store
```

## build

```bash
ln -s $WORKSPACE/web-store $GOROOT/web-store
cd dev-06/go/web-store
go build ./...
```
