IMAGE = docker12345rekcod/rss_frontend:latest

dist: 
	npm run build

build-image: dist
	docker build --platform linux/amd64 -t $(IMAGE) .

push-image: build-image
	docker push $(IMAGE)