FROM golang:1.27-bookworm
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# COPY *.go ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080
CMD ["/docker-gs-ping"]

