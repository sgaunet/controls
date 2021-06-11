
FROM golang:1.16.4-alpine AS builder
LABEL stage=builder

RUN apk add --no-cache git upx

WORKDIR /go/src/
COPY src /go/src
COPY src/go.mod /go/src

RUN cd /go/src && \
    go get && \
    CGO_ENABLED=0 GOOS=linux go build . 

RUN ls -al /go/src/
RUN upx /go/src/controls
##################################################################



FROM alpine:3.12.1

RUN addgroup -S ctluser_group -g 1000 && adduser -S ctluser -G ctluser_group --uid 1000
RUN apk add --update --no-cache bash \
    && rm -rf /tmp/* /var/cache/apk/*

RUN wget https://github.com/sgaunet/mdtohtml/releases/download/0.2/mdtohtml_0.2_Linux_x86_64.tar.gz \
    && tar zxvf mdtohtml_0.2_Linux_x86_64.tar.gz \
    && mv mdtohtml /usr/bin/mdtohtml \
    && rm mdtohtml_0.2_Linux_x86_64.tar.gz

COPY --from=builder /go/src/controls            /usr/bin/controls
RUN chmod +x /usr/bin/controls

USER ctluser
CMD [ "/usr/bin/controls" ]