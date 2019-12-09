fact :: Integer -> Integer
fact 0 = 1
fact n = n * fact (n-1)

main = do
    putStrLn "factorial of 5"
    print(fact 100)
