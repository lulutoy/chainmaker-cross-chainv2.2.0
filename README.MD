## 项目简介
跨链简单理解就是两条或多条链之间进行数据交换。初版跨链采用网关的方案实现，该方案基于一个假设：在联盟链环境中，所有节点都是可信的，会存在异常节点，但不会存在恶意节点。为了解决可信性的问题，后续在新的版本中会采用Relay的方案。

# 版本列表：
* [V1.2.0](https://git.chainmaker.org.cn/chainmaker/chainmaker-cross/-/tags/v1.2.0)
> 主要功能：
>   直连，代理模式下的 chainmaker 同构跨链
* [V1.2.4](https://git.chainmaker.org.cn/chainmaker/chainmaker-cross/-/tags/v1.2.0)
> Change Log：
>   修复了跨链SDK在并发场景下的使用问题
* [V2.1.0](https://git.chainmaker.org.cn/chainmaker/cross-chain/-/tags/v2.1.0)
> Change Log：
>   优化了跨链流程，添加了跨链网关http路由方式，优化了跨链SDK的使用方式

# 适配版本一览表

| chainmaker-cross | chainmaker-go    | chainmaker-spv   | chainmaker-sdk-go |
| :----------------| :----------------| :----------------| :-----------------|
| v1.2.0           | v1.2.0           | v1.2.0           | v1.2.0            |
| v1.2.4           | v1.2.4           | v1.2.3           | v1.2.4            |
| v2.1.0           | v2.0.0+          | v2.1.0           | v2.0.0            |

# 快速了解
访问以下链接来快速了解项目结和使用方法
* [设计方案](./docs/跨链设计方案.MD)
* [使用文档](./docs/跨链使用指南.MD)

# License
chainmaker-cross 项目在 Apache License, Version 2.0 (Apache-2.0) 下构建, 参考 [LICENSE](./LICENSE) 文件说明.