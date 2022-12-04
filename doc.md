# 开发教程


## 去掉黑色命令提示框

```
# 增加编译参数
-ldflags -H=windowsgui

# 完整编译命令
go build -ldflags -H=windowsgui
```


## 设置程序图标

为生成的程序添加图标
```
fyne package -icon icon.png
```
