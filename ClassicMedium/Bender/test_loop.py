# Test in which Bender cannot reach the goal and instead loops

import core

SAMPLE_INPUT_DIMS = '15 15'
SAMPLE_INPUT_GRID_LINES = [
'###############',
'#      IXXXXX #',
'#  @          #',
'#E S          #',
'#             #',
'#  I          #',
'#  B          #',
'#  B   S     W#',
'#  B   T      #',
'#             #',
'#         T   #',
'#         B   #',
'#N          W$#',
'#        XXXX #',
'###############']

SAMPLE_EXPECTED_OUTPUT_LINES = [
'LOOP',
]

h, w = core.parse_dims(SAMPLE_INPUT_DIMS)
grid = core.parse_grid(SAMPLE_INPUT_GRID_LINES, w, h)
core.run(grid, SAMPLE_EXPECTED_OUTPUT_LINES)
