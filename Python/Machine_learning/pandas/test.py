import time
import os

# 获取进程的pid
pid = os.getpid()
print('pid: ', pid)

# 将pid写入本地文件
f1 = open(file='count_pid.txt', mode='w')
f1.write(pid.__str__())
f1.close()

# 开始计数并打印
n = 0
while True:
    n += 1
    print(n)
    time.sleep(1)
