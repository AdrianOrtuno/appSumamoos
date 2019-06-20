FROM golang

COPY . /go/src/github.com/user/appesports-back
WORKDIR /go/src/github.com/user/appesports-back

# added vendor services will need to be included here
RUN go get ./Vendor/gorilla/mux

RUN go get ./
RUN go build
	
EXPOSE 8080