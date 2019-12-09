fact :: Integer -> Integer
fact n | n == 0 = 1
       | n /= 0 = n * fact(n-1)
       
roots :: (Float, Float, Float) -> (Float, Float)
roots(a,b,c) = (x1,x2) where
    x1 = e + sqrt d / (2*a)
    x2 = e - sqrt d / (2*a)
    d = (b^2) - (4*a*c) 
    e = -b/(2*a)
       
main = do 
    putStrLn "factorial of 5"
    print (fact 5)
    putStrLn "roots of the equation x²-8x+6"
    print (roots(1,-8,6))
