import stack
s = stack.Stack()
print("Добро пожаловать в нашу консоль для стека")
while True:
     cmd = input("sas ")
     cmd_w = cmd.split()
     try:
          if cmd_w[0] == "push":
               s.push(cmd_w[1])
          elif cmd_w[0] == "pop":
               print(s.pop())
          elif cmd_w[0] == "dump":
               for i,x in enumerate(s.dump()):
                    print(i,x)
          elif cmd_w[0] == "exit":
               break
     except:
          print("Произошла ошибка")
          
          
