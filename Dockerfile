FROM    golang:1.7.3
COPY    . /go/src/github.com/moul/chatsvc
WORKDIR /go/src/github.com/moul/chatsvc
CMD     ["chatsvc"]
EXPOSE  8000 9000
RUN     make install
