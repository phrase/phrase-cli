FROM ubuntu:22.04

RUN dpkg --add-architecture i386
RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates wine-stable wine32 xvfb

ENV WINEDEBUG fixme-all
ENV WINEARCH win32
ENV FORCE_COLOR 1

# # Install Inno Setup binaries
RUN curl -SL "http://www.jrsoftware.org/download.php/is.exe" -o is.exe
RUN xvfb-run wine is.exe /VERYSILENT /SUPPRESSMSGBOXES /NORESTART
