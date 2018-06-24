FROM golang:1.8-alpine

RUN apk update && \
	apk add git

WORKDIR /go/src/github.com/dougrich/goproj
COPY . /go/src/github.com/dougrich/goproj
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -o /a.out

FROM scratch

COPY --from=0 /a.out /a.out
COPY ./templates /templates

ENTRYPOINT [ "/a.out" ]
