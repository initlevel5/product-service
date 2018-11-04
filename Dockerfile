FROM alpine
ADD ./product-service /usr/src/app/
ENTRYPOINT ["/usr/src/app/product-service"]
EXPOSE 8080
