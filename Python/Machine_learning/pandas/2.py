# 1.sys.exit()
# 执行该语句会直接退出程序，这也是经常使用的方法，也不需要考虑平台等因素的影响，一般是退出Python程序的首选方法。
# 该方法中包含一个参数status，默认为0，表示正常退出，也可以为1，表示异常退出。
# 该方法引发的是一个SystemExit异常(这是唯一一个不会被认为是错误的异常)，当没有设置捕获这个异常将会直接退出程序执行，当然也可以捕获这个异常进行一些其他操作。

# import sys
# from time import sleep
#
# for i in range(10):
#     print("正在运行中。。。", i)
#     sleep(1)
#     try:
#         if i == 5:
#             print("准备退出程序")
#             sys.exit()
#     except SystemExit:
#         print("程序退出啦。。。")

# 2、os._exit()
# 效果也是直接退出，不会抛出异常，但是其使用会受到平台的限制，但我们常用的Win32平台和基于UNIX的平台不会有所影响。
# import sys
# from time import sleep
# import os
#
# for i in range(10):
#     print("正在运行中。。。", i)
#     sleep(1)
#     if i == 5:
#         print("准备退出程序")
#         os._exit(0)

# 3、os.kill()
# 一般用于直接Kill掉进程，但是只能在UNIX平台上有效。
# 基本原理：该函数是模拟传统的UNIX函数发信号给进程，其中包含两个参数：一个是进程名，即所要接收信号的进程；一个是所要进行的操作。
# 操作（第二个参数）的常用取值为：
# SIGINT         终止进程        中断进程
# SIGTERM        终止进程        软件终止信号
# SIGKILL        终止进程        杀死进程
# SIGALRM        闹钟信号


# 4、Windows下Kill进程
# 既然在Linux下能够进行上述操作，那么Windows下也能够有相关的操作。
# 这里使用的是os.popen()，该方法是用于直接执行系统命令，而在Windows下其实就是使用taskkill来kill掉进程，其基本形式是，
# taskkill/pid程序的PID号码
# 可以直接在CMD窗口下试下这个命令。
# 可以先打开一个计算器程序，然后使用tasklist查看该程序的pid，这里是620，所以对应的Python代码是:
# import os
#
# if __name__ == '__main__':
#     # pid = 20292
#     os.popen('taskkill.exe /pid:' + str(pid))
