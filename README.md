# service example

[![reportcard](https://goreportcard.com/badge/github.com/gomatic/service-example)](https://goreportcard.com/report/github.com/gomatic/service-example)
[![build](https://travis-ci.org/gomatic/service-example.svg?branch=master)](https://travis-ci.org/gomatic/service-example)
[![godoc](https://godoc.org/github.com/gomatic/service-example?status.svg)](https://godoc.org/github.com/gomatic/service-example)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

## Install

    go install github.com/gomatic/service-example
    
## Test

**Run the service**
    
    API_PORT=5000 service-example --debug >service-example.log 2>&1 & service_pid=$!

**Test the service**

    curl -s localhost:5000/health         # The health check
    curl -s localhost:5000/api/model.json # The OpenAPI documentation

debug routes

    curl -s localhost:4999/debug/vars     # This debug route provides runtime information

### RPC

**Install RPC client helper**

    go install github.com/gomatic/service-example/cmd/service-example-client

**Call the service through the RPC using the client helper**

    API_PORT=5000 service-example-client this is a great example message

### Cleanup

    kill ${service_pid}
    rm service-example.log # Maybe have a look at the log before doing this.
