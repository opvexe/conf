# MySQL备份与恢复

备份

```
mysqldump -uroot -p123456 "数据库名称" > "文件路径"(./back.sql)

```

恢复

```
mysql -uroot -p123456 back(恢复的表名) < back.sql
```