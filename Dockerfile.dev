# goのバージョン
FROM golang:1.19.3-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# golangci-lintのインストール
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.35.0
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app