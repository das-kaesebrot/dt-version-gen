FROM python:3.14-alpine@sha256:2a77c2640cc80f5506babd027c883abc55f04d44173fd52eeacea9d3b978e811

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