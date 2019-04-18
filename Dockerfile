FROM golang AS multistage

WORKDIR /upstream
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

##

FROM alpine:3.9
WORKDIR /app
COPY --from=multistage /upstream/upstream /app/upstream
EXPOSE 3000
ENTRYPOINT [ "./upstream" ]
