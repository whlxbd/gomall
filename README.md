# gomall

## 项目需求
三、功能需求
认证中心
- 分发身份令牌
- 续期身份令牌（高级）
- 校验身份令牌

用户服务
- 创建用户
- 登录
- 用户登出（可选）
- 删除用户（可选）
- 更新用户（可选）
- 获取用户身份信息

商品服务
- 创建商品（可选）
- 修改商品信息（可选）
- 删除商品（可选）
- 查询商品信息（单个商品、批量商品）

购物车服务
- 创建购物车
- 清空购物车
- 获取购物车信息

订单服务
- 创建订单
- 修改订单信息（可选）
- 订单定时取消（高级）

结算
- 订单结算

支付
- 取消支付（高级）
- 定时取消支付（高级）
- 支付

AI大模型
- 订单查询
- 模拟自动下单

## 项目结构
```bash
.
├── app
│   ├── aiorder
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── auth
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── cart
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── checkout
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── frontend
│   │   ├── biz
│   │   ├── conf
│   │   ├── hertz_gen
│   │   ├── infra
│   │   ├── middleware
│   │   ├── script
│   │   ├── static
│   │   ├── template
│   │   ├── types
│   │   └── utils
│   ├── order
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── payment
│   │   ├── biz
│   │   ├── conf
│   │   └── script
│   ├── product
│   │   ├── biz
│   │   ├── conf
│   │   ├── ghz_test
│   │   └── script
│   └── user
│       ├── biz
│       ├── conf
│       └── script
├── common
│   ├── clientsuite
│   ├── mtl
│   ├── rmq
│   ├── serversuite
│   └── utils
│       └── pool
├── deploy
│   └── config
├── idl
│   └── frontend
├── rpc_gen
│   ├── kitex_gen
│   │   ├── aiorder
│   │   ├── auth
│   │   ├── cart
│   │   ├── checkout
│   │   ├── order
│   │   ├── payment
│   │   ├── product
│   │   └── user
│   └── rpc
│       ├── aiorder
│       ├── auth
│       ├── cart
│       ├── checkout
│       ├── order
│       ├── payment
│       ├── product
│       └── user
└── scripts
```