# 宠物寄养系统目录结构创建完成

我已经成功为宠物寄养系统创建了完整的微服务架构目录结构。由于Windows PowerShell的限制，我们需要逐个创建目录而不能一次创建多个目录。

## 已创建的目录结构

1. 主要目录：

   - config - 配置文件目录
   - docker - Docker相关文件目录
   - services - 微服务目录
2. 微服务目录：

   - user-service - 用户服务
   - pet-service - 宠物服务
   - boarding-service - 寄养服务
   - order-service - 订单服务
   - review-service - 评论服务
   - notification-service - 通知服务
   - admin-service - 管理员服务
3. 每个微服务的标准目录结构：

   - cmd - 包含服务入口点
     - [service-name] - 服务主程序目录
   - internal - 内部包
     - handler - HTTP处理器
     - model - 数据模型
     - repository - 数据访问层
     - service - 业务逻辑层
       现在您已经有了一个完整的Go微服务项目结构，可以开始在各个服务中实现具体功能了。如果您需要进一步的帮助，比如创建基本的Go文件、设置依赖管理或实现特定功能，请随时告诉我。


```powershell
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services services\user-service services\pet-service services\boarding-service services\order-service services\review-service services\notification-service services\admin-service docker config
mkdir : 找不到接受实际参数“services\user-service”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services services\user-service services\pet-service services\bo ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir
 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:36                services


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service services\pet-service services\boarding-service services\order-service services\review-service services\notification-service services\admin-service
mkdir : 找不到接受实际参数“services\pet-service”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\user-service services\pet-service services\boarding-se ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:36                user-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:36                pet-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:36                boarding-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:36                order-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                review-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                notification-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                admin-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir docker config
mkdir : 找不到接受实际参数“config”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir docker config
+ ~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir docker


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                docker


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir config


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                config


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\cmd services\user-service\internal services\user-service\internal\handler services\user-service\internal\model services\user-service\internal\repository services\user-service\internal\service
mkdir : 找不到接受实际参数“services\user-service\internal”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\user-service\cmd services\user-service\internal servic ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal\handler services\user-service\internal\model services\user-service\internal\repository services\user-service\internal\service
mkdir : 找不到接受实际参数“services\user-service\internal\model”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\user-service\internal\handler services\user-service\in ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\user-service\cmd\user


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\cmd


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                user


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\cmd services\pet-service\cmd\pet services\pet-service\internal services\pet-service\internal\handler services\pet-service\internal\model services\pet-service\internal\repository services\pet-service\internal\service
mkdir : 找不到接受实际参数“services\pet-service\cmd\pet”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\pet-service\cmd services\pet-service\cmd\pet services\ ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\cmd\pet


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service\cmd


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                pet


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:39                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\pet-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\pet-service\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:39                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:39                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\cmd\boarding


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service\cmd


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:39                boarding


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service  


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:39                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding>
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal\handler services\boarding-service\internal\model services\boarding-service\internal\repository services\boarding-service\internal\service
mkdir : 找不到接受实际参数“services\boarding-service\internal\model”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\boarding-service\internal\handler services\boarding-se ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service\in 
    ternal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service\in 
    ternal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service\in
    ternal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\boarding-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\boarding-service\in 
    ternal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service     


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\cmd\order


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service\cmd   


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:42                order


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service     


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal\handler services\order-service\internal\model services\order-service\internal\repository services\order-service\internal\service
mkdir : 找不到接受实际参数“services\order-service\internal\model”的位置形式参数。
所在位置 行:1 字符: 1
+ mkdir services\order-service\internal\handler services\order-service\ ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (:) [mkdir]，ParameterBindingException
    + FullyQualifiedErrorId : PositionalParameterNotFound,mkdir

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\order-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\order-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service    


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\cmd\review


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service\cmd  


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                review


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service    


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service\inte 
    rnal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service\inte 
    rnal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service\inte 
    rnal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\review-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\review-service\inte 
    rnal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:43                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\cmd\notification


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic
    e\cmd


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                notification


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\notification-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\notification-servic 
    e\internal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\cmd


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service     


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                cmd


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\cmd\admin


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service\cmd   


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                admin


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service     


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                internal


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\internal\handler


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                handler


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\internal\model


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                model


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\internal\repository


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                repository


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> mkdir services\admin-service\internal\service


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\admin-service\inter 
    nal


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> dir


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:37                config
d-----         2025/6/18      3:37                docker
d-----         2025/6/18      3:13                resource
d-----         2025/6/18      3:37                services
-a----         2025/6/18      3:31          32651 README.md


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> dir services


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:44                admin-service
d-----         2025/6/18      3:39                boarding-service
d-----         2025/6/18      3:44                notification-service
d-----         2025/6/18      3:43                order-service
d-----         2025/6/18      3:38                pet-service
d-----         2025/6/18      3:43                review-service
d-----         2025/6/18      3:37                user-service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> dir services\user-service /s
dir : 第二个路径段不得为驱动器或 UNC 名称。
参数名: path2
所在位置 行:1 字符: 1
+ dir services\user-service /s
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (E:\Development\...es\user-service:String) [Get-ChildItem]，Argum 
    entException
    + FullyQualifiedErrorId : DirArgumentError,Microsoft.PowerShell.Commands.GetChildItemCommand

PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> dir services\user-service -Recurse


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service      


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                cmd
d-----         2025/6/18      3:38                internal


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\cmd  


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                user


    目录: E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding\services\user-service\intern 
    al


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2025/6/18      3:38                handler
d-----         2025/6/18      3:38                model
d-----         2025/6/18      3:38                repository
d-----         2025/6/18      3:38                service


PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
PS E:\Development\DevTool\GitHub\CyanHub\EduDev\BlockChainAppDev\PetBoarding> 
```
