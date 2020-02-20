#!/bin/bash

function usage() {
cat <<_EOT_
Usage:
  $0 contest last_task

Description:
  コンテスト用のフォルダを作成するスクリプト

Options:
  --help, -h  print this.
_EOT_
exit 1
}

case ${1} in
    help|--help|-h)
        usage
    ;;
esac

readonly CONTEST=$1"/"
readonly LAST_TASK=$2
readonly BASE_DIR="./contest/"

# 開始位置判定
start_task=1
if [[ ${LAST_TASK} =~ ^[A-Z]$ ]]; then
  start_task="A"
elif [[ ${LAST_TASK} =~ ^[a-z]$ ]]; then
  start_task="a"
elif [[ ${LAST_TASK} =~ [1-9]?[0-9] ]]; then
  start_task=1
else
  echo "最終問題番号の指定が誤りです。lastには「1-99」、「A-Z」または「a-z」を指定してください。"
  exit 1
fi

if [ -e $BASE_DIR$CONTEST ]; then
  echo "すでに{$CONTEST}ディレクトリが存在します"
  exit 1
else
  # コンテストのディレクトリを作成
  mkdir $BASE_DIR$CONTEST

  # 問題のディレクトリを作成
  eval mkdir $BASE_DIR$CONTEST{$start_task..$LAST_TASK}
  eval echo $BASE_DIR$CONTEST{$start_task..$LAST_TASK} | xargs -n 1 cp -rv ./src/template/*
  exit 0
fi
