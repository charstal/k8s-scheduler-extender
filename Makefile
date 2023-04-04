MAJOR=1
MINOR=15
PATCH=12
VERSION=v$(MAJOR).$(MINOR).$(PATCH)


COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z) GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))
BUILDENVVAR=CGO_ENABLED=0


image:
	docker build -t scheduler-plugin-agent:$(VERSION) .

remote: image
	docker tag scheduler-plugin-agent:$(VERSION) registry.cn-hangzhou.aliyuncs.com/charstal/scheduler-plugin-agent:$(VERSION)
	docker push registry.cn-hangzhou.aliyuncs.com/charstal/scheduler-plugin-agent:$(VERSION)