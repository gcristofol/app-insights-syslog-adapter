# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9

COPY ./mysyslog/*.go /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog/
WORKDIR /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog

# Fetch or manage dependencies here
RUN go get github.com/spf13/viper
RUN go get github.com/Microsoft/ApplicationInsights-Go/appinsights
RUN go get gopkg.in/mcuadros/go-syslog.v2

# Build the mysyslog command inside the container.
RUN ls -lisa 
RUN go build
RUN pwd
RUN ls -lisa


# Run the outyet command by default when the container starts.
#COPY ./mysyslog /go/bin/
RUN chmod a+x /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog/mysyslog
#ENTRYPOINT /go/src/github.com/gcristofol/app-insights-syslog-adapter/mysyslog/mysyslog

# Document that the service listens on port 8080.
EXPOSE 5000
