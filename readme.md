# goPetStore: use gin and sqlx
使用 go语言 实现的 jpetstore<br/>
通过重构原无框架版，使用 gin 和 sqlx 进行框架开发<br/>
原无框架版：https://github.com/SwordHarry/gopetstore<br/>

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
template + gin + sqlx + mysql<br/>
原采用 gorm，但其具有一定入门门槛和学习成本，其关联查询和关联插入搞不清楚。。。<br/>
并且比对起来像 java 的 herbinate，故移用轻小的 sqlx<br/>
