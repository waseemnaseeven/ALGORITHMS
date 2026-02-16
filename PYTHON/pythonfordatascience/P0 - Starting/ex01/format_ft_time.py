import time
from datetime import datetime

# Get the current time in seconds since the epoch
current_time = time.time()

# Format the time in seconds since January 1, 1970
formatted_time = f"Seconds since January 1, 1970: {current_time:,.4f} or {current_time:.2e} in scientific notation"

# Get the current date
current_date = datetime.now()

# Format the current date in the desired format
formatted_date = current_date.strftime("%b %d %Y")

# Print the formatted results
print(formatted_time)
print(formatted_date)
