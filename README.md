# mongoBackup
备份阿里云mongo数据至本地或其他服务器



```bash
[root@iZuf6238xylvm0t7dxz4l9Z ~]# mongoback -h
NAME:
   backup - 下载阿里云中mongo的备份文件

USAGE:
   mongoback [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --mongo value, -m value  mongo的实例ID，多个实例以英文逗号隔开
   --date value, -d value   获取哪天的备份，默认获取当天凌晨的备份文件，格式：yyyy-MM-dd (default: 2021-07-14)
   --hour value, -H value   在控制台设置的备份时间段，默认是：03:00-04:00 (default: 03:00-04:00)
   --path value, -p value   指定备份文件存储路径 (default: /data/mongo-backup)
   --internal, -i           是否使用内网下载，默认使用公网地址 (default: false)
   --help, -h               show help (default: false)
```


