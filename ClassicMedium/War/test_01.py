# Test 01 - 3 cards per player - no wars

import core

INPUT_LINES = [
'3',
'AD',
'KC',
'QC',
'3',
'KH',
'QS',
'JC'
]

EXPECTED_OUTPUT = "1 3"

line_idx = 0
num_cards = int(INPUT_LINES[line_idx])
line_idx = line_idx + 1

deck_unparsed_1 = INPUT_LINES[line_idx:line_idx+num_cards]
line_idx = line_idx + num_cards

num_cards = int(INPUT_LINES[line_idx])
line_idx = line_idx + 1
deck_unparsed_2 = INPUT_LINES[line_idx:line_idx+num_cards]

core.run([deck_unparsed_1, deck_unparsed_2], EXPECTED_OUTPUT)