import core

def test_01():
	core.debug_log("Running Test 01")
	links = {0 :[1], 1:[0,2], 2:[1,3,4], 3:[2], 4:[2]}
	core.run(links, 2)

def test_02():
	core.debug_log("Running Test 02")
	links = {0 :[1], 1:[0,2,4], 2:[1,3], 3:[2], 4:[1,5,6], 5:[4], 6:[4]}
	core.run(links, 2)

if __name__ == "__main__":
	test_01()
	test_02()

