import core

def test_01():
	core.debug_log("Running Test 01")
	budgets = [3, 100, 100]
	c = 100
	expected_output = [3, 48, 49]
	core.run(c, budgets, expected_output)

def test_02():
	core.debug_log("Running Test 02")
	budgets = [40, 40, 40]
	c = 100
	expected_output = [33, 33, 34]
	core.run(c, budgets, expected_output)

def test_03():
	core.debug_log("Running Test 03")
	budgets = [100, 1, 60]
	c = 100
	expected_output = [1, 49, 50]
	core.run(c, budgets, expected_output)

def test_04():
	core.debug_log("Running Test 04")
	budgets = [20, 20, 20]
	c = 100
	expected_output = "IMPOSSIBLE"
	core.run(c, budgets, expected_output)

def test_05():
	core.debug_log("Running Test 05")
	budgets = [3, 3, 3]
	c = 3
	expected_output = [1, 1, 1]
	core.run(c, budgets, expected_output)

def test_06():
	core.debug_log("Running Test 06")
	budgets = [10, 100, 100]
	c = 100
	expected_output = [10, 45, 45]
	core.run(c, budgets, expected_output)

def test_07():
	core.debug_log("Running Test 07")
	budgets = [5, 10, 5]
	c = 10
	expected_output = [3, 3, 4]
	core.run(c, budgets, expected_output)

if __name__ == "__main__":
	test_01()
	test_02()
	test_03()
	test_04()
	test_05()
	test_06()
	test_07()

