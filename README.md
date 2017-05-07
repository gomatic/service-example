# example

    go get github.com/gomatic/service-example
    service-example & server_pid=$!
    service-example-client
    kill ${server_pid}
