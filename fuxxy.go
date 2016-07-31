package fuxxy

import "sort"

type (
	match struct {
		i      int
		points int
	}

	byPointsDesc []match
)

func (a byPointsDesc) Len() int           { return len(a) }
func (a byPointsDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPointsDesc) Less(i, j int) bool { return a[i].points > a[j].points }

func Matches(names []string, value string) []string {

	if len(names) == 0 || len(value) == 0 {
		return nil
	}

	var (
		valueRunes = []rune(value)
		matches    = make([]match, 0, 8)
	)

	for i, name := range names {
		var j, pts, in int
		for _, r := range name {
			if r == valueRunes[j] {
				pts += in
				j, in = j+1, in+1
				if j == len(valueRunes) {
					matches = append(matches, match{i: i, points: pts})
					break
				}
			} else {
				in = 0
				pts--
			}
		}
	}

	sort.Stable(byPointsDesc(matches))

	ret := make([]string, len(matches))
	for i, match := range matches {
		ret[i] = names[match.i]
	}

	return ret

}
