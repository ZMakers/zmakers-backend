# 合约相关
|Contract Name|Contract Address|Description|
|:---:|:---:|:---:|
|DigitalMediaStore|0x0245a2E012D92D9aD56A4870E20AF3aBF15a8252|contract that manges artworks|
|ApprovedCreatorRegistry|0x0033fe788b04908fb0D011E08DEe0dE4E771803f|operator account registry contract|
|ApprovedCreatorRegistryV1|0x803B37F847816C7eF1931A3F718906dA610994BA|old version of ApprovedCreatorRegistry contract|
|DigitalMediaCore|0x2F1fFd21aB5B19b8105C0Fb28D4ad8E8D9b14800|Main driver contract|

# 一、接口域名
|序号|接口环境|域名|对应ip和端口|说明|
|:---:|:---:|:---:|:---:|:---:|
|1|测试环境||||

# 二、接口列表
|序号|接口名称|请求地址|请求方式|
|:---:|:---:|:---:|:---:|
|1|注册|/user|Post|
|2|登录|/user/login|Post|
|3|创建专辑|/collections|Post|

# 三、接口
## 1、注册
### 请求类型
Post

### 请求地址
/user

### 请求参数
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|Name|用户名|String|是|
|Email|注册邮件|String|是|
|Password|密码|String|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|id|用户id|uuid格式的String|
|name|用户名|String|
|email|用户注册的email|String|
|password|密码|String|
|hash|密码的hash|String|
|token|获取访问api权限的token，有过期时限|String|
### 应答报文示例
```bigquery
{
    "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
    "name": "ZJJ01",
    "email": "zzzz@asd1.com",
    "Password": "123@123",
    "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
    "token": ""
}
```

## 2、登录
### 请求类型
Post

### 请求地址
/user/login

### 请求参数
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|Name|用户名|String|是|
|Email|注册邮件|String|是|
|Password|密码|String|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|id|用户id|uuid格式的String|
|name|用户名|String|
|email|用户注册的email|String|
|password|密码|String|
|hash|密码的hash|String|
|token|获取访问api权限的token，有过期时限|String|
### 应答报文示例
```bigquery
{
    "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
    "name": "ZJJ01",
    "email": "zzzz@asd1.com",
    "Password": "123@123",
    "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJaSkowMSIsInN1YiI6IjEyM0AxMjMiLCJleHAiOjE2Mzg2OTQ0MDgsIm5iZiI6MTYzODY4NzIwOH0.Zq1wt0LOKTBc6bgkiv1VW3EHQlqavZENrh01mJwhVtk"
}
```

## 3、创建专辑
### 请求类型
Post

### 请求地址
/colelctions

### 请求参数
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|Owner|专辑拥有者地址|String|是|
|MetaPath|专辑存储地址|String|是|
|TxHash|交易地址|String|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|owner|专辑拥有者地址|String|
|metaPath|专辑存储地址|String|
|txHash|交易地址|String|
### 应答报文示例
```bigquery
{
    "owner": "0xea674fdde714fd979de3edf0f56aa9716b898ec8",
    "metaPath": "http://io.ipfs.com ",
    "txHash": "0x1234567890123456789012345678901234567890123456789012345678901234"
}
```
