# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9

COPY ./mysyslog /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog
WORKDIR /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/spf13/viper
RUN go get github.com/Microsoft/ApplicationInsights-Go/appinsights
RUN go get gopkg.in/mcuadros/go-syslog.v2

RUN go build

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/mysyslog

# Document that the service listens on port 8080.
EXPOSE 5000
