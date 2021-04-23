<div align="center" style="vertical-align: center;">
  <img src="https://cdn.rpaka.dev/icon/scene-picks.png" height="80px" style="margin-right: 15px;" />
  <h1>ScenePicks</h1>
  <h1>楽天株式会社 夏の陣2020</h1>
  <h2>C日程チームF</h2>
</div>

<br/>

<div align="center">
  <img src="https://cdn.rpaka.dev/useimage/scene-picks/demo.gif" />
</div>

## 概要

漫画、アニメ、本など自分の好きな作品のセリフを共有することを通して作品を紹介し、人々を繋げる WEB サービスのプロトタイプです。

## 背景

このプロトタイプは楽天株式会社サマーインターン「夏の陣 2020」でチームメンバーと共に制作しました。

## 利用

### ディレクトリ構成

```
.
├──app // フロントサイドコード
├──docs // データ定義など
├──server // サーバサイドコード
│    ├──gen // swaggerが吐き出すファイルの保存先
│    ├──handler
│    └──main.go
├──docker-compose.yml
├──Dockerfile
└──README.md
```

### サーバの起動

環境変数を設定

```
export DBENV=localhost
```

ローカルでサーバを起動する

```
# 起動
make local/run
```

docker を利用してサーバを起動

```
# 起動
make docker/run

# 停止
make docker/stop
```

## 構成

![システム構成](https://cdn.rpaka.dev/arch/scene-picks.png)
