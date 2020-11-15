#
# build stage
#

FROM golang:latest as builder

RUN mkdir -p /build
WORKDIR /build
COPY . .

RUN GOSUMDB=off go mod download
RUN ./contrib/scripts/build_bundle.sh

#
# run stage
#

FROM busybox:glibc

MAINTAINER smirnoffl@mail.ru
WORKDIR /dist
COPY --from=builder /build/dist/ /dist/

ENTRYPOINT ["./inmemory"]

