# WebWork

下面的操作不对
//操作，将这个镜像移动到unraid
//先生成image
//sh ./make_image.sh
//再导出镜像
//docker save e315365fa787(镜像ID) > Desktop/demo.tar

VPS 拉 mysql docker
[root@144 ~]# docker pull mysql:5.7.28
Digest: sha256:b38555e593300df225daea22aeb104eed79fc80d2f064fde1e16e1804d00d0fc
Status: Downloaded newer image for mysql:5.7.28
docker.io/library/mysql:5.7.28

[root@144 ~]# docker images
REPOSITORY                     TAG       IMAGE ID       CREATED         SIZE
jonssonyan/trojan-panel-core   latest    1de02ac2ce9a   17 months ago   313MB
jonssonyan/trojan-panel-ui     latest    88cfb161d5dd   17 months ago   29.6MB
jonssonyan/trojan-panel        latest    53e760c54202   17 months ago   41.4MB
redis                          6.2.7     4b1123a829a1   2 years ago     113MB
caddy                          2.6.2     006d393a4e6a   2 years ago     46.8MB
mariadb                        10.7.3    daf0f023c28d   3 years ago     414MB
mysql                          5.7.28    db39680b63ac   5 years ago     437MB


启动mysql
docker run --name mysql5.7 -p 3307:3306 -v /data/mysql/datadir:/var/lib/mysql -v /data/mysql/conf.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=stanhu520 -d mysql:5.7.28



# 示例：创建 2GB 的 Swap 文件（替换路径和大小）
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile  # 设置严格权限
sudo mkswap /swapfile     # 格式化为 Swap
sudo swapon /swapfile     # 启用 Swap
如果 fallocate 报错可以用，再rvtfrgdm的代码就行
# 正确示例：使用更小的缓冲区（bs=1M）创建 2GB Swap 文件
sudo dd if=/dev/zero of=/swapfile bs=1G count=2 status=progress

dd命令报
报 dd: memory exhausted by input buffer of size 1073741824 bytes (1.0 GiB)
可以改用
sudo dd if=/dev/zero of=/swapfile bs=1M count=2048 status=progress