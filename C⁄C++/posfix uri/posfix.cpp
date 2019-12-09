#include <iostream>
#include <stack>
#include <queue>

int main() {

    std::stack <char> op;
    std::queue <char> res;
    char expression[310];
    int n=1;

    std::cin >> n;

    while(n) {
        std::cin >> expression;

        for(int i = 0; expression[i] != '\0'; i++){
            if(expression[i] == '+' || expression[i] == '-' || expression[i] == '*' ||
                    expression[i] == '/' || expression[i] == '^' )
                op.push(expression[i]);
            else if(expression[i] == '(')
                continue;
            else if(expression[i] == ')'){
                res.push(op.top());
                op.pop();
                continue;
            }
            else
                res.push(expression[i]);
        }

        while(!res.empty()){
            std::cout << res.front() << std::flush;
            res.pop();
        }
        std::cout << std::endl;

        n--;
    }

}
