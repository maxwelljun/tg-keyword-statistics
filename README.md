## 给羊王开发的关键词查询记录统计机器人
## 功能
- 按照小时\天\周\月\年\总量 统计关键词查询的次数
- 自助查询相关统计结果
- 每个小时\每天 自动推送相关结果

## 编译
### 安装依赖
- `go get -u github.com/mattn/go-sqlite3`
- `go get -u github.com/go-telegram-bot-api/telegram-bot-api`
- `go get -u github.com/robfig/cron`

### 编译
在代码文件路径下运行 `go build ./`


## 使用方法
- 初次使用 `./statistics tg-bot-token` , 会将token存到数据库中
- 后面使用 `./statistics` , 无需输入token
- 统计方法: 按照"找手电筒"这种格式的关键词会被统计在内,关键词是"手电筒"
- 查询方法: 拥有相关查询权限的人可以使用命令 `/hour` `/day` `/week` `/month` `/year` `/sum` 查询相关结果. 默认返回前20条,可以添加返回数量的参数,例如`/day 101`

## 说明
因为是定制软件,没有考虑针对多种使用条件的情况,特此说明
- 统计群组是写死的某些群组
- 查询权限是写死的几个人,其他人查不了
