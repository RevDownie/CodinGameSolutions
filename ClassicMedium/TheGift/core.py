# RULES
# Determine how much each person contributes to a cost based on their budget.
# They cannot pay more than the budget
# Payment must be made in integers
# The aim is to make the highest payment as low as possible (this means that people with low budgets should pay the most percentage wise)
# If payment cannot be made in full return IMPOSSIBLE otherwise output the payment amount for each person

import sys

# Logging to stdout is reserved for returning the answer. To debug log
# we must log to stderr
#
def debug_log(msg):
	sys.stderr.write(str(msg) + '\n')

# Run the simulation and output the answer to codingame
# Note: the given budgets are unsorted
# 
def run(gift_amount, budgets, expected_output = None):

	# Check to make sure they have enough in total to pay the gift
	total_budget = sum(budgets)
	if total_budget < gift_amount:
		output = "IMPOSSIBLE"
		if expected_output != None:
			debug_log("SUCCESS" if expected_output == output else "FAILED")

		# Send the answer to CodinGame via stdout
		print(output)
		return

	# Each person pays, staring with the person with the lowest budget who pays the highest percent relative to their budget
	# as the puzzle wants us to return the lowest possible highest (so we make the high budget peeps pay less)
	remaining = gift_amount
	budgets.sort()

	# NOTE: It seems that the output order should match the sorted order and not the original input order (which is convenient)
	output = [0] * len(budgets)
	for i in range(0, len(budgets)):
		should_pay = remaining/(len(budgets)-i)
		pays = min(budgets[i], should_pay)
		remaining -= pays
		output[i] = pays

	if expected_output != None:
		debug_log("SUCCESS" if expected_output == output else "FAILED")

	# Send the answer to CodinGame via stdout
	for o in output:
		print(o)


# Used when running in CodinGame to read input from stdinput
#
if __name__ == "__main__":
	# N is the number of people/budgets
	n = int(raw_input())

	# C is the total gift amount
	c = int(raw_input())

	expected_output = None

	# The amount each person has to spend (unordered)
	budgets = []
	for i in xrange(n):
		budgets.append(int(raw_input()))

	run(c, budgets, expected_output)