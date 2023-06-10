"""
私钥认证远程执行命令
"""
# !/usr/bin/python
# -*- coding: utf-8 -*-
import paramiko
import sys

hostname = '192.168.1.215'
port = 22
username = 'root'
key_file = '/root/.ssh/id_rsa'
cmd = " ".join(sys.argv[1:])


def ssh_conn(command):
    client = paramiko.SSHClient()
    key = paramiko.RSAKey.from_private_key_file(key_file)
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(hostname, port, username, pkey=key)
    stdin, stdout, stderr = client.exec_command(command)  # 标准输入，标准输出，错误输出
    result = stdout.read()
    error = stderr.read()
    if not error:
        print(result)
    else:
        print(error)
    client.close()


if __name__ == "__main__":
    ssh_conn(cmd)
