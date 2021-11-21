build-linux:
	cd cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rss

build-image:
	docker build -t docker12345rekcod/rss_backend:latest .

push-image: build-image
	docker push docker12345rekcod/rss_backend:latest