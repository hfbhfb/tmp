## 临时用项目

通过控制台创建ELB Ingress

https://support.huaweicloud.com/usermanual-cce/cce_10_0251.html


另外推荐一个华为云的Ingress材料，里面有讲基础原理（数据流，配置流）和Ingress支持的Service类型。
路由概述：https://support.huaweicloud.com/usermanual-cce/cce_10_0094.html

## end

1.支持给部署的应用分配固定IP，按照给的方法测试失败
2.服务网格测试目前的不是springcloud，所以大部分功能测试不了，需要给我们个demo
3.安全性，就是网络朔源的问题，在容器场场景，如何通过流量/抓包来判断是哪个业务使用的。
4.https证书该怎么在服务端配置
5.日志易对接@华为 于欣彤 这些问题麻烦配合下哈
