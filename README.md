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

# ブランチの運用
mainからブランチを切る

PRは一人以上のapproveがあればmerge（CICDで設定したい）

## チケットの管理
github projectsで管理

## 関連URL
### 議事録
https://docs.google.com/document/d/1K5yJ9He1sn9mGebhSDGGxvkcrGUwJDdW4fwPFyt8csY/edit