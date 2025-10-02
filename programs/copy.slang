[A] IF X != 0 GOTO B
    Z9 <- Z9 + 1
    IF Z9 != 0 GOTO C
[B] X <- X -1
    Y <- Y + 1
    Z <- Z + 1
    Z8 <- Z8 + 1
    IF Z8 != 0 GOTO A
[C] IF Z != 0 GOTO D
    Z7 <- Z7 + 1
    IF Z7 != 0 GOTO E
[D] Z <- Z - 1
    X <- X + 1
    Z6 <- Z6 + 1
    Z99 <-  COPY(X)
    IF Z6 != 0 GOTO C

