IMAGE = docker12345rekcod/rss_frontend:latest

dist: 
	npm run build

build-image: dist
	docker build -t $(IMAGE) .

push-image: build-image
	docker push $(IMAGE)