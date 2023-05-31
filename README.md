# Minecraft DynmapやBluemapに対応したプロキシ

MinecrafのプラグインであるDynmapやBluemapは相対パスを使用しており、

Nginxなどでリバースプロキシをする際に

パス単位にする（`example.com/map`など）のが非常にめんどくさい。

そのめんどくさい工程を自動化したhttpプロキシ

```yaml
listen: 127.0.0.1:8080 -> 起動アドレスとポート
location: 
    example: -> 名前（適当可） 
        path: /sa -> パス
        proxy_pass: http://127.0.0.1:3000 -> バックエンドのアドレス
コンフィグは自動生成されます。
```

使用方法

```
go download
go build -o httpproxy
chmod +x httpproxy
./httpproxy
```

コンフィグの生成・設置位置を変えたい場合は

```
./httpproxy -c /home/shiro/configs/httpproxy.yml
```

停止する際は、`^C`もしくはコンソールに`stop`と入力すると停止します。

専用のPterodactylEggもあるよ！ [http-proxy-pterodactyl](https://github.com/shiro8613/http-proxy-pterodactyl)