
# syntax=docker/dockerfile:1
# Copyright (C) 2022, Berachain Foundation. All rights reserved.
# See the file LICENSE for licensing terms.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#######################################################
###           Stage 0 - Build Arguments             ###
#######################################################

ARG GO_VERSION
ARG GOARCH=amd64
ARG GOOS=linux
ARG APP_NAME=genbld
ARG CMD_PATH=./genbld

#######################################################
###         Stage 1 - Build the Application         ###
#######################################################

FROM golang:${GO_VERSION}-alpine as builder

# Set the working directory
WORKDIR /workdir

# Copy go.mod and go.sum files
COPY ./genbld/go.sum ./genbld/go.mod ./genbld/

# Download the go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN env GOOS=${GOOS} GOARCH=${GOARCH} && \
    go build \
    -mod=readonly \
    -trimpath \
    -o /workdir/bin \
    ${CMD_PATH}

COPY /workdir/bin/${APP_NAME} /bin/

#######################################################
###        Stage 2 - Prepare the Final Image        ###
#######################################################

FROM polard/base:v0.0.0 as basepolard

RUN apk add --no-cache bash

# Build args
ARG APP_NAME

COPY /workdir/bin/genbld /bin/
COPY /workdir/bin/${APP_NAME} /bin/

CMD ["sh", "-c", "sleep infinity"]
