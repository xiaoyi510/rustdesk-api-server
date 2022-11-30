FROM alpine:3.15
RUN mkdir /app

# 复制配置文件
COPY conf /app/conf
RUN mv /app/conf/config_default.yml /app/conf/config.yml

# 复制主文件
COPY rustdesk-api-server /app
WORKDIR /app
ENTRYPOINT ["./rustdesk-api-server"]

# 导出端口号
EXPOSE [21114]