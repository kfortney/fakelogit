FROM golang:alpine as build
LABEL maintainer="kwfortney@gmail.com"
ENV GO_SRC $GOPATH/src
ENV FAKELOG_GITHUB github.com/kfortney/fakelogit
ENV FAKELOG_ROOT $GO_SRC/$FAKELOG_GITHUB

RUN mkdir -p $FAKELOG_ROOT
WORKDIR $FAKELOG_ROOT
COPY . .
RUN go get -u $FAKELOG_GITHUB...
ENTRYPOINT ["fakelogit"]