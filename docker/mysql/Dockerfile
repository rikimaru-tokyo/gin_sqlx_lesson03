FROM mysql:latest

ADD docker/mysql/conf.d/my.cnf /etc/mysql/my.cnf
RUN chmod 644 /etc/mysql/my.cnf
# ↑ Windows+Virtualbox環境(共有フォルダ使用)では、my.cnfのパーミッションがどうしても777になってしまい、
#   読み込みができず文字化けします。
#   当Dockerfileはmy.cnfを確実に読み込ませるための回避策として定義しております。
#      参考：https://qiita.com/koyo-miyamura/items/4d1430b9086c5d4a58a5
