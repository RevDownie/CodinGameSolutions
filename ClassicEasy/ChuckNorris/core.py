import sys, math, itertools

def getBlockType(in_bit):
    if in_bit == '1':
        return "0"
    else:
        return "00"

message = input()

binary = ' '.join(format(ord(x), '07b') for x in message).replace(" ", "")
#print(binary)

unary = ""

#00 0 0 0 00 00 0 0 00 0 0 0

for key, items in itertools.groupby(binary):
    unary += getBlockType(key) + " "
    for item in items:
        unary += "0"
    unary += " "

print(unary[:-1])