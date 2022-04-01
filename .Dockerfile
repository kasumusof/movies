
FROM golang:1.18
ENV GOLANG_VERSION 1.18

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN go install -v github.com/gobuffalo/pop/soda
RUN make createprod
RUN go build -o main .


CMD ["/app/main"]