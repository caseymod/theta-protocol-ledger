FROM golang:1.22rc2-bookworm
ENV GOPATH=/app
ENV PATH=$GOPATH/bin:$PATH
WORKDIR /app/src/github.com/thetatoken/theta
COPY . .
RUN make install
RUN cp -r ./integration/testnet_amber ../
EXPOSE 28888
CMD theta start --config=../testnet_amber/node --password="qwertyuiop"

