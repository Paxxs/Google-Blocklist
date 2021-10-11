# Google-BlockList 搜索引擎过滤黑名单

很多垃圾站点出现在搜索引擎结果中让平常的搜索变得非常不便。

结合 uBlacklist 插件实现对 Bing/DucDuckGo/Google 等搜索结果进行过滤，提高搜索体验。

> 之前火狐浏览器 uBlacklist谷歌同步 + cobaltdisco + 此 repo 的自定规则用得还不错。
> 
> 但是近期中文环境下的谷歌，被 penzai.com 的团队凭借着一己之力，成功让谷歌前N页全是 “小X知识网” “小X百科网” 等内容农场站点，几乎变得无法使用。
> 
> 之前靠过滤域名的方式过滤他们非常吃力，因为他们使用了大量域名。经群友提醒，可以通过 uBlacklist正则匹配标题 来实现过滤他们的垃圾内容。但后来发现 cobaltdisco 最近几天也实现了相关内容。
> 
> 故更新了一下本 repo 来去重过滤内容


## 使用

配合插件 uBlacklist 实现过滤搜索引擎内容：

- [Chrome Web Store](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)
- [Firefox Add-ons](https://addons.mozilla.org/en-US/firefox/addon/ublacklist/)
- [Mac App Store](https://apps.apple.com/app/ublacklist-for-safari/id1547912640)


### 订阅规则

```
https://raw.githubusercontent.com/Paxxs/Google-Blocklist/master/uBlacklist_subscription.txt
```

## TODO

- [x] 过滤垃圾站 penzai.com （“小X知识网” “小X百科网”）
- [x] 自动合并订阅去重
- [ ] custom 规则内嵌原因

## 内容来源

- 使用 Google 搜索时收集
- [cobaltdisco/Google-Chinese-Results-Blocklist](https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist)