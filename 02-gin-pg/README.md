# 
- install <a href="https://go.dev/dl/"> go </a>
- init go
```go
go mod init example/gin
```
- install gin
```go
go get -u github.com/gin-gonic/gin
```
- install and setup air
    - install air
    ```
    go install github.com/cosmtrek/air@latest
    ```

    - recheck go bin in path
    ```bash
    echo $PATH | grep $(go env GOPATH)/bin
    ```
    - if not have `/Users/sothanav/go/bin` in `PATH` add PATH in `bash.rc` or `.zshrc`
    ```bash
    ...
    PATH=$PATH:$(go env GOPATH)/bin
    ...
    ```
    - init air
    ```
    air init
    ```
- run server
```
air -c .air.toml
```