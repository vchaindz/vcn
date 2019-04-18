# vcn - vChain CodeNotary
# 
# Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
# This software is released under GPL3.
# The full license information can be found under:
# https://www.gnu.org/licenses/gpl-3.0.en.html

FROM golang:1.12-stretch

COPY . /vcn

WORKDIR /vcn
RUN make install

ENTRYPOINT [ "vcn" ]