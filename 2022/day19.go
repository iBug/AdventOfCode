package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("19-1", Solution19_1)
	RegisterSolution("19-2", Solution19_2)
}

type Blueprint19 struct {
	ore, clay, obsidian, obsidian2, geode, geode2 int
	maxOreCost                                    int
}

func (bp *Blueprint19) getMaxOreCost() int {
	bp.maxOreCost = Max(bp.ore, bp.clay, bp.obsidian, bp.geode)
	return bp.maxOreCost
}

func (bp Blueprint19) String() string {
	return fmt.Sprintf("[Ore: %d, Clay: %d, Obsidian: %d+%d, Geode: %d+%d]", bp.ore, bp.clay, bp.obsidian, bp.obsidian2, bp.geode, bp.geode2)
}

type Resource19 struct {
	ore, clay, obsidian, geode int
}

func (r Resource19) Add(o Resource19) Resource19 {
	return Resource19{
		ore:      r.ore + o.ore,
		clay:     r.clay + o.clay,
		obsidian: r.obsidian + o.obsidian,
		geode:    r.geode + o.geode,
	}
}

func (r Resource19) canBuildGeodeBot(b Blueprint19) bool {
	return r.ore >= b.geode && r.obsidian >= b.geode2
}

func TimeToNextOreBot19(bp Blueprint19, count Resource19, bots Resource19) int {
	if bots.ore >= bp.maxOreCost {
		return -1
	}
	oreNeeded := bp.ore - count.ore
	if oreNeeded <= 0 {
		return 0
	}
	return (oreNeeded + bots.ore - 1) / bots.ore
}

func TimeToNextClayBot19(bp Blueprint19, count Resource19, bots Resource19) int {
	if bots.clay >= bp.obsidian2 {
		return -1
	}
	oreNeeded := bp.clay - count.ore
	if oreNeeded <= 0 {
		return 0
	}
	return (oreNeeded + bots.ore - 1) / bots.ore
}

func TimeToNextObsidianBot19(bp Blueprint19, count Resource19, bots Resource19) int {
	if bots.clay == 0 {
		return -1
	}
	if bots.obsidian >= bp.geode2 {
		return -1
	}
	oreNeeded := bp.obsidian - count.ore
	if oreNeeded <= 0 {
		oreNeeded = 0
	}
	clayNeeded := bp.obsidian2 - count.clay
	if clayNeeded <= 0 {
		clayNeeded = 0
	}
	return Max((oreNeeded+bots.ore-1)/bots.ore, (clayNeeded+bots.clay-1)/bots.clay)
}

func TimeToNextGeodeBot19(bp Blueprint19, count Resource19, bots Resource19) int {
	if bots.obsidian == 0 {
		return -1
	}
	oreNeeded := bp.geode - count.ore
	if oreNeeded <= 0 {
		oreNeeded = 0
	}
	obsidianNeeded := bp.geode2 - count.obsidian
	if obsidianNeeded <= 0 {
		obsidianNeeded = 0
	}
	return Max((oreNeeded+bots.ore-1)/bots.ore, (obsidianNeeded+bots.obsidian-1)/bots.obsidian)
}

func Recurse19(minutes int, bp Blueprint19, count Resource19, bots Resource19) int {
	if minutes == 0 {
		return count.geode
	}
	if minutes == 1 {
		return count.geode + bots.geode
	}
	if minutes == 2 {
		result := count.geode + bots.geode*2
		if count.canBuildGeodeBot(bp) {
			result += 1
		}
		return result
	}

	maxGeodeCount := count.geode + bots.geode*minutes
	timeNeeded := 0

	// Next bot is geode bot
	timeNeeded = TimeToNextGeodeBot19(bp, count, bots)
	if timeNeeded >= 0 && timeNeeded < minutes {
		newCount := count
		for i := 0; i < timeNeeded; i++ {
			newCount = newCount.Add(bots)
		}
		newCount.ore -= bp.geode
		newCount.obsidian -= bp.geode2

		newResult := Recurse19(minutes-timeNeeded-1, bp, newCount.Add(bots), bots.Add(Resource19{0, 0, 0, 1}))
		if newResult > maxGeodeCount {
			maxGeodeCount = newResult
		}
	}

	// Next bot is obsidian bot
	timeNeeded = TimeToNextObsidianBot19(bp, count, bots)
	if timeNeeded >= 0 && timeNeeded < minutes {
		newCount := count
		for i := 0; i < timeNeeded; i++ {
			newCount = newCount.Add(bots)
		}
		newCount.ore -= bp.obsidian
		newCount.clay -= bp.obsidian2

		newResult := Recurse19(minutes-timeNeeded-1, bp, newCount.Add(bots), bots.Add(Resource19{0, 0, 1, 0}))
		if newResult > maxGeodeCount {
			maxGeodeCount = newResult
		}
	}

	// Next bot is clay bot
	timeNeeded = TimeToNextClayBot19(bp, count, bots)
	if timeNeeded >= 0 && timeNeeded < minutes {
		newCount := count
		for i := 0; i < timeNeeded; i++ {
			newCount = newCount.Add(bots)
		}
		newCount.ore -= bp.clay

		newResult := Recurse19(minutes-timeNeeded-1, bp, newCount.Add(bots), bots.Add(Resource19{0, 1, 0, 0}))
		if newResult > maxGeodeCount {
			maxGeodeCount = newResult
		}
	}

	// Next bot is ore bot
	timeNeeded = TimeToNextOreBot19(bp, count, bots)
	if timeNeeded >= 0 && timeNeeded < minutes {
		newCount := count
		for i := 0; i < timeNeeded; i++ {
			newCount = newCount.Add(bots)
		}
		newCount.ore -= bp.ore

		newResult := Recurse19(minutes-timeNeeded-1, bp, newCount.Add(bots), bots.Add(Resource19{1, 0, 0, 0}))
		if newResult > maxGeodeCount {
			maxGeodeCount = newResult
		}
	}

	return maxGeodeCount
}

func ParseBlueprint19(line string) (blueprintId int, bp Blueprint19) {
	f := strings.Fields(line)
	blueprintId, _ = strconv.Atoi(strings.TrimSuffix(f[1], ":"))
	bp.ore, _ = strconv.Atoi(f[6])
	bp.clay, _ = strconv.Atoi(f[12])
	bp.obsidian, _ = strconv.Atoi(f[18])
	bp.obsidian2, _ = strconv.Atoi(f[21])
	bp.geode, _ = strconv.Atoi(f[27])
	bp.geode2, _ = strconv.Atoi(f[30])
	bp.getMaxOreCost()
	return
}

func Solution19_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		blueprintId, bp := ParseBlueprint19(scanner.Text())
		result := Recurse19(24, bp, Resource19{0, 0, 0, 0}, Resource19{1, 0, 0, 0})
		// fmt.Printf("Blueprint %d %v: %d\n", blueprintId, bp, result)
		total += blueprintId * result
	}
	fmt.Println(total)
}

func Solution19_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 1
	for scanner.Scan() {
		blueprintId, bp := ParseBlueprint19(scanner.Text())
		result := Recurse19(32, bp, Resource19{0, 0, 0, 0}, Resource19{1, 0, 0, 0})
		// fmt.Printf("Blueprint %d %v: %d\n", blueprintId, bp, result)
		total *= result
		if blueprintId == 3 {
			break
		}
	}
	fmt.Println(total)
}
