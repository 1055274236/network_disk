# 网盘

## 前端

[前端](https://github.com/1055274236/network-disk-front-end)

~~正在编写中....~~

暂停开发

## 主要功能

用户登录/注册，同账号异地登录登出，多线程下载，断点续传，秒传，日志等。

## 功能优化

### 储存空间优化

将文件 md5、sha1 作为文件唯一标识符，创建文件的过程只是在数据库中建立一个到某静态文件的索引，已达到节省存储空间的目的，同时还能够作为秒传的标识。

## 功能实现

### 上传

上传主要有用户端先创建文件夹，其次创建文件，然后根据文件 ID 进行上传。
