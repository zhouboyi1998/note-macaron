FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-macaron
COPY ./main /go/note-macaron
COPY ./application-docker.yaml /go/note-macaron

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18083
ENTRYPOINT ["./main"]
