+++
date = "2019-10-07"
title = "Build caddy server hugo+github"
slug= "Build-caddy-server-hugo+github"
tags = ["linux"]
+++

# Build caddy server hugo+github

caddy跟hugo兩個都是使用golang這個語言寫的，效能妥妥的

caddy還有很多優點除了效能好，還有自動簽證書和configure檔案方便好寫，並且有很多外掛功能，我們這次也會用到其中一個騷操作。

## 1.我們需要golang來編譯hugo的source code

這裡在linux上，我推薦先安裝brew再透過brew來安裝hugo，之後更新會方便很多，brew也會幫你把需要的套件一併安裝

[Homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux)

[Install Hugo](https://gohugo.io/getting-started/installing/)

p.s如果中途發生ram不夠的情況，可以看我上一篇文章

安裝完成後可以下

    hugo version

來確認是否安裝成功

## 2.再來安裝caddy，可以按照底下網址操作

[caddy官方脚本一键安装与使用](https://medium.com/@jestem/caddy%E5%AE%98%E6%96%B9%E8%84%9A%E6%9C%AC%E4%B8%80%E9%94%AE%E5%AE%89%E8%A3%85%E4%B8%8E%E4%BD%BF%E7%94%A8-1e6d25154804)

裡面比較需要注意的是在配置systemd時

    sudo curl -s  https://raw.githubusercontent.com/mholt/caddy/master/dist/init/linux-systemd/caddy.service  -o /etc/systemd/system/caddy.service
    sudo systemctl daemon-reload
    sudo systemctl enable caddy.service
    sudo systemctl status caddy.service

可以看個人需求，我個人在ProtectHome這邊從true改成false，方便之後的github webhook，因為只有跑hugo所以不太需要擔心安全性的問題，但如果有跑其他的服務這裡可以斟酌一下

    ;Hide /home, /root, and /run/user. Nobody will steal your SSH-keys.
    ProtectHome=false

## 3.接下來我們開始編寫caddyfile和處理github webhook

首先當然要再github創新一個新的repo，然後在caddyfile中填入必要資訊，比如說repo的網址，要選擇哪個branch，要放到哪裡，接受到這個請求後要執行那個指令，你的webhooks key是什麼，都填好之後我們將畫面回到github上。

![](../Untitled-0bd7ed8d-6ad1-490f-b14e-6d5f6ce5f7bf.png)

然後進入設定裡面有Webhooks這個選項，點選Add webhook

![](../Untitled-51487f9a-1be7-4641-b6ca-df54c06aeca7.png)

填入你要接受webhooks的網址和Secret，並將content type選成application/json和將ssl verification打開((反正caddy自帶https，然後點選Just the push event。

## 4.完成

之後你就可以將你邊寫好的md放入clone下來的hugo資料夾裡面然後再push上github，你的caddy server就會接受到通知他會clone最新的repo下來並自動執行hugo指令，一個自動更新的hugo server就完成摟!