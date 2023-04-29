FROM python:3-alpine

ARG WORKDIR="/version-gen"
RUN mkdir -pv ${WORKDIR}
COPY *.py ${WORKDIR}

RUN apk update && \
    apk upgrade && \
    apk add bash shadow

RUN usermod --shell /bin/bash root

RUN chmod +x "${WORKDIR}/version-gen.py" && \
    ln -s "${WORKDIR}/version-gen.py" /usr/bin/version-gen

ENTRYPOINT [ "/usr/bin/version-gen" ]