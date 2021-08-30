# vc-mop


## gore
インストール

`go get github.com/motemen/gore/cmd/gore`

.zshrcに下記を追加。

`export GOPATH=$(go env GOPATH)`

`export PATH=$PATH:$GOPATH/bin`


## golang.org/x/tools/cmd
`go get golang.org/x/tools/cmd/...`  
`go get golang.org/x/lint/golint`  

## cmds
`go mod tidy`

`go vet ./...`  
`golint -set_exit_status ./...`