FROM alpine:latest

#setup the golang build environments
ENV GOPATH /go
ENV GOREPO github.com/allen13/example-api-app
RUN mkdir -p $GOPATH/src/$GOREPO
ADD . $GOPATH/src/$GOREPO
WORKDIR $GOPATH/src/$GOREPO

#install, build project, and then remove golang build tools all in one step
#this minimizes image sizes since the build tools won't be saved to a run layer
RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
		git \
		go \
		build-base \
	&& go build -o example-api-app cmd/example-api-app/main.go \
	&& apk del .build-deps \
	&& rm -rf $GOPATH/pkg

EXPOSE 1323

#This is important. Don't run containers as root!
USER 1001

CMD ./example-api-app api
