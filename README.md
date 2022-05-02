# 一、需要导入的模块（Go）
1. 阿里云
>go get github.com/aliyun/alibaba-cloud-sdk-go/services/ecs

2. 腾讯云
>go get -v -u github.com/tencentcloud/tencentcloud-sdk-go

3. 华为云
>go get -u github.com/huaweicloud/huaweicloud-sdk-go-v3
go get github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic

4. aws
>go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
github.com/aws/aws-sdk-go-v2/credentials

5. google
>go get cloud.google.com/go/compute/apiv1
go get google.golang.org/genproto/googleapis/cloud/compute/v1
go get github.com/googleapis/google-cloud-go

6.Azure
>go get github.com/Azure/azure-sdk-for-go/sdk/azcore/to
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
go get github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute

# 二、支持的功能清单
1. 虚拟机
    * 阿里云
        - [ ] 查询-单个
        - [ ] 查询-列表
        - [ ] 查询-状态
        - [ ] 创建
        - [ ] 关机
        - [ ] 启动
        - [ ] 重启
        - [ ] 删除
        - [ ] 修改实例类型
    * 腾讯云
        - [ ] 查询-单个
        - [ ] 查询-列表
        - [ ] 查询-状态
        - [ ] 创建
        - [ ] 关机
        - [ ] 启动
        - [ ] 重启
        - [ ] 删除
        - [ ] 修改实例类型
    * 华为云
        - [ ] 查询-单个
        - [ ] 查询-列表
        - [ ] 查询-状态
        - [ ] 创建
        - [ ] 关机
        - [ ] 启动
        - [ ] 重启
        - [ ] 删除
        - [ ] 修改实例类型
    * Google
        - [ ] 查询-单个
        - [ ] 查询-列表
        - [ ] 查询-状态
        - [ ] 创建
        - [ ] 关机
        - [ ] 启动
        - [ ] 重启
        - [ ] 删除
        - [ ] 修改实例类型
    * AWS
        - [x] 查询-单个 
        - [x] 查询-列表
        - [x] 查询-状态
        - [x] 创建
        - [x] 关机
        - [x] 启动
        - [x] 重启
        - [x] 删除
        - [x] 修改实例类型
    * Azrue
        - [ ] 查询-单个
        - [ ] 查询-列表
        - [ ] 查询-状态
        - [ ] 创建
        - [ ] 关机
        - [ ] 启动
        - [ ] 重启
        - [ ] 删除
        - [ ] 修改实例类型
2. 磁盘
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
    * Google
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
        - [ ] 变更配置
3. 快照
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Google
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
4. 镜像
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Google
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
5. VPC
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Google
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
6. 子网
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Google
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
7. 安全组
    * 阿里云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 腾讯云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * 华为云
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Google
        - [ ] ~~不支持安全组~~
    * AWS
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
    * Azrue
        - [ ] 创建
        - [ ] 查询
        - [ ] 删除
8. 公共资源
   * 机房
   * 可用区
   * 实例类型
