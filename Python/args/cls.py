class A:
    def run(self):
        pass

    def ok(self):
        pass


if __name__ == '__main__':
    a = A()
    print(type(a.__getattribute__("run")))
    print(a.__getattribute__("run"))
