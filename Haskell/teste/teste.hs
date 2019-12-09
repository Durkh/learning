
fType :: Int -> Int -> Int
fType x y = x*x + y*y

main = do
    
    let var = 25
    if var `rem` 2 == 0
       then putStrLn "Even"
       else putStrLn "Odd"
       
    print()
    print()
    print(fType 2 4)
    print(show var)
    print (readInt (show var))
    print (surface $ Circle 10 20 10)
    
readInt :: String -> Int
readInt = read 

data Area = Circle Float Float Float
surface :: Area -> Float
surface (Circle _ _ r) = pi * r ^ 2
