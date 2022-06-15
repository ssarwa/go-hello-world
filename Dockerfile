FROM golang:1.18-buster AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-hello-world
# Create a new release build stage
FROM gcr.io/distroless/base-debian10
# Set the working directory to the root directory path
WORKDIR /
# Copy over the binary built from the previous stage
COPY --from=builder /go-hello-world /go-hello-world
EXPOSE 8080
ENTRYPOINT ["/go-hello-world"]