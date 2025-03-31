FROM golang:1.14.4-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
#crea la imagen para trabajar con go
#> docker build --tag goservice .

#ejecutar la imagen para levantar el contenedor go
#docker run -it --rm -p 8082:8081 --name mygoservice goservice