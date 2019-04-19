FROM golang AS multistage

WORKDIR /echo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

##

FROM alpine:3.9
WORKDIR /app
COPY --from=multistage /echo/echo /app/echo
EXPOSE 3000
ENTRYPOINT [ "./echo" ]
