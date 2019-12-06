package day6

func Count(m [][]string) int {
	count := 0

	orbits := buildOrbitTree(m)

	for name := range orbits {
		cursor := name
		for cursor != "COM" {
			cursor = orbits[cursor].parent
			count++
		}
	}

	return count
}

func TransfersBetween(m [][]string, org string, dest string) int {
	jumps := 0
	orbits := buildOrbitTree(m)

	orgPath := listOrbitsToCom(orbits, org)
	destPath := listOrbitsToCom(orbits, dest)

	cursor := orgPath[org].parent
	for {
		if _, ok := destPath[cursor]; ok {
			break
		}
		jumps++

		cursor = orgPath[cursor].parent
	}

	destCur := destPath[dest].parent
	for {
		if destCur == cursor {
			break
		}

		jumps++
		destCur = destPath[destCur].parent
	}

	return jumps
}

type orbit struct {
	parent string
}

func listOrbitsToCom(o map[string]orbit, org string) map[string]orbit {
	l := make(map[string]orbit)

	cursor := org
	for cursor != "COM" {
		l[cursor] = o[cursor]
		cursor = o[cursor].parent
	}

	return l
}

func buildOrbitTree(m [][]string) map[string]orbit {
	orb := make(map[string]orbit)

	for _, o := range m {
		orb[o[1]] = orbit{parent: o[0]}
	}

	return orb
}
