scar: server/*.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o snsbridge ./server/

clean:
	rm -f scar
