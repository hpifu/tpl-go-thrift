FROM centos:centos7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY docker/tpl-go-thrift /var/docker/tpl-go-thrift
RUN mkdir -p /var/docker/tpl-go-thrift/log

EXPOSE 7060

WORKDIR /var/docker/tpl-go-thrift
CMD [ "bin/echo", "-c", "configs/echo.json" ]
