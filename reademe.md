# goPetStore: use gin and sqlx
使用 go语言 实现的 jpetstore
通过重构原无框架版，使用 gin 和 sqlx 进行框架开发
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
- 订单模块
    - order
    - lineItem
    - sequence

### 架构
template + gin + sqlx + mysql
原采用 gorm，但其具有一定入门门槛和学习成本，其关联查询和关联插入搞不清楚。。。
并且比对起来像 java 的 herbinate，故移用轻小的 sqlx
