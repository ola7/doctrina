FROM iron/base

EXPOSE 8989

ADD rsp-service-linux-amd64 /
ADD healthchecker-linux-amd64 /

HEALTHCHECK --interval=3s --timeout=3s CMD ["./healthchecker-linux-amd64", "-port=8989"] || exit 1
ENTRYPOINT ["./rsp-service-linux-amd64"]