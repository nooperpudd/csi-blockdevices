module github.com/nooperpudd/csi-blockdevices

go 1.20

require (
	github.com/akutz/gofsutil v0.1.2
	github.com/container-storage-interface/spec v0.1.0
	github.com/rexray/gocsi v1.2.2
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.51.0

)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.7
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
//github.com/container-storage-interface/spec => github.com/container-storage-interface/spec v0.2.0
)

require (
	github.com/akutz/gosync v0.1.0 // indirect
	github.com/coreos/etcd v3.3.13+incompatible // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20230209195159-6f3db454fdf8 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.3 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/thecodeteam/gofsutil v0.1.2 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
