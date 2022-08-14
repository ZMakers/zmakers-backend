# 合约相关
|Contract Name|Contract Address|Description|
|:---:|:---:|:---:|
|DigitalMediaStore|0xe4f2527977ad08f684f35e623af28db48cd8ded4|contract that manges artworks|
|ApprovedCreatorRegistry|0x76f890ea86af16213201cd7d8180626d14fbe84a|operator account registry contract|
|DigitalMediaStoreV1|0x62bf912dc67c84cf82e2e6b66e444a00dbd4fe69|old version of ApprovedCreatorRegistry contract|
|DigitalMediaCore|0xc80ebd7726654db5d3ee899884696e81629cafb3|Main driver contract|

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
    "code": 200,
    "data": {
        "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
        "name": "ZJJ01",
        "email": "zzzz@asd1.com",
        "Password": "123@123",
        "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
        "token": ""
    }
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
    "code": 200,
    "data": {
        "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
        "name": "ZJJ01",
        "email": "zzzz@asd1.com",
        "Password": "123@123",
        "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJaSkowMSIsInN1YiI6IjEyM0AxMjMiLCJleHAiOjE2Mzg2OTQ0MDgsIm5iZiI6MTYzODY4NzIwOH0.Zq1wt0LOKTBc6bgkiv1VW3EHQlqavZENrh01mJwhVtk"
    }
}
```

## 3、创建专辑
### 请求类型
Post

### 请求地址
/colelctions

### 请求头
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|token|访问api的权限token，登录后可获得|String|是|

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
    "code": 200,
    "data": {
        "owner": "0xea674fdde714fd979de3edf0f56aa9716b898ec8",
        "metaPath": "http://io.ipfs.com ",
        "txHash": "0x1234567890123456789012345678901234567890123456789012345678901234"
    }
}
```

## 4、创建专辑
### 请求类型
Post

### 请求地址
/colelctions# 合约相关
|Contract Name|Contract Address|Description|
|:---:|:---:|:---:|
|DigitalMediaStore|0xe4f2527977ad08f684f35e623af28db48cd8ded4|contract that manges artworks|
|ApprovedCreatorRegistry|0x76f890ea86af16213201cd7d8180626d14fbe84a|operator account registry contract|
|DigitalMediaStoreV1|0x62bf912dc67c84cf82e2e6b66e444a00dbd4fe69|old version of ApprovedCreatorRegistry contract|
|DigitalMediaCore|0xc80ebd7726654db5d3ee899884696e81629cafb3|Main driver contract|

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
|4|创建艺术品|/media|Post|
|5||||
|6||||

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
    "code": 200,
    "data": {
        "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
        "name": "ZJJ01",
        "email": "zzzz@asd1.com",
        "Password": "123@123",
        "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
        "token": ""
    }
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
    "code": 200,
    "data": {
        "id": "2f94b4fc-46bc-4d42-bbb2-aa5a7d0e5fe5",
        "name": "ZJJ01",
        "email": "zzzz@asd1.com",
        "Password": "123@123",
        "hash": "347db5f31963daaeb91a0743ebac7b865ce56a8db425a7d2d1554e6143f2b808",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJaSkowMSIsInN1YiI6IjEyM0AxMjMiLCJleHAiOjE2Mzg2OTQ0MDgsIm5iZiI6MTYzODY4NzIwOH0.Zq1wt0LOKTBc6bgkiv1VW3EHQlqavZENrh01mJwhVtk"    
    }
}
```

## 3、创建专辑
### 请求类型
Post

### 请求地址
/colelctions

### 请求头
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|token|访问api的权限token，登录后可获得|String|是|

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
    "code": 200,
    "data": {
        "owner": "0xea674fdde714fd979de3edf0f56aa9716b898ec8",
        "metaPath": "http://io.ipfs.com ",
        "txHash": "0x1234567890123456789012345678901234567890123456789012345678901234"
    } 
}
```

## 4、创建艺术品
### 请求类型
Post
### 请求地址
/media

### 请求头
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|token|访问api的权限token，登录后可获得|String|是|

### 请求参数
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|Id|艺术品编号|uint64|是|
|TotalSupply|艺术品流通量|uint32|是|
|CollectionId|艺术品所属专辑|uint64|是|
|MetadataPath|存储地址|String|是|
|NumRelease|发行量（<=流通量）|uint32|是|
|TxHash|交易id|String|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|id|艺术品编号|uint64|
|totalSupply|艺术品流通量|uint32|
|collectionId|艺术品所属专辑|uint64|
|metadataPath|存储地址|String|
|numRelease|发行量（<=流通量）|uint32|
|txHash|交易id|String|
### 应答报文示例
```bigquery
{
    "code": 200,
    "data" {
        "id": 121,
        "totalSupply": 10,
        "collectionId": 1,
        "metadataPath": "sdasda",
        "numRelease": 5,
        "txHash": "0xwqeq"
    }
}
```

## 5、创建艺术品发行

## 6、购买艺术品
### 请求类型
Post
### 请求地址
/media/buy

### 请求头
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|token|访问api的权限token，登录后可获得|String|是|

### 请求参数
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|TxHash|交易编号|String|是|
|From|销售者|String|是|
|To|购买者|String|是|
|MediaId|艺术品编号|int|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|txHash|交易编号|String|是|
|from|销售者|String|是|
|to|购买者|String|是|
|mediaId|艺术品编号|int|是|

### 应答报文示例
```bigquery
{
    "code": 200,
    "data": {
        "txHash": "0x0ab91284ce4ab2ef0f8017fc12c236bdc48548a7b26017897a97784bb7909fd8",
        "from": "0x7b3c17dda766f0c8c9fff8b00a27b307410ade6f",
        "to": "0x811bfaf0a74bd9d27a57d726d39db6b7a76e15cd",
        "mediaId": 23
    }
}
```

## 7、生成nft
### 请求类型
Post
### 请求地址
/nft?nftOwner="用户地址"

### 请求头
|参数|描述|类型|是否必填|
|:---:|:---:|:---:|:---:|
|nftOwner|token拥有者|String|是|

### 应答报文
|参数|描述|类型|
|:---:|:---:|:---:|
|txHash|交易编号|String|是|

### 应答报文示例
```bigquery

{
    "code": 200,
    "data": {
        "txHash": "0x3e2b2e1b1905afa8bd4331375b6ed4ac3a88b656629ab9acf9a4498728692228"
    }
}
```
