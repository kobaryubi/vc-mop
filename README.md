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

## packages
`go get github.com/joho/godotenv`  
`go get github.com/lib/pq`  

## cmds
`go mod tidy`

## LINE
### Messaging API
- Channel type: Messaging API
- Provider: MasahikoKobayashi
- Channel name: notification
- Channel description: notification
- Category: 個人
- Subcategory: 個人（IT・コンピュータ）
- Email address: kobaryubi6@gmail.com

## PostgreSQL
`/usr/local/bin/brew install postgresql@12`  
`echo 'export PATH="/usr/local/opt/postgresql@12/bin:$PATH"' >> ~/.zshrc`  
`postgres -D /usr/local/var/postgresql@12`  
`psql -h localhost -p 5432 -U kobayashimasahiko -d postgres`  
