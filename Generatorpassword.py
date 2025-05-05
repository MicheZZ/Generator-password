from random import *
app = ["0","1","2","3","4","5","6","7","8","9",]
bad = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]
cap = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]
den = ["!","@","#","№","$","%","^","&","?","*","_","=", "+"]

guf = [app,bad,cap, den]
fuck = [app,bad,cap]

g = False
while g != True:
    try:
        y = int(input("Введи длину пароля: "))
        g = True
    except ValueError:
        print('Ошибка! Введите целое число')


print("Надо использовать спец символы? Такие как '!#$@%&*?^_+=-'")
key2 = None
key3 = None
while key2 != True:
    answer = input()
    if answer.lower() != "да" and answer.lower() != "нет":
        print("Напиши да или нет")
    else:
        if answer.lower() == "да":
            key3 = True
            key2 = True
        else:
            key2 = True
            key3 = False
password = []
if key3 == True:
    for _ in range(y):
        k  = randint(0, 3)
        i = len(guf[k])-1
        j = randint(0, i)
        password.append(guf[k][j])
else:
    for _ in range(y):
        k = randint(0, 2)
        i = len(fuck[k]) -1
        j = randint(0, i)
        password.append(fuck[k][j])
print(*password, sep="")
input()