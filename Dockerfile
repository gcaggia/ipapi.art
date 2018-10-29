# STEP 1 build executable binary
FROM golang:alpine as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# STEP 2 build a small image (around 6 MB)
FROM scratch
COPY --from=builder /app/main /app/main
COPY --from=builder /app/assets/ /app/assets/
COPY --from=builder /app/templates/ /app/templates/
WORKDIR /app
EXPOSE 18000
ENTRYPOINT ["/app/main"]