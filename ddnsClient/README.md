# 修改域名解析记录文档示例

该项目为根据传入参数修改解析记录。文档示例，该示例**无法在线调试**，如需调试可下载到本地后替换 [AK](https://usercenter.console.aliyun.com/#/manage/ak) 以及参数后进行调试。

## 运行条件

- 下载并解压需要语言的代码;


- 在阿里云帐户中获取您的 [凭证](https://usercenter.console.aliyun.com/#/manage/ak)并通过它替换下载后代码中的 ACCESS_KEY_ID 以及 ACCESS_KEY_SECRET;

- 执行对应语言的构建及运行语句

## 执行步骤

下载的代码包，在根据自己需要更改代码中的参数和 AK 以后，可以在**解压代码所在目录下**按如下的步骤执行

- Go
- *Go 环境版本必须不低于 1.10.x*
- *安装 SDK 核心库 OpenAPI*
```sh
go get github.com/alibabacloud-go/darabonba-openapi/v2/client
```
- *执行命令*
```sh
GOPROXY=https://goproxy.cn,direct go run ./main
```
## 使用的 API

-  UpdateDomainRecord 根据传入参数修改解析记录。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Alidns/2015-01-09/UpdateDomainRecord)

## 返回示例

*实际输出结构可能稍有不同，属于正常返回；下列输出值仅作为参考，以实际调用为准*


- JSON 格式 
```js
{
    "RequestId": "536E9CAD-DB30-4647-AC87-AA5CC38C5382",
    "RecordId": "9999985"
}
```
- XML 格式 
```xml
<UpdateDomainRecordResponse>
    <RequestId>536E9CAD-DB30-4647-AC87-AA5CC38C5382</RequestId>
    <RecordId>9999985</RecordId>
 </UpdateDomainRecordResponse>
```

## 更新 go.mod 中的模块路径

```
go mod edit -module newnme

```

## 打包可行文件

```sh
go build -o ddnsClient
```
## 运行
```sh
./ddnsClient --recordId 7****44 --domainName infra --ipValue *.*.*.*
```
