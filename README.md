<p align="center"><a href="https://keigo.konnyaku256.dev" target="_blank" rel="noopener"><img src="https://user-images.githubusercontent.com/29503528/95951305-21e87180-0e31-11eb-9a2c-467eef7ea2bb.png" width="120" /></a></p>

# keiGo
[keiGo](http://keigo.konnyaku256.dev/)は翻訳アプリのような使用感の敬語変換Webアプリケーションです。
左側の原文ペインに変換したいテキストを入力すると、右側の敬語ペインに変換結果が表示されます。

![keiGo Web App preview](https://user-images.githubusercontent.com/29503528/95950858-455eec80-0e30-11eb-9e16-932ee57b8041.png)

現在、
- 丁寧語（です/ます調）
- 尊敬語（相手を立てる表現）
- 謙譲語（自分を下げる表現）

の3種類の敬語の変換に対応しています。

## Frontend
主な採用技術
- React
- TypeScript
- styled-components
- webpack

React + TypeScriptで開発しています。
コンポーネントのスタイリングはstyled-componentsを使ったCSS in JSで書いています。
また、create-react-appなどは使わず、できるだけ最小構成のReactで開発を進めています。

[Frontendのコードとドキュメントはこちら](frontend/)

## Backend
主な採用技術
- Golang
- gin
- kagome
- Docker

Golang、ginでAPI Serverを開発しています。
アーキテクチャはMVCです。
最も重要な敬語変換ロジックはGolang向け形態素解析ツールのkagomeと独自の変換プログラムを組み合わせて実装しています。
また、上記の変換ロジックは想定するケースが膨大にあるため、testingパッケージを使ってユニットテストを導入しています。
Dockerを活用し、アプリケーションはコンテナ化して開発しています。

[Backendのコードとドキュメントはこちら](backend/)

## Infrastructure
### Netlify
Frontendのデプロイ先として使用しています。
GitHubのリポジトリと連携するだけでよいので、素早く低コストにデプロイを実現できています。

### Cloud Run
Backendアプリケーション（敬語変換サーバ）の実行環境として採用しています。
Dockerコンテナ化したアプリケーションをフルマネージドで公開、運用することができ、サーバの管理コストを大幅に下げることができています。
さらに、Cloud RunへのデプロイはGitHub Actionsを介しているため、Backendの担当者はアプリケーションの開発に集中できます。
また、デフォルトでHTTPSに対応しているため、Mixed contentの問題に対応する手間も省けています。
