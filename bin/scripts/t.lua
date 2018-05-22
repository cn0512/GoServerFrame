package.path = package.path ..';bin\\scripts\\?.lua';
print(package.path)
local t = require('t1')
print(t)
t.foo()