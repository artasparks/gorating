package gorating

// Construct a map from a player id to all the games they played.
func PlayerMaps(games []Game) (map[string][]Game, map[string]PlayerRating) {
	m := make(map[string][]Game)
	pr := make(map[string]PlayerRating)
	for _, g := range games {
		for _, p := range []PlayerRating{g.PlayerOne(), g.PlayerTwo()} {
			if _, ok := m[p.UnqiueId()]; !ok {
				m[p.UnqiueId()] = make([]Game, 0, 5)
			}
			if _, ok := pr[p.UnqiueId()]; !ok {
				pr[p.UnqiueId()] = p
			}
			arr := m[p.UnqiueId()]
			m[p.UnqiueId()] = append(arr, g)
		}
	}
	return m, pr
}

func FilterGames(player Player, games []Game) []Game {
	m := make([]Game, 0, 5)
	for _, g := range games {
		if g.PlayerOne().UnqiueId() == player.UnqiueId() ||
			g.PlayerTwo().UnqiueId() == player.UnqiueId() {
			m = append(m, g)
		}
	}
	return m
}
