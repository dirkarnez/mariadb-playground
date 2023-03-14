FROM golang:buster

# RUN adduser --disabled-password --gecos '' docker
# RUN adduser docker sudo
# RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# USER docker

COPY . /go/src/github.com/dirkarnez/mariadb-playground/
WORKDIR /go/src/github.com/dirkarnez/mariadb-playground/

RUN apt-get update; \
	apt-get install -y --no-install-recommends \
		bash \
		sudo \
	&& \
	rm -rf /var/lib/apt/lists/* && \
	go build -o mariadb-playground

ENTRYPOINT ["/bin/bash"]

# CMD ./mariadb-playground --database=default --docker=true	

EXPOSE 5000