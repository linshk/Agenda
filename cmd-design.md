## 用户注册
1. 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
2. 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。
```bash
register a new user, a unique username, a password, an email and a phone required

Usage:
  Agenda register [flags]

Flags:
  -c, --contact string    phone number
  -e, --email string      email address
  -h, --help              help for register
  -p, --password string   password
  -u, --username string   username
```
## 用户登录
1. 用户使用用户名和密码登录 Agenda 系统。
2. 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。
```bash
login a user, a username and a password required

Usage:
  Agenda login -u [username] -p [password] -e [email] -c [phone] [flags]

Flags:
  -h, --help              help for login
  -p, --password string   password
  -u, --username string   username
```
## 用户登出
1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。
```bash
logout a user

Usage:
  Agenda logout [flags]

Flags:
  -h, --help   help for logout
```
## 用户查询
1. 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。
```bash
view all registered users with username, email and phone

Usage:
  Agenda queryuser [flags]

Flags:
  -h, --help   help for queryuser
```
## 用户删除

1. 已登录的用户可以删除本用户账户（即销号）。
2. 操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
3. 删除成功则退出系统登录状态。删除后，该用户账户不再存在。
4. 用户账户删除以后：
    - 以该用户为 发起者 的会议将被删除
    - 以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。
```bash
deleteuser a user

Usage:
  Agenda deleteuser [flags]

Flags:
  -h, --help   help for deleteuser
```

```bash
zjq:Agenda zjq$ go run main.go deleteuser
[Cmd]   2018/10/27 20:10:15 deleteuser
[Error] 2018/10/27 20:10:15 no user has been logined, login and try again
zjq:Agenda zjq$ go run main.go register -u me -p p -e me@qq.com -c 110
[Cmd]   2018/10/27 20:11:40 register --username=me --password=p --email=me@qq.com --phone=110
[OK]    2018/10/27 20:11:40 register successfully
zjq:Agenda zjq$ go run main.go register -u me -p p -e me@qq.com -c 110
[Cmd]   2018/10/27 20:11:42 register --username=me --password=p --email=me@qq.com --phone=110
[Error] 2018/10/27 20:11:42 username already exists
zjq:Agenda zjq$ go run main.go register -u me -p p -e me@qq.com 
[Cmd]   2018/10/27 20:11:53 register --username=me --password=p --email=me@qq.com --phone=
[Error] 2018/10/27 20:11:53 a username, a password, an email and a phone required
zjq:Agenda zjq$ go run main.go register -u you -p p -e you@qq.com -c 120
[Cmd]   2018/10/27 20:12:23 register --username=you --password=p --email=you@qq.com --phone=120
[OK]    2018/10/27 20:12:23 register successfully
zjq:Agenda zjq$ go run main.go register -u he -p p -e he@qq.com -c 119
[Cmd]   2018/10/27 20:12:37 register --username=he --password=p --email=he@qq.com --phone=119
[OK]    2018/10/27 20:12:37 register successfully
zjq:Agenda zjq$ go run main.go login -u me -p p
[Cmd]   2018/10/27 20:12:54 login --username=me --password=p
[OK]    2018/10/27 20:12:54 login successfully
zjq:Agenda zjq$ go run main.go login -u me -p p
[Cmd]   2018/10/27 20:12:57 login --username=me --password=p
[Error] 2018/10/27 20:12:57 a user has been logined, logout and try again
zjq:Agenda zjq$ go run main.go logout
[Cmd]   2018/10/27 20:13:11 logout
[OK]    2018/10/27 20:13:11 logout successfully
zjq:Agenda zjq$ go run main.go logout
[Cmd]   2018/10/27 20:13:13 logout
[Error] 2018/10/27 20:13:13 no user has been logined, login and try again
zjq:Agenda zjq$ go run main.go login -u me -p a
[Cmd]   2018/10/27 20:13:22 login --username=me --password=a
[Error] 2018/10/27 20:13:22 username or password error
zjq:Agenda zjq$ go run main.go login -u a -p a
[Cmd]   2018/10/27 20:13:25 login --username=a --password=a
[Error] 2018/10/27 20:13:25 username or password error
zjq:Agenda zjq$ go run main.go login -u me -p p
[Cmd]   2018/10/27 20:13:31 login --username=me --password=p
[OK]    2018/10/27 20:13:31 login successfully
zjq:Agenda zjq$ go run main.go queryuser
[Cmd]   2018/10/27 20:13:45 queryuser
[OK]    2018/10/27 20:13:45 
         username               email          phone
 1             me           me@qq.com            110
 2            you          you@qq.com            120
 3             he           he@qq.com            119
zjq:Agenda zjq$ go run main.go logout
[Cmd]   2018/10/27 20:13:55 logout
[OK]    2018/10/27 20:13:55 logout successfully
zjq:Agenda zjq$ go run main.go queryuser
[Cmd]   2018/10/27 20:13:57 queryuser
[Error] 2018/10/27 20:13:57 no user has been logined, login and try again
zjq:Agenda zjq$ go run main.go login -u me -p p
[Cmd]   2018/10/27 20:14:04 login --username=me --password=p
[OK]    2018/10/27 20:14:04 login successfully
zjq:Agenda zjq$ go run main.go deleteuser
[Cmd]   2018/10/27 20:14:18 deleteuser
[OK]    2018/10/27 20:14:18 deleteuser successfully
zjq:Agenda zjq$ go run main.go deleteuser
[Cmd]   2018/10/27 20:14:20 deleteuser
[Error] 2018/10/27 20:14:20 no user has been logined, login and try again
zjq:Agenda zjq$ go run main.go login -u me -p p
[Cmd]   2018/10/27 20:14:28 login --username=me --password=p
[Error] 2018/10/27 20:14:28 username or password error
zjq:Agenda zjq$ go run main.go login -u you -p p
[Cmd]   2018/10/27 20:14:33 login --username=you --password=p
[OK]    2018/10/27 20:14:33 login successfully
zjq:Agenda zjq$ go run main.go queryuser
[Cmd]   2018/10/27 20:14:43 queryuser
[OK]    2018/10/27 20:14:43 
         username               email          phone
 1            you          you@qq.com            120
 2             he           he@qq.com            119
```