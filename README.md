# washing-ocr
洗濯表示のラベル画像を受け取り、カテゴライズして該当する文言を返すline bot

## 起動
```
make run
```

`http://localhost:8080` -> API URL

## 技術スタック
- Go 1.19

## アーキテクチャ
DDD
- main.go -> エントリーファイル, ルーティング
- entity/ -> domain層
- entity/model -> entity
- entity/repository -> repository
- application/ -> application層, application service
- infra/ -> infrastructure層, external api

## 環境変数
```
rootに.envファイルを作成して以下の環境変数を追加
```

| NAME                     | Description                  |
| ------------------------ | ---------------------------- |
| LINE_BOT_CHANNEL_SECRET  | line bot channel sercret     |
| LINE_BOT_CHANNEL_TOKEN   | line bot channel token       |

# ブランチの運用
mainからブランチを切る

PRは一人以上のapproveがあればmerge（CICDで設定したい）

## チケットの管理
github projectsで管理

## LINE Messaeging APIの運用 ※要改善
各自公式LINEを作成する。[公式LINE作成のやり方はこちら](https://qiita.com/yuki_0920/items/cbdbd5220a6a8b4eef19)
ngrokを使ってフォワーディングされたurl+/api/v1/replyをLINEの設定からwebhook urlに保存する(例:) https://<span>XXXXXXXX</span>.ngrok.io/api/v1/reply )。[webhookの設定方法](https://qiita.com/mininobu/items/b45dbc70faedf30f484e)


## 関連URL
### 議事録
https://docs.google.com/document/d/1K5yJ9He1sn9mGebhSDGGxvkcrGUwJDdW4fwPFyt8csY/edit