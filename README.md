# Project Doctrina

Project Doctrina is a playground project for myself to learn a new tech stack.

At its infancy, Doctrina is a rock-paper-scissors application that is written in golang and deployed as a container.

## Building Docker Images

Note that all Docker images are Linux based. In order to build a golang executable for Linux engines, remember to set GOOS.

Example when building rsp-service: 

```
$ export GOOS=linux
$ go build -o rsp-service-linux-amd64
```

Then set GOOS back to your operating system (e.g. ```darwin``` for Mac).
