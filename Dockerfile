FROM python:3.14-alpine@sha256:cc95388e96eeaa0a7dbf78d51d0d567cc0e9e2ae3ead2637877858de9b41a7bf

ENV DEBIAN_FRONTEND=noninteractive
ENV PYTHONUNBUFFERED=1

RUN apk update && \
    apk upgrade

ARG WORKDIR="/version-gen"
RUN mkdir -pv ${WORKDIR}
COPY *.py ${WORKDIR}

RUN mkdir -pv /root/app

WORKDIR /root/app

RUN chmod +x "${WORKDIR}/version-gen.py" && \
    ln -s "${WORKDIR}/version-gen.py" /usr/bin/version-gen

CMD [ "sh" ]