FROM progrium/busybox
MAINTAINER Brad Pinter <brad.pinter@gmail.com

ADD ./stage/littlegopher /bin/littlegopher

ENV DOCKER_HOST unix:///tmp/docker.sock

ENTRYPOINT ["/bin/littlegopher"]
