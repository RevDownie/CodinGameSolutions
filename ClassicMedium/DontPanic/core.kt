// OBJECTIVE
// The NPCs walk in single file and in a given direction. For each floor find the point at which a blocker should be placed (if at all)
// to direct the NPCs to the lift or exit. Make sure at least one NPC reaches the exit

// INPUT - Game
// Number of floors - Int
// Width of floor in cells - Int
// Number of rounds/loops - Int
// Index of floor containing exit (Y) - Int
// Index of cell containing exit (X) - Int
// Num. NPCs - Int
// Ignore - Int
// Num lifts - NL - Int
// NL lines containing Y, X of lift

// INPUT - Loop
// Floor of leading NPC (Y) - Int
// Cell of leading NPC (X) - Int
// Direction of leading NPC (LEFT, RIGHT) - String

// OUTPUT - Loop
// A single string each loop with either
//	- WAIT - Do nothing
//	- BLOCK - Cause the leading NPC to stop and act as a direction changer when hit by other NPCs

// NOTES
// 1 lift per floor
// Don't block in front of a lift as it will be blocked off
// Going off the edge of the floor is bad

// TO CONSIDER
// We only really care about the leading NPC at any time
// If they are heading in a direction away from the lift and towards certain death then block. Otherwise do nothing

import java.util.*
import java.io.*

data class NPCState(val floor : Int, val pos : Int, val dir : String)

/// For the given floor attempt to find the direction from the given 
/// location to the exit on that floor. Will return null if no exit found on the floor
///
fun tryGetDirectionToExit(npcState : NPCState, exitFloor : Int, exitPos : Int) : String? {
	return when {
		npcState.floor != exitFloor -> null
		exitPos > npcState.pos -> "RIGHT"
		exitPos < npcState.pos -> "LEFT"
		else -> npcState.dir
	} 
}

/// For the given floor attempt to find the direction from the given 
/// location to the lift on that floor. Will return null if no exit found on the floor
///
fun tryGetDirectionToLift(npcState : NPCState, liftFloorPositions : IntArray) : String? {
	val liftPos = liftFloorPositions[npcState.floor]
	return when {
		liftPos == -1 -> null
		liftPos > npcState.pos -> "RIGHT"
		liftPos < npcState.pos -> "LEFT"
		else -> npcState.dir
	}
}

/// Read data from CodinGame via stdin and output answers via stout. The strategy is
/// that we block if we are heading away from the "goal" on the current floor 
///
fun main(args : Array<String>) {
	val input = Scanner(System.`in`)

	val numFloors = input.nextInt()

	// Holds the pos of the lift for each floor. -1 if no lift
	val liftFloorPositions = IntArray(numFloors, {-1})

	val width = input.nextInt()
	val numLoops = input.nextInt()
	val exitFloor = input.nextInt()
	val exitPos = input.nextInt()
	val numNPCs = input.nextInt()

	val numAdditionalLifts = input.nextInt() // Ignore in this episode
	val numLifts = input.nextInt() 
	for (i in 0 until numLifts) {
		val liftFloor = input.nextInt()
		val liftPos = input.nextInt()
		liftFloorPositions[liftFloor] = liftPos;
	}

	var prevBlockedFloor = -1

	for (i in 0 until numLoops) {
		// The input data for each loop corresponds to the current leading NPC
		val floor = input.nextInt()
		val pos = input.nextInt()
		val direction = input.next()

		// Floor will be -1 if there are no spawned NPC
		if (floor >= 0) {
			val npcState = NPCState(floor, pos, direction)
			// Only one block required per floor. We always move up floors
			if (prevBlockedFloor < floor) {
				val dirToExit = tryGetDirectionToExit(npcState, exitFloor, exitPos)
				val dirToLift = tryGetDirectionToLift(npcState, liftFloorPositions)
				//Check if we are heading towards one of our mutually exclusive goals
				val shouldBlock = dirToExit != npcState.dir || dirToLift != npcState.dir
				if (shouldBlock == true) {
					// If not then block
					prevBlockedFloor = floor
					println("BLOCK")
					continue
				}				
			}
		}

		//Nothing to be done - either no NPCs or we are heading in the right direction
		println("WAIT")
	}
}