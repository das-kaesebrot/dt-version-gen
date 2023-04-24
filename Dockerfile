FROM python:3-alpine

ARG WORKDIR="/version-gen"
RUN mkdir -pv ${WORKDIR}
COPY *.py ${WORKDIR}

RUN chmod +x "${WORKDIR}/version-gen.py" && \
    ln -s "${WORKDIR}/version-gen.py" /usr/bin/version-gen

ENTRYPOINT [ "/usr/bin/version-gen" ]