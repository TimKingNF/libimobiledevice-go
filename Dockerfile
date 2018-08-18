FROM golang:1.10

# Install libimobiledevice
# Let's build whatever we have, and will clean up later
RUN apt-get update
RUN apt-get install -y git build-essential automake libtool pkg-config
RUN apt-get install -y python-dev libplist-dev libplist++-dev libusb-1.0-0-dev libfuse-dev libssl-dev

ENV PKG_CONFIG_PATH="/usr/local/lib/pkgconfig"
ENV LD_LIBRARY_PATH="/usr/local/lib"
ENV LDFLAGS='-L/usr/local/lib/'
RUN cd /root && git clone https://github.com/libimobiledevice/libusbmuxd.git && cd libusbmuxd && ./autogen.sh --prefix=/usr/local && make && make install
RUN cd /root && git clone https://github.com/libimobiledevice/libimobiledevice.git && cd libimobiledevice && ./autogen.sh --enable-debug-code --prefix=/usr/local && make && make install
RUN cd /root && git clone https://github.com/libimobiledevice/usbmuxd.git && cd usbmuxd && ./autogen.sh --prefix=/usr/local && make && make install
RUN cd /root && git clone https://github.com/libimobiledevice/ifuse.git && cd ifuse && ./autogen.sh --prefix=/usr/local && make && make install
RUN cd /root && rm -rf *

# RUN apt-get remove -y python-dev libplist-dev libplist++-dev libusb-1.0-0-dev libfuse-dev libssl-dev git build-essential automake libtool pkg-config
# RUN apt-get clean

# Golang starts here
WORKDIR $GOPATH/src/github.com/alyyousuf7/libimobiledevice-go
ADD . .
