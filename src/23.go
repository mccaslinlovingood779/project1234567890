import numpy as np
from scipy.optimize import minimize

def f(x):
    y = x[0]**2 + 3*x[1] - 5
    return y

x_0 = [1, 1]
result = minimize(f, x_0)

print("Optimized solution:", result.x)
