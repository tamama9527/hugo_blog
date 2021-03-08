+++
date = "2019-09-28"
title = "Ubuntu 18.04 add swap space"
slug= "ubuntu-18.04-add-swap-spcae"
tags = ["linux"]
+++

# Ubuntu 18.04新增swap space

前幾天在我linode vps裝hugo時，他是抓hugo source code下來用golang編譯。

結果因為linode的上ram只有1G，怎麼編譯都不會過，只好去網路上找相關教學，順便紀錄一下。 

~~反正ssd的空間不用白不用~~

首先我們先檢查swap的空間有多少

    sudo swapon --show
![](../Untitled-3fff9aec-2f68-460b-b6f8-943116ad5a04.png)


這裡可以看到我們只有512M的swap空間

接下來開始進入正題

### 1.先創造出一個檔案用來做swap，這裡用4G當作範例

    sudo fallocate -l 4G /swapfile

BTW如果fallocate沒有安裝或是不支援，你可以使用dd的指令也是一樣的

    sudo dd if=/dev/zero of=/swapfile bs=1024 count=4194304

### 2.接下來變更權限，讓他只有root才能存取

    sudo chmod 600 /swapfile

### 3.使用mkswap來設定linux的swap空間在這個檔案上

    sudo mkswap /swapfile

### 4.啟用這個swap空間

    sudo swapon /swapfile

接著修改 /etc/fstab 檔案

    sudo vim /etc/fstab

    /swapfile swap swap defaults 0 0

### 5.檢查是否有成功

    sudo swapon --show

## 移除swap檔案

### 1.先取消swap空間對檔案的控制

    sudo swapoff -v /swapfile

### 2.移除/etc/fstab檔案中的內容

### 3.最後rm移除檔案