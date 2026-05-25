FROM golang:1.26-trixie@sha256:0f6b034c99663ea8957e7dae99124e37374cbe7fcb5b5646f19b185f8f976279

RUN GO111MODULE=on go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.19.0
RUN GO111MODULE=on go install k8s.io/code-generator/cmd/conversion-gen@v0.29.3

RUN mkdir /gatekeeper
WORKDIR /gatekeeper
