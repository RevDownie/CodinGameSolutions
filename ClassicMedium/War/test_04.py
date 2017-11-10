# Test 04 - Wars

import core

INPUT_LINES = [
'5',
'8C',
'KD',
'AH',
'QH',
'2S',
'5',
'8D',
'2D',
'3H',
'4D',
'3S'
]

EXPECTED_OUTPUT = "2 1"

line_idx = 0
num_cards = int(INPUT_LINES[line_idx])
line_idx = line_idx + 1

deck_unparsed_1 = INPUT_LINES[line_idx:line_idx+num_cards]
deck_unparsed_1.reverse()
line_idx = line_idx + num_cards

num_cards = int(INPUT_LINES[line_idx])
line_idx = line_idx + 1
deck_unparsed_2 = INPUT_LINES[line_idx:line_idx+num_cards]
deck_unparsed_2.reverse()

core.run([deck_unparsed_1, deck_unparsed_2], EXPECTED_OUTPUT)