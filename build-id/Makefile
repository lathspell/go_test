all:
	go build -i -ldflags "-X main.build=$$(git rev-parse --short HEAD)" build_id.go 

list_deps:
	go list -f '{{join .Deps "\n"}}' | xargs go list -f '{{.ImportPath}}'

list_deps_non_std:
	go list -f '{{join .Deps "\n"}}' | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'	

