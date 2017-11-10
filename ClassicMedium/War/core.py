# CARD GAME RULES
# 2 players
# Dealt variable number of cards at the beginning of the game
# Cards are dealt face down with the next card lying on top of the old card (LIFO - stack)
# FIGHT - Each round the player reveals the top card and the player with the higher card takes both cards and puts them to the bottom of their deck (all of p1 cards first then the p2 cards)
# WAR - If the cards are equal then both players place the next 3 cards in their deck face down and FIGHT using that sub-deck (wars can form a chain)
# Once a player wins a war all the cards won are added to their deck
# Game is over when one player has no cards remaining and the winner is the other player
# Aces are high
# Each fight counts as a game round
# If during a WAR a player runs out of cards then the game is considered a draw

# NOTES
# We don't have to handle an infinite game
# Number of cards per player does not exceed 1000

# INPUT
# N - number of cards for player 1
# The cards of player 1, 1 per line in the format "NumSuit" e.g. 10D
# M - number of cards for player 2
# The cards of player 2, 1 per line in the format "NumSuit" e.g. 10D

# OUTPUT
# 1 or 2 if a clear winner followed by a space and the number of rounds completed
# PAT if drawn game

# STATE
# Round number
# Players decks (and sub decks)

# TO CONSIDER
# The player has a main root deck and sub decks for each war, this forms a chain but we don't need a hierarchy.
# The suits are irrelevant so we should just convert to numeric values 

import sys

# Ordered lowest to highest
VALUE_ORDER = ['2','3','4','5','6','7','8','9','10','J','Q','K','A']

# Logging to stdout is reserved for returning the answer. To debug log
# we must log to stderr
#
def debug_log(msg):
	sys.stderr.write(str(msg) + '\n')

# Convert the deck in format ["AD","KH","2H"] to numerically
# comparable values (we keep the same offsets for easier debugging)
#
def convert_deck_to_values(deck):
	return map(lambda c: VALUE_ORDER.index(c[0:-1]) + 2, deck)

# Return the index of the player that still has cards if the other player has none
# 
def try_get_winner_index(decks):
	for p in range(0, len(decks)):
		if len(decks[p]) == 0:
			# The other player won
			return (p + 1) % 2

	return -1

# Runs the simulation with the given input data
#
def run(unparsed_decks, expected_output = None):
	decks = map(convert_deck_to_values, unparsed_decks)
	debug_log(str.format("STARTING DECKS: {}", decks))

	round_num = 0
	who_won_game = -1

	for x in range(0, 100): # Supposed to be a forever loop but adding a cap in case I mess up and get stuck

		# New round
		round_num = round_num + 1

		# Reveal (remove) the top card for each player and compare (Note: our top card is the last card in the array)
		fight_cards = [decks[0].pop(), decks[1].pop()]

		who_won_fight = 0 if fight_cards[0] > fight_cards[1] else 1 if fight_cards[1] > fight_cards[0] else -1

		if who_won_fight >= 0:
			# Winner takes the cards - order is player 1 then player 2
			for c in fight_cards:
				decks[who_won_fight].insert(0, c) 
			debug_log(str.format("DECKS: {}", decks))

		else:
			# War
			None

		debug_log(str.format("FIGHT: Round: {}, Cards: {} vs {}. Winner: {}", round_num, fight_cards[0], fight_cards[1], who_won_fight + 1))

		# Check for a winner (i.e. if a player has no cards left they lose)
		who_won_game = try_get_winner_index(decks)
		if who_won_game >= 0:
			debug_log(str.format("WINNER: {}", who_won_game + 1))
			break

	# Output answer
	answer = str.format("{} {}", who_won_game + 1, round_num)
	if expected_output != None:
		debug_log("Finished: " + "SUCCESS" if answer == expected_output else "FAILED")
	print(answer)



# Used when running in CodinGame to read input from stdinput
#
if __name__ == "__main__":

	decks_unparsed = [[],[]]
	for p in range(0, 2):
		num_cards = int(raw_input())

		for i in range(0, num_cards):
			decks_unparsed[p].append(raw_input())

		# This is a bit confusing. We consider the Nth element to be the top of the deck and the rules say cards are dealt face down on top
		# of the deck, which I interpreted as the first card dealt is on the bottom and the last one on the top, however the test cases
		# only work if this is reversed
		decks_unparsed[p].reverse()

	run(decks_unparsed)