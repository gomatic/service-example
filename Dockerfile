FROM scratch

WORKDIR /

ENV HOME=/
ENV PWD=/
ENV PATH=/
ENV TMP=/
ENV TEMP=/

ADD service-linux-amd64 /

EXPOSE 2999 3000 3001

ENTRYPOINT ["/service-linux-amd64"]
