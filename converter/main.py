import pandas as pd

df = pd.read_csv("exemplo.csv", header=16)
print(df.to_string())