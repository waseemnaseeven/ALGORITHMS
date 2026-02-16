import numpy as np

def give_bmi(height: list[int | float], weight: list[int | float]) -> list[int | float]:
    data = np.array((height, weight))
    L = []
    for i in range(data.shape[1]):
        h = data[0, i]
        w = data[1, i]

        # Convert height to meters for BMI calculation
        tmp = float(w / (h/100)**2)
        bmi = round(tmp, 2)
        L.append(bmi)
    return L

def apply_limit(bmi: list[int | float], limit: int) -> list[bool]:
    return [value > limit for value in bmi]

if __name__ == "__main__":
    height = [271, 115]
    weight = [165.3, 38.4]
    bmi = give_bmi(height, weight)
    print(bmi, type(bmi))

    limit = 26
    limit_result = apply_limit(bmi, limit)
    print(limit_result)
