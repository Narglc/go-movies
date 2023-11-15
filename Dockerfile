#FROM scratch
FROM busybox:latest

WORKDIR /app
COPY . /app

#EXPOSE 8000, 默认直接使用编译好的端口
CMD ["./go_movies"]
#RUN echo "just a test"
