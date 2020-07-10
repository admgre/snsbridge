FROM golang:1.14.2-alpine as build-env
RUN apk add --no-cache git make
ADD . /go/src/github.com/admgre/snsbridge
WORKDIR /go/src/github.com/admgre/snsbridge
RUN go get -d -v ./server/...
RUN make && mv snsbridge /
RUN adduser -h /runhome -s /bin/sh -D run

FROM scratch
COPY --from=build-env /etc/ssl/certs /etc/ssl/certs
COPY --from=build-env /etc/passwd /etc/passwd
COPY --from=build-env /runhome /runhome
COPY --from=build-env /snsbridge /snsbridge
USER run
ENTRYPOINT [ "/snsbridge" ]