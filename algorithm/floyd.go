package algorithm

import "math"

type Graph struct {
	To int
	Wt float64
}

func FloydWarshall(g [][]Graph) [][]float64 {
	dist := make([][]float64, len(g))

	for i := range dist {
		di := make([]float64, len(g))
		for j := range di {
			di[j] = math.Inf(i)
		}

    di[i] = 0
    dist[i] = di
	}

  for u, graphs := range g {
    for _, v := range graphs {
      dist[u][v.To] = v.Wt
    }
  }

  for k, dk := range dist {
    for _, di := range dist {
      for j, dij := range di {
        if d := di[k] + dk[j]; dij > d {
          di[j] = d
        }
      }
    }
  }

  return dist
}
