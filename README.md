# bee myblog reflactor

# 步骤

1. bee api bee-myblog
2. git init 
3. git remote add origin https://username:password@github.com/xxxxx  ( git remote remove origin )
4. git push origin master:dev(master)
5. git pull origin dev(master) master
6. bee run 

7. go get -u dep
8. dep init
```python
# 依赖管理帮助
dep help ensure
# 添加一条依赖
dep ensure -add github.com/bitly/go-simplejson
# 这里 @= 参数指定的是 某个 tag
dep ensure -add github.com/bitly/go-simplejson@=0.4.3
# 添加后一定记住执行 确保 同步
dep ensure
# 同理建议使用
dep ensure -v
# 更新依赖
dep ensure -update -v
#  删除没有用到的 package
dep prune -v
```

go-simplejson
go-jwt


引用需要两个集合，内嵌只用一个集合
# mongo 数据库设计
博文对评论： 一对多
博文对归档： 多对一
博文对标签： 多对少
实体：
1. blogs 博文
    - _ id 
    - 创建的时间 createdAt   new Date()
    - 更新的时间 updatedAt   new Date()     
    - 简述 summary           string
    - 标题 title             string 
    - 评论数 commentCount    int 
    - 评论   comments        数组 引用comments
    - 阅读数 readCount       int 
    - 标签 tags              内嵌tags 
        - 名字 name

2. comments 评论
    - _ id
    - 用户名 username          string 
    - 内容 content             string
    - 时间 createdAt           new Date()
3. archives 归档
    - _ id
    - 名字 name                string
    - 包含的博文 blogs         数组 引用blogs
    
4. collections 收藏文档资源
    - _ id 
    - title 标题      string
    - author 作者     string
    - url 连接        string
    - time  收藏时间  new Date()
    
# redis 设计
- [ ] 缓存readCount
- [ ] 标签存在 缓存中
- [x] 缓存 jwt 字段 的refresh token


问：有人说本地缓存密码和用户名，自动登录不就行了么？

答：客户端缓存用户名和密码，容易受到黑客攻击，如果使用了token机制，就算黑客盗取了用户的access_token与refresh_token，只需重新登录，就可令原来的token失效。
> 解释：
> 1. 在登录的时候：返回给前端一个 access_token , 一个 refresh_token ,并且在redis 中写入 userId_access_token : access_token 过期时间一天,userId_refresh_token:refresh_token 过期时间一个月
> 2. 每次请求的时候都判断一下，redisConn.get(userId_access_token)，如果有值，返回true 如果为空，直接401
> 3. redisConn.get(userId_refresh_token)，如果有值，那么刷新access_token,refresh_token，如果也没值，return 401
> 前端逻辑:
> 1. 将access_token, refresh_token 存在 localstorage里
> 2. 每次访问接口带上 access_token
> 3. if  401  那么就要 refresh_token
> 4. 那么就带上 refresh_token json请求到 服务器 /v1/auth/refresh_token
> 5. 如果回应是 access_token  那么刷新localstorage继续访问
> 6. 如果是 401，那么 跳到登录界面

**这样既实现了自动登录，又保证了安全**

# TODO：
- [x] Dockerfile
- [x] docker-compose.yml
- [x] .gitignore

====
- [x] swagger api 文档生成
- [x] logs 日志记录
- [x] tests  单元 集成测试

====
- [ ] script ci cd 
- [ ] .travis.yml

====
- [x] LICENSE

# API设计
### 1. blogs  /v1/blogs?xxx
1. get （时间变成年月日，评论内嵌进去，找出blog的归档信息）
    - 不带query_params 返回所有 blogs
    - 带参数
        + offset=起始值
        + limit=返回的最大值
        + sort=更新的时间（后期可以加上什么按 阅读数，评论数排序）
        + tag=标签 
       
2. post 
    - title
    - summary
    - tags
    - archive
    
### 2. blog   /v1/blog/_ id
1. get（时间变成年月日，评论内嵌进去，找出blog的归档信息）

2. post(x)

3. put（发送其中任意之一）
    - title
    - summary
    - tags
    - archive

4. delete（删除）

### 3. archives  /v1/archives
1. get （把blog信息内嵌进去）

2. post （添加归档）
    - name
    
### 4. archive   /v1/archive/_ id
1. get（得到一篇归档）

2. post(x)

3. put(x)

4. delete(删除)

### 5. collections  /v1/collections
1. get（得到所有的收藏信息）
    - no query params
    - 带参数
        + sort=
        + offset= 
        + limit=
        
2. post(更新)
    - title
    - author
    - url
    - time
    
### 6. collection  /v1/collection/_ id
(x)











