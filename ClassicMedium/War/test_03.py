# Test 03 - 26 cards per player - no wars

import core

INPUT_LINES = [
'26',
'6H',
'7H',
'6C',
'QS',
'7S',
'8D',
'6D',
'5S',
'6S',
'QH',
'4D',
'3S',
'7C',
'3C',
'4S',
'5H',
'QD',
'5C',
'3H',
'3D',
'8C',
'4H',
'4C',
'QC',
'5D',
'7D',
'26',
'JH',
'AH',
'KD',
'AD',
'9C',
'2D',
'2H',
'JC',
'10C',
'KC',
'10D',
'JS',
'JD',
'9D',
'9S',
'KS',
'AS',
'KH',
'10S',
'8S',
'2S',
'10H',
'8H',
'AC',
'2C',
'9H'
]

EXPECTED_OUTPUT = "2 56"

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