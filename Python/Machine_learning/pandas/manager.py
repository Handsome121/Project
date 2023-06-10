from kill import kill

f1 = open(file='count_pid.txt', mode='r')
pid = f1.read()
f1.close()

kill(pid=pid)
