# RULES
# Each line displays the run-length encoding of the previous line e.g. 1 => 1 1 => 2 1 => 1 2 1 1

# INPUT
# R - The starting number of line 1
# L - The line to display (the first line is 1)

# OUTPUT
# The contents of the line numbered L as a single space separated line

# NOTES
# 0 < R < 100
# 0 < L <= 25

# TO CONSIDER
# Is there a formulaic way of doing this or do we just brute force from the root node?

import sys

# Logging to stdout is reserved for returning the answer. To debug log
# we must log to stderr
#
def debug_log(msg):
	sys.stderr.write(str(msg) + '\n')

# Convert the given line into run-length encoding e.g. 1,1,1 becomes 3,1
# Input and outputs are lists of integers
#
def run_encode(line):
	encoded = []
	count = 0
	val = line[0]
	
	for i in range(0, len(line)):
		if line[i] != val:
			encoded.append(count)
			encoded.append(val)
			count = 1
			val = line[i]
		else:
			count = count + 1

	encoded.append(count)
	encoded.append(val)
	return encoded

# Find the line l for starting input r
# 
def run(r , l, expected = None):
	line = [r]

	for i in range(1, l):
		line = run_encode(line)
		debug_log(line)
	
	# Convert to the output format and submit to CodinGame
	as_strings = map(str, line)
	output_line = ' '.join(as_strings)
	if expected != None:
		debug_log("SUCCESS" if expected == output_line else "FAILED")
	print(output_line)

# Used when running in CodinGame to read input from stdinput
#
if __name__ == "__main__":
	r = int(raw_input())
	l = int(raw_input())

	run(r, l)