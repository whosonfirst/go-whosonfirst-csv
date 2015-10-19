prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-csv; then rm -rf src/github.com/whosonfirst/go-whosonfirst-csv; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-csv
	cp *.go src/github.com/whosonfirst/go-whosonfirst-csv/

deps:   self

fmt:
	go fmt *.go

bin: 	self

