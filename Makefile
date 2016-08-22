TAG = <1.0.0>
PREFIX = <PREFIX>

.PHONY: all
all: deploy

.PHONY: build
build:
	go build -o whatismyipaddress -tags netgo -ldflags "-extldflags '-lm -lstdc++ -static'" .
	docker build -t $(PREFIX):$(TAG) .

.PHONY: build
push:
	docker push $(PREFIX):$(TAG)

.PHONY: run
run: build
	docker run --rm -p 7070:7070 $(PREFIX):$(TAG)

.PHONY: deploy
deploy: push
	kubectl create -f spec.yaml

.PHONY: deployOnly
deployOnly:
	kubectl create -f spec.yaml


.PHONY: delete
delete:
	kubectl delete --ignore-not-found -f spec.yaml