import os


def kill(pid):
    if os.name == 'nt':
        # Windows系统
        cmd = "taskkill /pid " + str(pid) + "/f"
        try:
            os.system(cmd)
            print(pid, "killed")
        except Exception as e:
            print(e)
    elif os.name == 'posix':
        # Linux系统
        cmd = 'kill ' + str(pid)
        try:
            os.system(cmd)
        except Exception as e:
            print(pid, 'killed')
            print(e)
    else:
        print('Undefined os.name')
