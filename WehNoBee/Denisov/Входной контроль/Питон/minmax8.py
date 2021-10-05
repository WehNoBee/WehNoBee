a = int(input())
chisla = int(input())
mine = 0
minperv = 1
minposl = 1
for i in range (chisla):
    if chisla < mine:
        mine = chisla
        minperv = i
        minposl = i
    else:
        if chisla == mine:
            minposl = i
print(minperv,minposl)