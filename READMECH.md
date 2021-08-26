# 合约相关
|合约名|合约地址|描述|
|:---:|:---:|:---:|
|DigitalMediaStore|0x0245a2E012D92D9aD56A4870E20AF3aBF15a8252|艺术品存储合约|
|ApprovedCreatorRegistry|0x0033fe788b04908fb0D011E08DEe0dE4E771803f|操作账户注册合约|
|ApprovedCreatorRegistryV1|0x803B37F847816C7eF1931A3F718906dA610994BA|老版本操作账户注册合约|
|DigitalMediaCore|0x2F1fFd21aB5B19b8105C0Fb28D4ad8E8D9b14800|主合约|

# 接口
* 创建 Collection

请求头 :

|请求头|请求内容|说明|
|:---:|:---:|:---:|
|Authorization|Basic secretKey|访问token|
|Content-Type|application/json|请求方式|
``请求方式``：POST

```请求头```：

|key|value|Description|
|:---:|:---:|:---:|
|creator|String||
|MetaPath|String|存储collection的ipfs目录|
|TxHash|String|交易号|

``请求示例``：
```
localhost:3000/collections

```

``响应示例``：
```

```

|参数|类型|备注|
|:---:|:---:|:---:|
|Owner|String||
|MetaPath|String|存储collection的ipfs目录|
|TxHash|String|交易号|


