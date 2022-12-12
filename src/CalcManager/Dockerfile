FROM golang:1.19.2

WORKDIR /usr/src/app

ENV LOKI_URL=http://loki:3100/api/prom/push
ENV LOKI_FORMAT=proto
ENV LOKI_SORCE=backend
ENV LOKI_JOB=calc

COPY go.mod go.* ./
RUN go mod download && go mod verify
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .
RUN go build -gcflags="all=-N -l" -v -o /usr/local/bin/app ./... 

EXPOSE 2345 
EXPOSE 5001

CMD ["dlv","debug","--listen=:2345","--headless=true","--api-version=2","--accept-multiclient"]

