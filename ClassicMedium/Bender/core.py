# SYMBOLS
# @ = Start location
# $ = End/Death location
# # ,X = Obstacles
# N,S,E,W = Direction modifiers
# I = Priority reverser
# B = Modifer to allow destruction of 'X' obstacles only - Toggle
# T = Teleporter - portals a 2 locations, doubly linked

# RULES
# 1. When bender encounters an obstacle he changes direction - in priority order - S,E,N,W
# 2. Encountering I will reverse priority order
# 3. Encountering a direction modifier will force a direction change
# 4. Encountering B will toggle X destruction on and off (B is not consumed)
# 5. Encountering T will teleport but maintain direction
# 6. Don't change direction on a blank cell i.e. keep going in the same direction

# INPUT
# Grid of CxL as lines with unbreakable surround. Max is 100x100

# OUTPUT
# Display the sequnce of moves taken for each cell - SOUTH, EAST, etc newline separated
# If the $ cannot be reached return "LOOP"

# STATE
# Direction priority - which changes based on whether an I has been encountered
# X destruction ability - which changes based on whether a B has been encountered
# Grid - X can be destroyed
# Steps - Steps taken til now
# Current location - x,y within grid
# Current direction
# Number of broken obstacles - used for loop state checking

# TO CONSIDER
# How do we detect a loop? - If we enter a cell with the same state more than once?
# Can we combine the loop check player state into a single value - inverted, breaker mode, direction, num broken obstacles?

# NOTES
# Grid origin (0,0) is top left (NW corner)

import sys
from itertools import ifilter

# Data container to holds grid state
#
class Grid:
	def __init__(self, cells, w, h):
		self.cells = cells
		self.history = map(lambda c: set(), cells)
		self.width = w
		self.height = h

# Logging to stdout is reserved for returning the answer. To debug log
# we must log to stderr
#
def debug_log(msg):
	sys.stderr.write(str(msg) + '\n')

# Parse the dimensions of the grid in the format "NumLines NumCols" into
# a tuple of actual numbers
# 
def parse_dims(unparsed):
	return tuple(map(int, unparsed.split()))

# Parse the grid lines into a grid structure of the given width and height
# 
def parse_grid(unparsed, w, h):
	return Grid([c for l in unparsed for c in l], w, h)

# Locate the x,y of start symbol @
#
def find_start_location(grid):
	index = grid.cells.index('@')
	return (index % grid.width, index / grid.width)

def find_teleporter_locations(grid):
	indices = [i for i, c in enumerate(grid.cells) if c == 'T']
	return map(lambda i: (i % grid.width, i / grid.width), indices)

# Check if the symbol will block movement based on the current player state
#
def is_blocker(symbol, breaker_mode):
	return symbol == '#' or (breaker_mode == False and symbol == 'X')

# A loop case occurs when entering a cell more than once with the same state
#
def is_looping(grid, loc, player_loop_state):
	return player_loop_state in grid.history[loc[1] * grid.width + loc[0]]

# Simulate taking a step (if possible) and identify the new location and contents
#
def simulate_step(grid, loc, direction_mod, breaker_mode):
	new_loc = direction_mod[1](loc)
	symbol = grid.cells[new_loc[1] * grid.width + new_loc[0]]
	if is_blocker(symbol, breaker_mode) == False:
		return new_loc, symbol

	#Blocked		
	return (None, symbol)

# Converts our shortened step into the output format
#
def format_output_step(step):
	mapper = {
		'N': 'NORTH',
		'E': 'EAST',
		'S': 'SOUTH',
		'W': 'WEST',
	}

	return mapper[step]

# Runs the simulation with the given input data
#
def run(grid, expected_output=None):
	debug_log(str.format("Grid is {}x{}", grid.width, grid.height))

	current_loc = find_start_location(grid)
	debug_log(str.format("Start: {}", current_loc))

	teleporter_locs = find_teleporter_locations(grid)
	debug_log(str.format("Teleporters: {}", teleporter_locs))

	# State
	steps = []
	next_step = None
	breaker_mode = False
	looping = False
	inverted = False
	num_broken_obstacles = 0
	# Ordered in normal priority order. Converts the location into the next location based on direction
	direction_mods = [('S', lambda l: (l[0], l[1] + 1)), ('E', lambda l: (l[0] + 1, l[1])), ('N', lambda l: (l[0], l[1] - 1)), ('W', lambda l: (l[0] - 1, l[1]))]

	for x in range(0, 1000): # Supposed to be a forever loop but adding a cap in case I mess up and get stuck

		# Prefer to keep going in the same direction if we can
		if next_step != None:
			direction_mod = next(ifilter(lambda m: m[0] == next_step, direction_mods), None)
			new_loc, symbol = simulate_step(grid, current_loc, direction_mod, breaker_mode)
			if new_loc == None:
				next_step = None

		# If we can't go in our preferred direction we pick a new direction based on the current direction priorities order
		if next_step == None:
			for direction_mod in direction_mods:
				new_loc, symbol = simulate_step(grid, current_loc, direction_mod, breaker_mode)
				if new_loc != None:
					next_step = direction_mod[0]
					break

		# Don't even think this is possible by the rules of the game but checking to make sure we don't get stuck
		if next_step == None:
			debug_log("Blocked")
			looping = True
			break

		# Take a step and log the direction we've travelled
		current_loc = new_loc
		debug_log("Step: " + next_step)
		steps.append(next_step)

		# Check if we are looping
		history_state = (breaker_mode, inverted, steps[-1], num_broken_obstacles)
		if is_looping(grid, current_loc, history_state) == True:
			debug_log("Looping")
			looping = True
			break

		# Make sure we update the visit history of the cell now before applying any mods as we check above without mods applied
		grid.history[current_loc[1] * grid.width + current_loc[0]].add(history_state)

		# Check if we are forced to change direction
		direction_mod = next(ifilter(lambda m: m[0] == symbol, direction_mods), None)
		if direction_mod != None:
			debug_log("Force Direction Change")
			next_step = direction_mod[0]
			continue

		# Check if we've reached the end
		elif symbol == '$':
			debug_log(str.format("Reached End: {}", current_loc))
			break

		# Check if we've found a teleporter and find the location of it's pair
		elif symbol == 'T':
			new_loc = filter(lambda t: t != current_loc, teleporter_locs)[0]
			debug_log(str.format("Teleporting from {} to: {}", current_loc, new_loc))
			current_loc = new_loc

		# Check if we've found a beer which toggles breaker mode on and off
		elif symbol == 'B':
			breaker_mode = not breaker_mode
			debug_log(str.format("Breaker Mode: {}", breaker_mode))

		# Check if we've found a breakable obstacle
		elif breaker_mode == True and symbol == 'X':
			debug_log(str.format("Smash obstacle at {}", current_loc))
			grid.cells[current_loc[1] * grid.width + current_loc[0]] = ' '
			num_broken_obstacles = num_broken_obstacles + 1

		# Check if we should invert the direction priorities
		elif symbol == 'I':
			debug_log("Invert")
			inverted = not inverted
			direction_mods.reverse()

	# The grid deals in single letter steps but they want the output as words
	formatted_steps = ['LOOP'] if looping == True else map(format_output_step, steps)

	if expected_output != None:
		debug_log("Finished: " + "SUCCESS" if formatted_steps == expected_output else "FAILED")

	# Write the answer out
	for s in formatted_steps:
		print(s)

# Used when running in CodinGame to read input from stdinput
#
if __name__ == "__main__":
	h, w = parse_dims(raw_input())
	unparsed_grid_lines = []
	for l in range(0, h):
		unparsed_grid_lines.append(raw_input())
	grid = parse_grid(unparsed_grid_lines, w, h)
	run(grid)