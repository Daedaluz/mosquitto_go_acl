auth.so: base.c base.go main.go
	go build -buildmode=c-shared -o auth.so
