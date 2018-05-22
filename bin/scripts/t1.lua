--[[print("hello yk")

for n in pairs(_G) do 
    print(n) 
end
--]]

--test require

local exports = {}
 
function exports.foo()
    print 'Hello World'
end
 
return exports