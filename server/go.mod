module github.com/MinBZK/logboek-dataverwerkingen-logboek/server

go 1.22

require (
	github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go v0.0.0-00010101000000-000000000000
	github.com/gocql/gocql v1.6.0
	github.com/mattn/go-sqlite3 v1.14.22
	google.golang.org/grpc v1.63.2
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240521202816-d264139d666e // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go => ../libs/logboek-go
