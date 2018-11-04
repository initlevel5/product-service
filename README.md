# product-service

The product service project

#### Usage

1. Build application
    ```
    make build
    ```

2. Build and run docker image
    ```
    ./docker-setup.sh
    ```

3. Fetch a query

    ```
    curl -X POST -H "Content-Type: application/json" \
    --data '{"query":"{product(id:\"1002\"){id title created price}}"}' \
    http://localhost:8080/query
    ```

#### Todo:
- [x] Use mock implementation
- [ ] Use PostgreSQL
- [ ] Add authentication & authorization
- [x] Add unit test cases