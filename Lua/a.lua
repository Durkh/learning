a= {}

--[[for i = 0, 5 do
a[i] = io.read();
end

i = 1
par = 0

while(i <=5) do
   if(a[i] %2 == 0) then
      par= par+1;
   end
   i=i+1;
end

print( par, "valores pares")
--]]

print("---------------------")

print(a.lua)

a.lua = 5
