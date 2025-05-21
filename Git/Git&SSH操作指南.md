# Git与SSH配置完全指南

## 一、Git基础配置
##### 1. 安装Git

访问 [Git官网](https://git-scm.com/downloads) 下载并安装最新版本的Git。

##### 2. 配置用户信息

打开PowerShell（右键点击开始菜单选择"Windows PowerShell (管理员)"），设置全局用户名和邮箱：

```bash
git config --global user.name "YourGitHubUsername"
git config --global user.email "your.email@example.com"
```

将`"YourGitHubUsername"`和`"your.email@example.com"`替换为自己的用户名和邮箱



### 3. 查看配置信息git config --list

## 二、SSH密钥配置

1. ##### 检查已有SSH密钥

   ```bash
   ls ~/.ssh
   ```

   如果看到`id_rsa.pub`或`id_ed25519.pub`等文件，表示已有SSH密钥。

2. ##### 生成新的SSH密钥

   ```bash
   ssh-keygen -t ed25519 -C "your.email@example.com"
   ```

   按提示操作：

- 直接回车使用默认路径
- 输入密码（可选但推荐）

3. ##### 启动SSH代理

4. ##### 启动SSH代理服务

​	

```bash
eval "$(ssh-agent -s)"
```



##### 将SSH密钥添加到代理

​	

```bash
ssh-add ~/.ssh/id_ed25519
```



## 三、GitHub配置

#### 1. 复制SSH公钥到剪贴板

##### Windows系统

```bash
cat ~/.ssh/id_ed25519.pub | clip
```



##### Linux/macOS系统

```sh
pbcopy < ~/.ssh/id_ed25519.pub
```



#### 2. 添加SSH密钥到GitHub

1. 登录**GitHub**，点击右上角头像 → **Settings** → **SSH and GPG keys**
2. 点击"**New SSH key**"
3. Title填写"**Personal Laptop**"等描述
4. **Key**字段粘贴刚才复制的内容
5. 点击"**Add SSH key**"

#### 3. 测试SSH连接

```bash
ssh -T git@github.com
```



##### 如果看到以下内容，表示连接成功：

```bash
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```



## 四、Git仓库操作

#### 1. 克隆远程仓库

##### 使用**SSH**协议

```bash
git clone git@github.com:username/repo-name.git
```



#### 2. 添加文件到暂存区

```bash
git add .  # 添加所有文件

git add filename  # 添加指定文件
```



##### 3. 提交到本地仓库

```bash
git commit -m "提交说明"
```



##### 4. 推送到远程仓库

```bash
git push origin branch-name
```



##### 5. 从远程仓库拉取

```bash
git pull origin branch-name
```



## 五、常见问题解决
#### 1. 权限被拒绝

##### 检查SSH密钥是否正确添加到GitHub

```bash
ssh -T git@github.com
```

#####  重新生成SSH密钥并添加到GitHub



#### 2. 忘记SSH密码

##### 删除原有密钥

```bash
rm ~/.ssh/id_ed25519*
```



##### 重新生成密钥

```bash
ssh-keygen -t ed25519 -C "your.email@example.com"
```



#### 3. 切换HTTPS到SSH协议# 查看当前远程地址

```bash
git remote -v
```



##### 修改为SSH协议

```bash
git remote set-url origin git@github.com:username/repo-name.git
```



# 六、高级配置

#### 1. 配置多个GitHub账户

创建`~/.ssh/config`文件：

##### 主账户

```bash
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519
```



##### 工作账户

```bash
  Host github-work
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_work
```

##### 克隆时使用：

```bash
git clone git@github-work:username/repo-name.git
```



#### 2. 配置自动保存密码# 保存密码到Windows凭据管理器

```bash
git config --global credential.helper wincred
```



#### 3. 配置Git别名

##### **简化常用命令**

```bash
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```

##### **使用方法：**

```bash
git co master  # 等同于git checkout master`
```

