FROM debian

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -q && \
    apt-get -y install curl locales && \
    locale-gen "en_US.UTF-8" && \
    export LANG=en_US.UTF-8

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -q && \
    apt-get -y install texlive-xetex \
                       texlive-fonts-recommended \
                       texlive-fonts-extra \
                       texlive-lang-english

ENV UGID='90' UGNAME='latexer'

RUN addgroup --gid "${UGID}" "${UGNAME}"
RUN useradd --no-create-home -g "${UGNAME}" "${UGNAME}"

COPY myapp /latexapi

RUN mkdir /latexapi_tex && \
    chown latexer:90 /latexapi_tex &&\
    chown latexer:90 /latexapi

EXPOSE 8080

ENV GIN_MODE=release

WORKDIR /latexapi_tex

COPY templates/ /latexapi_tex/templates

# Add addional resources, eg images etc
# COPY logo.pdf /latexapi_tex/

USER latexer

CMD ["/latexapi"]
