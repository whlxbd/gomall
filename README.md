# gomall

## 代码生成命令
```bash
kitex -module whlxbd.github.com/gomall idl/*.proto
```

```bash
# 进入 auth 服务目录并生成代码
cd rpc/auth
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.auth -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/auth.proto
cd ../..

# 进入 cart 服务目录并生成代码
cd rpc/cart
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.cart -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/cart.proto
cd ../..

# 进入 checkout 服务目录并生成代码
cd rpc/checkout
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.checkout -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/checkout.proto
cd ../..

# 进入 order 服务目录并生成代码
cd rpc/order
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.order -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/order.proto
cd ../..

# 进入 payment 服务目录并生成代码
cd rpc/payment
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.payment -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/payment.proto
cd ../..

# 进入 product 服务目录并生成代码
cd rpc/product
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.product -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/product.proto
cd ../..

# 进入 user 服务目录并生成代码
cd rpc/user
kitex -module whlxbd.github.com/gomall -service whlxbd.github.com.gomall.user -use whlxbd.github.com/gomall/kitex_gen -I ../../idl ../../idl/user.proto
cd ../..
```