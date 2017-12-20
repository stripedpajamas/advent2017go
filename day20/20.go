package day20

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type Particle struct {
	position           map[string]int
	velocity           map[string]int
	acceleration       map[string]int
	distanceFromOrigin int
	alive              bool
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func printParticles(particles []*Particle) {
	for i, p := range particles {
		fmt.Printf("%d:\n\tPosition: %d,%d,%d\n\tVelocity: %d,%d,%d\n\tAcceleration: %d,%d,%d\n\tDFO: %d\n\n",
			i, p.position["x"], p.position["y"], p.position["z"], p.velocity["x"], p.velocity["y"], p.velocity["z"],
			p.acceleration["x"], p.acceleration["y"], p.acceleration["z"], p.distanceFromOrigin)
	}
}

func dfo(p *Particle) int {
	dfo := 0
	for _, p := range p.position {
		dfo += int(math.Abs(float64(p)))
	}
	return dfo
}

func parseInput(input [][]byte) []*Particle {
	output := make([]*Particle, len(input))
	for idx, line := range input {
		split := bytes.Split(line, []byte(", "))
		// [p=<-3787,-3683,3352>, v=<41,-25,-124>, a=<5,9,1>]
		splitP := bytes.Split(split[0], []byte(","))
		splitV := bytes.Split(split[1], []byte(","))
		splitA := bytes.Split(split[2], []byte(","))
		// [p=<xx, xx, xx>]
		px, err := strconv.Atoi(string(splitP[0][3:]))
		checkErr(err)
		py, err := strconv.Atoi(string(splitP[1]))
		checkErr(err)
		pz, err := strconv.Atoi(string(splitP[2][:len(splitP[2])-1]))
		checkErr(err)
		vx, err := strconv.Atoi(string(splitV[0][3:]))
		checkErr(err)
		vy, err := strconv.Atoi(string(splitV[1]))
		checkErr(err)
		vz, err := strconv.Atoi(string(splitV[2][:len(splitV[2])-1]))
		checkErr(err)
		ax, err := strconv.Atoi(string(splitA[0][3:]))
		checkErr(err)
		ay, err := strconv.Atoi(string(splitA[1]))
		checkErr(err)
		az, err := strconv.Atoi(string(splitA[2][:len(splitA[2])-1]))
		checkErr(err)

		output[idx] = &Particle{
			alive: true,
			position: map[string]int{
				"x": px,
				"y": py,
				"z": pz,
			},
			velocity: map[string]int{
				"x": vx,
				"y": vy,
				"z": vz,
			},
			acceleration: map[string]int{
				"x": ax,
				"y": ay,
				"z": az,
			},
		}

		output[idx].distanceFromOrigin = dfo(output[idx])
	}
	return output
}

func Solve1(particles []*Particle) int {
	for i := 0; i < 1000; i++ {
		// go through 1000 units of time and then measure
		for _, particle := range particles {
			particle.velocity["x"] += particle.acceleration["x"]
			particle.velocity["y"] += particle.acceleration["y"]
			particle.velocity["z"] += particle.acceleration["z"]
			particle.position["x"] += particle.velocity["x"]
			particle.position["y"] += particle.velocity["y"]
			particle.position["z"] += particle.velocity["z"]
			particle.distanceFromOrigin = dfo(particle)
		}
	}

	// measure
	closestIdx := 0
	distance := particles[closestIdx].distanceFromOrigin
	for i, p := range particles {
		if p.distanceFromOrigin < distance {
			closestIdx = i
			distance = p.distanceFromOrigin
		}
	}

	return closestIdx
}

func handleCollisions(particles []*Particle) {
	posMap := make(map[string]int)
	killMap := make(map[int]bool)
	// loop through particles and update posmap
	for idx, p := range particles {
		if p.alive {
			posId := fmt.Sprintf("%d,%d,%d", p.position["x"], p.position["y"], p.position["z"])
			if i, found := posMap[posId]; found {
				// posMap already had this position, kill both
				killMap[idx], killMap[i] = true, true
			} else {
				posMap[posId] = idx
			}
		}
	}

	for idx, _ := range killMap {
		particles[idx].alive = false
	}
}

func Solve2(particles []*Particle) int {
	for i := 0; i < 1000; i++ {
		handleCollisions(particles)
		for _, particle := range particles {
			if particle.alive {
				particle.velocity["x"] += particle.acceleration["x"]
				particle.velocity["y"] += particle.acceleration["y"]
				particle.velocity["z"] += particle.acceleration["z"]
				particle.position["x"] += particle.velocity["x"]
				particle.position["y"] += particle.velocity["y"]
				particle.position["z"] += particle.velocity["z"]
				particle.distanceFromOrigin = dfo(particle)
			}
		}
	}
	// count alive particles
	total := 0
	for _, p := range particles {
		if p.alive {
			total++
		}
	}
	return total
}
