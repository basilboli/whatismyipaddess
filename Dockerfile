FROM golang
ADD . /go/src/github.com/basilboli/whatismyipaddress
RUN go install github.com/basilboli/whatismyipaddress
WORKDIR /go/src/github.com/basilboli/whatismyipaddress
EXPOSE 7070
ENTRYPOINT /go/bin/whatismyipaddress
