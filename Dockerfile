FROM golang:1.5

RUN apt-get -q update && \
      apt-get install --no-install-recommends -y --force-yes -q \
      ca-certificates \
      zsh \
      curl \
      git \
      cmake \
      && \
      apt-get clean && \
      rm -rf /var/lib/apt/lists/*

RUN go get golang.org/x/tools/cmd/present
COPY . /presentation
WORKDIR /presentation

RUN present
