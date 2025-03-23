import numpy as np

def sample_data():
    # Create some synthetic data
    x = np.random.rand(100)
    y = 2 * x + 3 + np.random.randn(100)

    return x, y

# Generate random samples for training purposes
x_train, y_train = sample_data()
print("Training Data:")
print(x_train)
print(y_train)
