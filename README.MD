# Google-BlockList 搜索引擎过滤黑名单

很多垃圾站点出现在搜索引擎结果中让平常的搜索体验变得非常不便。

本项目结合 uBlacklist 插件实现对 Bing/DucDuckGo/Google 等搜索结果进行过滤，提高相关搜索体验。

前情提要：

> 之前火狐浏览器 uBlacklist 谷歌同步 + cobaltdisco + 此 repo 的自定规则用得还不错。
> 
> 但是近期中文环境下的谷歌，被 penzai.com 的团队凭借着一己之力，成功让谷歌前N页全是 “小X知识网” “小X百科网” 等内容农场垃圾站点，几乎让谷歌中文检索瘫痪变得无法使用。
> 
> 之前靠过滤域名的方式来过滤他们会非常吃力，因为他们使用了大量域名。经群友提醒，可以通过 uBlacklist正则匹配标题 来实现过滤他们的垃圾内容。但后来发现 cobaltdisco 最近几天也实现了相关内容。
> 
> 故创建了本 repo 来自动生成 uBlacklist 订阅并去重内容


## 使用

配合插件 uBlacklist 实现过滤搜索引擎内容。

根据你的浏览器来选择：

- [Chrome Web Store](https://chrome.google.com/webstore/detail/ublacklist/pncfbmialoiaghdehhbnbhkkgmjanfhe)
- [Firefox Add-ons](https://addons.mozilla.org/en-US/firefox/addon/ublacklist/)
- [Mac App Store](https://apps.apple.com/app/ublacklist-for-safari/id1547912640)

### 订阅规则

订阅会全自动更新，不需要你 clone 到本地生成，在 uBlacklist 的设置选项中找到「订阅」并添加我们的项目链接地址即可。

Github : `https://git.io/gblock`

```
https://raw.githubusercontent.com/Paxxs/Google-Blocklist/develop/uBlacklist_subscription.txt
```

CDN加速（适合中国内陆用户）：

```shell
https://cdn.jsdelivr.net/gh/Paxxs/Google-Blocklist@develop/uBlacklist_subscription.txt
```

## TODO

- [x] 过滤垃圾站 penzai.com （“小X知识网” “小X百科网”）
- [x] 自动合并订阅去重
- [ ] custom 规则内嵌原因
- [ ] 单独域名IP列表

## Reference

- 使用 Google 搜索时收集
- [cobaltdisco/Google-Chinese-Results-Blocklist](https://github.com/cobaltdisco/Google-Chinese-Results-Blocklist)
- [dallaslu/penzai-list: list of g.penzai.com](https://github.com/dallaslu/penzai-list)
- [liubiantao/uBlacklist-Websites: 谷歌搜索结果中屏蔽机器翻译网站](https://github.com/liubiantao/uBlacklist-Websites)
- [arosh/ublacklist-stackoverflow-translation: Stack Overflow の机械翻訳サイトの除外用フィルタ](https://github.com/arosh/ublacklist-stackoverflow-translation)
