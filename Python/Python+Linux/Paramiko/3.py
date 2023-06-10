"""
上传文件到远程服务器
"""
# !/usr/bin/python
# -*- coding: utf-8 -*-
import os, sys
import paramiko

hostname = '192.168.1.215'
port = 22
username = 'root'
password = '123456'
local_path = '/root/test.txt'
remote_path = '/opt/test.txt'
if not os.path.isfile(local_path):
    print(local_path + " file not exist!")
    sys.exit(1)
try:
    s = paramiko.Transport((hostname, port))
    s.connect(username=username, password=password)
except Exception as e:
    print(e)
    sys.exit(1)
sftp = paramiko.SFTPClient.from_transport(s)
# 使用put()方法把本地文件上传到远程服务器
sftp.put(local_path, remote_path)
# 简单测试是否上传成功
try:
    # 如果远程主机有这个文件则返回一个对象，否则抛出异常
    sftp.file(remote_path)
    print("上传成功.")
except IOError:
    print("上传失败!")
finally:
    s.close()
