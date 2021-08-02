FROM golang:1.16.4
RUN mkdir /gm-tool-backend
RUN mkdir /gm-tool-backend/storage
RUN mkdir /gm-tool-backend/storage/en
RUN mkdir /gm-tool-backend/storage/id

ADD . /gm-tool-backend

WORKDIR /gm-tool-backend

RUN go build -o main


CMD ["/gm-tool-backend/main"]