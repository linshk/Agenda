# Agenda
一个轻量级的会议管理命令行工具。

## 环境要求
- Go

## 安装说明
```bash
$ go get github.com/linshk/Agenda
$ go install github.com/linshk/Agenda
# 测试是否安装成功
$ Agenda -h
```

## 功能
```bash
Agenda is a meeting management application written in Go.
This application is a tool to manage your meetings efficiently

Usage:
  Agenda [command]

Available Commands:
  add          add participators to the meeting
  cancel       cancel a meeting sponsored by current user
  clear        clear all meetings sponsored by current user
  create       create a meeting
  deleteuser   delete a user
  help         Help about any command
  login        login a user
  logout       logout a user
  mm           A brief description of your command
  querymeeting query meetings by time interval
  queryuser    view all registered users
  quit         quit a meeting participated by current user
  register     register a new user
  remove       remove participators from the meeting

Flags:
      --config string   config file (default is $HOME/.Agenda.yaml)
  -h, --help            help for Agenda

Use "Agenda [command] --help" for more information about a command
```

### 用户注册
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

### 用户登录
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

### 用户登出
1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。
```bash
logout a user

Usage:
  Agenda logout [flags]

Flags:
  -h, --help   help for logout
```

### 用户查询
1. 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。
```bash
view all registered users with username, email and phone

Usage:
  Agenda queryuser [flags]

Flags:
  -h, --help   help for queryuser
```

### 用户删除

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

### 创建会议

1. 已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
- 会议主题(title)（在会议列表中具有唯一性）
- 会议参与者(participator)
- 会议起始时间(start time)
- 会议结束时间(end time)
2. 注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
3. 用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

```bash
create a meeting, with a unique title, a list of participators (a list of registered usernames separated by comma), and with start time and end time (format:2006-01-02T15:04:05).
  Meetings with conflicts are not allowed

Usage:
	Agenda create -t [title] -s [start time] -e [end time] -p [participators]

Flags:
  -t,  --title        string  title of meeeting
  -s,  --starttime    string  start time of meeting, format required: "2018-10-01T08:00:00"
  -e,  --endtime      string  end time of meeting, format required: "2018-10-01T08:00:00"
  -p,  --partipators  string  a list of partipators (identify each user by his name), use comma to separate
  -h,  --help                 help for create
```

### 增删会议参与者

1. 已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
2. 增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
3. 删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。


```bash
add/remove participators to/from the meeting that specified by the title, the meeting will be canceled if there is no participator after removing

Usage:
	Agenda add/remove -t [title] -p [participators]

Flags:		
  -t,  --title         string  title of meeeting
  -p,  --participators string  a list of partipators (identify each user by his name), use comma to separate	
  -h,  --help                  help for add/remove				
```

### 查询会议

1. 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
2. 用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
3. 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
4. 注意，查询会议的结果应包括用户作为 发起者或参与者的会议。

```bash
query meetings that current user participates or sponsors by time interval

Usage:
  Agenda querymeeting -s [start time] -e [end time]

Flags:
  -s,  --starttime   string  start time of the interval, format required: "2006-01-02T15:04:05"
  -e,  --endtime     string  end time of the interval, format required: "2006-01-02T15:04:05" 
  -h,  --help                help for querymeeting
```

### 取消会议

1. 已登录的用户可以取消自己发起的某一会议安排。
2. 取消会议时，需提供唯一标识：会议主题（title）。

```bash
cancel a meeting sponsored by current user

Usage:
  Agenda cancel -t [title]

Flags:
  -t,  --title  string  title of the meeting
  -h,  --help           help for cancel
```

### 退出会议

1. 已登录的用户可以退出自己参与的某一会议安排。
2. 退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议参与者人数为0，则会议也将被删除。

```bash
quit a meeting participated by current user, the meeting will be canceled if there is no participator after quiting

Usage:
  Agenda quit -t [title]

Flags:
  -t,  --title  string  title of the meeting
  -h,  --help           help for quit
```

### 清空会议

1. 已登录的用户可以清空 自己发起 的所有会议安排。

```bash
clear all meetings sponsored by current user

Usage:
  Agenda clear

Flags:
  -h,  --help    help for clear
```
