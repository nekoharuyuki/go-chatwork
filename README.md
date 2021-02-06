# go-chatwork
[![Github issues](https://img.shields.io/github/issues/nekoharuyuki/go-chatwork)](https://github.com/nekoharuyuki/go-chatwork/issues)
[![Github forks](https://img.shields.io/github/forks/nekoharuyuki/go-chatwork)](https://github.com/nekoharuyuki/go-chatwork/network/members)
[![Github stars](https://img.shields.io/github/stars/nekoharuyuki/go-chatwork)](https://github.com/nekoharuyuki/go-chatwork/stargazers)
[![Github license](https://img.shields.io/github/license/nekoharuyuki/go-chatwork)](https://github.com/nekoharuyuki/go-chatwork/)

go-chatwork is a sample program that sends notifications and tasks to chatwork.
  
**Chatwork API documentation site**  
https://developer.chatwork.com/ja/

## Usage ##

Enter the room id and token in config.ini.

```
[chatwork]
roomid = XXXXXXXXXXXXXXXXX 
token  = XXXXXXXXXXXXXXXXX
```

**Create a task in ChatWork**

```go
chatwork.CreateTask(message, config.Config.Token, userids, config.Config.Roomid)
```

**Post a message to chatwork**

```go
chatwork.Post(message, config.Config.Token, config.Config.Roomid)
```
