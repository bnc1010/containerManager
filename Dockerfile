FROM golang:latest


RUN mkdir /app

WORKDIR /app

RUN go env -w GOPROXY=https://goproxy.cn \
    && go env -w GO111MODULE=on \
    && go install github.com/cloudwego/hertz/cmd/hz@latest 


EXPOSE 8888

###########--For deploy--############
#COPY entrypoint.sh /entrypoint.sh
#ENTRYPOINT [ "/entrypoint.sh" ]
#####################################