FROM alpine

ADD . /src

RUN cd /src

RUN go build -o goapp

CMD ./goapp
