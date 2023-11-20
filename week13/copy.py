import copy
# 파이썬을 통한 깊은 복사(deepCopy)와 얕은 복사(shallowCopy)
a=[3,-9,11]
# a=[3,[5,-71],11]
b=a #shallow copy
c=a[:] #shallow copy
d=list(a) #shallow copy
e=a.copy() #shallow copy
f=copy.deepcopy(a) #deep copy

print(a,b,c,d,e,f)
a[1]=100
# a[1][0]=100
print(a,b,c,d,e,f)