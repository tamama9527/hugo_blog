+++
date = "2021-03-29"
title = "Notion to Hugo"
slug = "notion-to-hugo"
tags = ["golang","linux"]
+++
# Notion to Hugo



Notion是我很喜歡使用的程式，可以很視覺的安排工作內容，寫寫文章。

![](../9546ec4cc5ec495d8a5d9ccb4346003f-notion-labs-inc-logo-vector.png)
其中最重要一點是Notion支援markdown，這大大的方便我們這些寫程式的人們可以撰寫排版漂亮的文章了。

![](../9546ec4cc5ec495d8a5d9ccb4346003f-hugo-logo-wide.svg)
再來Hugo是一個很棒的靜態網頁編譯器，一樣也是用go寫的

基於上面幾個因素，我想把notion當作寫BLOG的平台，但這會有幾個困點點需要克服。

1. notion並沒有提供webhook的功能，所以當文章完成後並沒有辦法自動抓取下來
2. notion匯出的markdown跟一般的markdown有一點出入，並且為了hugo theme的模板，我們需要對匯出的檔案進行parser
3. 因為notion匯出的是一整包zip檔，他會將文章中用到的圖片放在一個資料夾底下，為了因應hugo，我們需要將媒體檔案跟markdown放在同一個目錄下。
4. hugo會將title當作網站的網址，遇到中文字時他會轉成url encode。

為了解決上述問題，走了很多看起來不太優美的路。

第一個問題的解決方法就是Notion支援推播到Slack，所以我寫了一個接受slack event webhook的程式。

接下來的兩個問題都是靠"正則"萬解，要特別提的是notion在匯出圖片時都會將圖片取名成Untitle{?}.png，如果我們不對檔名做特別的處理，這樣第二篇文章就會把圖片蓋掉了，我這裡用的是將Notion對每個文章都有唯一的pageID，將他放在檔名前面"pageID-Untitle.png"，當然做了這樣的修改，也必須在markdown裡面將它的相對路徑改好，一樣也是用正則處理(這世界上沒有一個正則處理不了的字串，如果有那就兩個)。

第四個問題，我是用slug解決，他會將中文轉成羅馬拼音。

source code : 

[](https://github.com/tamama9527/hugo_blog)