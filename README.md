# Project Doctrina

Project Doctrina is a playground project for myself to learn a new tech stack.

At its infancy, Doctrina is a rock-paper-scissors application that is written in golang and deployed as a container.

## Building Docker Images

Note that all Docker images are Linux based. In order to build a golang executable for Linux engines, remember to set GOOS, GOARCH, and CGO_ENABLED to properly cross-compile.

Example when building rsp-service: 

```
$ export GOOS=linux
$ export GOARCH=amd64
$ export CGO_ENABLED=0
$ go build -o rsp-service-linux-amd64
```

After, make sure you return to original values for your local system, (e.g. set GOOS to ```darwin``` for Mac).
