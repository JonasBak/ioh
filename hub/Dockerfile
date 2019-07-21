FROM golang as build

WORKDIR /build

ADD . /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=build /build/main .

CMD ["./main"]
