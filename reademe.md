# goPetStore: use gin and gorm
使用 go语言 实现的 jpetstore
通过重构原无框架版，使用 gin 和 gorm 进行框架开发
原无框架版：https://github.com/SwordHarry/gopetstore

### 业务模块
- 商品模块
    - category
    - product
    - item
    - search
- 购物车模块
    - cart
- 用户模块
    - account
        - TODO: update bug
- 订单模块
    - order
    - lineItem
    - sequence

### 架构
template + gin + gorm + mysql
