module github.com/KeatonBrink/distributed-chessboard-generation/cb_master

go 1.19

replace github.com/KeatonBrink/distrubuted-chessboard-generation/distributedchessboardgeneration => ../distributedchessboardgeneration/

require (
	github.com/KeatonBrink/distrubuted-chessboard-generation/distributedchessboardgeneration v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
