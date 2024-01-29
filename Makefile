build:
	go build -ldflags "\
	-s -w \
	-X github.com/trevinteacutter/mwo-helper/pkg/build.Build=v0.1.0 \
	-X github.com/trevinteacutter/mwo-helper/pkg/build.Commit=$(git rev-parse HEAD) \
	-X github.com/trevinteacutter/mwo-helper/pkg/build.Date=Never \
	-X github.com/trevinteacutter/mwo-helper/pkg/build.Runtime=$(go version | awk '{print $3;}') \
	" -o helper cmd/helper/main.go
