FROM golang:alpine AS dev
WORKDIR /rota
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o ./bin/rota

FROM alpine:latest AS prod
WORKDIR /rota
COPY --from=dev /rota/bin .
EXPOSE 1926 1927
ENTRYPOINT ["./rota"]