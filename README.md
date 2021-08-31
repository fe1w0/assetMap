# assetMap
基于扫描器技术探测⽹络内的数字资产，当前版本中(v0.0.1)实现功能如下：

- [x] 利用sync和的channel,实现并发协程
- [x] 利用nmap公开的指纹 [nmap-service-probes](https://svn.nmap.org/nmap/nmap-service-probes) ,来识别端口服务,

> 但依旧不全面,如正则匹配效果不佳,无banner输出的服务无法判断,如135,需要进一步改进和对nmap 的原理进一步学习

此外,针对绕过防火墙的绕过,如http请求走私，随机化地址扫描,以及速度和识别率上的进一步提高，如采用多次扫描,修改`ulimit`。以及，对最后数据的处理上，对网站进行指纹识别和分析服务之间的关系。

## 使用方法

```terminal                                                                                                                                                                 ─╯
❯ ./assetMap --help

NAME:
   assetMap - Simple assets scanner

USAGE:
   assetMap [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

AUTHOR:
   fe1w0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ip value, -i value           ip list
   --port value, -p value         port list (default: "22,23,53,80-139")
   --mode value, -m value         scan mode
   --timeout value, -t value      timeout (default: 4)
   --concurrency value, -c value  concurrency (default: 10)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)

```
## 设计思路





## 应用测试

![image-20210831035456698](http://img.xzaslxr.xyz/20210831085641.png)




## 参考与推荐阅读

[将nmap指纹集成到扫描器中.md - 小草窝博客 (hacking8.com)](https://x.hacking8.com/post-418.html)

[TideFinger/Web指纹识别技术研究与优化实现.md at master · TideSec/TideFinger (github.com)](https://github.com/TideSec/TideFinger/blob/master/Web指纹识别技术研究与优化实现.md)

[projectdiscover之naabu 端口扫描器源码学习.md - 小草窝博客 (hacking8.com)](https://x.hacking8.com/post-406.html)

[从 Masscan, Zmap 源码分析到开发实践 (seebug.org)](https://paper.seebug.org/1052/)

[Nmap原理02 - 编写自己的服务探测脚本 - 随风浪子的博客 - 博客园 (cnblogs.com)](https://www.cnblogs.com/liun1994/p/6986544.html)

[nmap-service-probes File Format | Nmap Network Scanning](https://nmap.org/book/vscan-fileformat.html)

[白帽子安全开发实战-赵海锋-微信读书 (qq.com)](https://weread.qq.com/web/reader/fd932b4072398309fd92017ke4d32d5015e4da3b7fbb1fa)

其中《白帽子安全开发实践》和 [将nmap指纹集成到扫描器中.md - 小草窝博客 (hacking8.com)](https://x.hacking8.com/post-418.html),两者提供了主要思路和帮助。
