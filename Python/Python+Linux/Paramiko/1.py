"""
SSH密码认证远程执行命令
"""

# !/usr/bin/python
# -*- coding: utf-8 -*-
import paramiko

hostname = '192.168.199.135'
port = 22
username = 'root'
password = '19950308.zxc'

client = paramiko.SSHClient()  # 绑定实例
client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
client.connect(hostname, port, username, password, timeout=5)
print(f"Successfully login to {hostname}\n")

stdin, stdout, stderr = client.exec_command('df -h')  # 执行bash命令
result = stdout.read().decode()
error = stderr.read().decode()
# 判断stderr输出是否为空，为空则打印执行结果，不为空打印报错信息
if not error:
    print(result)
else:
    print(error)
client.close()
