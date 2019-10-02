# install a go-environment
sudo apt-get install golang-go

# setup GOPATH to env
# create GOBIN and setup to ENV
mkdir -p ~/go/bin
export GOBIN=~/go/bin
export PATH=$PATH:$GOBIN


# install go sdk of docker engine.
# 
# https://docs.docker.com/develop/sdk/

go get github.com/docker/docker/client

go get github.com/containerd/containerd

go get github.com/docker/distribution
