"""
上传目录到远程服务器
"""

# !/usr/bin/python
# -*- coding: utf-8 -*-
import os, sys
import paramiko

hostname = '192.168.1.215'
port = 22
username = 'root'
password = '123456'
local_path = '/root/abc'
remote_path = '/opt/abc'
# 去除路径后面正斜杠
if local_path[-1] == '/':
    local_path = local_path[0:-1]
if remote_path[-1] == '/':
    remote_path = remote_path[0:-1]
file_list = []
if os.path.isdir(local_path):
    for root, dirs, files in os.walk(local_path):
        for file in files:
            # 获取文件绝对路径
            file_path = os.path.join(root, file)
            file_list.append(file_path)
else:
    print("Directory not exist!")
    sys.exit(1)
try:
    s = paramiko.Transport((hostname, port))
    s.connect(username=username, password=password)
    sftp = paramiko.SFTPClient.from_transport(s)
except Exception as e:
    print(e)
for local_file in file_list:
    # 替换目标目录
    remote_file = local_file.replace(local_path, remote_path)
    remote_dir = os.path.dirname(remote_file)
    # 如果远程服务器没目标目录则创建
    try:
        sftp.stat(remote_dir)
    except IOError:
        sftp.mkdir(remote_dir)
    print("%s -> %s" % (local_file, remote_file))
    sftp.put(local_file, remote_file)
s.close()
