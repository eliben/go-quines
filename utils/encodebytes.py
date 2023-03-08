# This script helps generate byte-encoded quines
import fileinput

cs = []
for line in fileinput.input():
    for c in line:
        cs.append(c)

print(", ".join(str(ord(c)) for c in cs))
