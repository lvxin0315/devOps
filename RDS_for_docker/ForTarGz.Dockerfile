FROM mysql:5.7 AS base1

ENV TZ Asia/Shanghai

RUN echo 'Asia/Shanghai' >/etc/timezone
RUN mkdir /base
COPY ./ /base
RUN chmod +x /base/*

#修改阿里源
RUN /base/setup1.sh


FROM base1 AS base2
#安装xtrabackup
RUN apt-get -y install curl lsb-release libdbd-mysql-perl rsync libcurl4-openssl-dev libaio1 libcurl4 libev4 libnuma1
RUN dpkg -i /base/percona-xtrabackup-24_2.4.20-1.buster_amd64.deb

#安装qpress
RUN tar xvf /base/qpress-11-linux-x64.tar
RUN chmod 775 qpress
RUN cp qpress /usr/bin

#恢复rds备份
FROM base2 AS base3

RUN mkdir -p /home/mysql/data

#处理压缩包
RUN tar -izxvf /base/hins_data.tar.gz -C /home/mysql/data

RUN innobackupex --defaults-file=/home/mysql/data/backup-my.cnf --apply-log /home/mysql/data

#阿里云文档中要求注释内容，https://help.aliyun.com/knowledge_detail/41817.html

#innodb_log_checksum_algorithm
#innodb_fast_checksum
#innodb_log_block_size
#innodb_doublewrite_file
#rds_encrypt_data
#innodb_encrypt_algorithm
#redo_log_version
#master_key_id
#server_uuid

RUN sed -i "s/innodb_log_checksum_algorithm/#innodb_log_checksum_algorithm/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/innodb_fast_checksum/#innodb_fast_checksum/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/innodb_log_block_size/#innodb_log_block_size/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/innodb_doublewrite_file/#innodb_doublewrite_file/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/rds_encrypt_data/#rds_encrypt_data/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/innodb_encrypt_algorithm/#innodb_encrypt_algorithm/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/redo_log_version/#redo_log_version/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/master_key_id/#master_key_id/g"  /home/mysql/data/backup-my.cnf
RUN sed -i "s/server_uuid/#server_uuid/g"  /home/mysql/data/backup-my.cnf
RUN chown -R mysql:mysql /home/mysql/data

#移动mysql数据文件
RUN mv  /base/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf


