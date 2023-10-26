import subprocess
import pytest


def test_01():
    cmd = ["ping", "www.baidu.com"]
    str1 = "往返"
    str2 = "最短"
    stdout, stderr = exec_command(cmd)
    assert is_exist_in_str(stdout, str1, str2) == 1, "failed"


def exec_command(cmd):
    proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, encoding="gbk")
    stdout, stderr = proc.communicate()

    return stdout, stderr


def is_exist_in_str(str0, str1, str2):
    index1 = str0.find(str1)
    index2 = str0.find(str2)

    if index1 > 0 and index2 > 0:
        return 0
