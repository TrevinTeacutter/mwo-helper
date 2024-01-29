package home

import (
	"strconv"
	"strings"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Scoreboard struct{}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{}
}

func (s *Scoreboard) Layout(gtx layout.Context, theme *material.Theme, score int, winner bool, users []api.UserDetails) layout.Dimensions {
	// var children []layout.StackChild
	//
	// for _, user := range details {
	// 	children = append(children, s.UserRow(theme, user))
	// }
	//
	// rows := []layout.FlexChild{
	// 	s.Headers(theme),
	// }
	//
	// if len(children) > 0 {
	// 	rows = append(
	// 		rows,
	// 		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
	// 			return layout.Flex{
	// 				Alignment: layout.Middle,
	// 				Axis:      layout.Horizontal,
	// 			}.Layout(
	// 				gtx,
	// 				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
	// 					label := material.Label(theme, unit.Sp(15), strconv.Itoa(score))
	//
	// 					label.Alignment = text.Middle
	//
	// 					return label.Layout(gtx)
	// 				}),
	// 				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
	// 					return layout.Stack{
	// 						Alignment: layout.Center,
	// 					}.Layout(gtx, children...)
	// 				}),
	// 			)
	// 		}),
	// 	)
	// }

	lances := make([]string, len(users))
	players := make([]string, len(users))
	mechs := make([]string, len(users))
	scores := make([]int, len(users))
	damage := make([]int, len(users))
	kills := make([]int, len(users))
	assists := make([]int, len(users))
	components := make([]int, len(users))
	teamDamage := make([]int, len(users))

	for index, player := range users {
		lances[index] = player.Lance
		players[index] = player.Username
		mechs[index] = player.MechName
		scores[index] = player.MatchScore
		damage[index] = player.Damage
		kills[index] = player.Kills
		assists[index] = player.Assists
		components[index] = player.ComponentsDestroyed
		teamDamage[index] = player.TeamDamage
	}

	columns := []layout.FlexChild{
		s.ScoreColumn(theme, score, len(users), winner),
		s.LanceColumn(theme, lances...),
		s.PlayerColumn(theme, players...),
		s.MechColumn(theme, mechs...),
		s.MatchScoreColumn(theme, scores...),
		s.DamageColumn(theme, damage...),
		s.KillsColumn(theme, kills...),
		s.AssistsColumn(theme, assists...),
		s.ComponentsColumn(theme, components...),
		s.TeamDamageColumn(theme, teamDamage...),
	}

	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(
		gtx,
		columns...,
	)
}

func (s *Scoreboard) ScoreColumn(theme *material.Theme, score int, playerCount int, winner bool) layout.FlexChild {
	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(
			gtx,
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				label := material.Label(theme, unit.Sp(15), "Score")

				label.MaxLines = 1
				label.Alignment = text.Middle

				return label.Layout(gtx)
			}),
			layout.Flexed(float32(1.0*playerCount), func(gtx layout.Context) layout.Dimensions {
				label := material.Label(theme, unit.Sp(15), strconv.Itoa(score))
				label.MaxLines = 1

				if winner {
					label.MaxLines = 2
					label.Text += "\n🏆"
				}

				label.Alignment = text.Middle

				return label.Layout(gtx)
			}),
		)
	})
}

func (s *Scoreboard) LanceColumn(theme *material.Theme, lances ...string) layout.FlexChild {
	children := make([]layout.FlexChild, 0, len(lances)+1)

	children = append(children,
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Lance").Layout(gtx)
		}),
	)

	for _, lance := range lances {
		value := lance
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), value).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) PlayerColumn(theme *material.Theme, players ...string) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Player").Layout(gtx)
		}),
	}

	for _, player := range players {
		value := player
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), value).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) MechColumn(theme *material.Theme, mechs ...string) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Mech").Layout(gtx)
		}),
	}

	for _, mech := range mechs {
		value := mech
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strings.ToUpper(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) HealthColumn(theme *material.Theme, healths ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Health").Layout(gtx)
		}),
	}

	for _, health := range healths {
		value := health
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), s.HealthText(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) MatchScoreColumn(theme *material.Theme, scores ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Match Score").Layout(gtx)
		}),
	}

	for _, score := range scores {
		value := score
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) DamageColumn(theme *material.Theme, damages ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Damage").Layout(gtx)
		}),
	}

	for _, damage := range damages {
		value := damage
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) KillsColumn(theme *material.Theme, kills ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Kills").Layout(gtx)
		}),
	}

	for _, kill := range kills {
		value := kill
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) AssistsColumn(theme *material.Theme, assists ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Assists").Layout(gtx)
		}),
	}

	for _, assist := range assists {
		value := assist
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) KMDDColumn(theme *material.Theme, kmdds ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "KMDD").Layout(gtx)
		}),
	}

	for _, kmdd := range kmdds {
		value := kmdd
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) ComponentsColumn(theme *material.Theme, components ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Components").Layout(gtx)
		}),
	}

	for _, component := range components {
		value := component
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) TeamDamageColumn(theme *material.Theme, damages ...int) layout.FlexChild {
	children := []layout.FlexChild{
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Team Damage").Layout(gtx)
		}),
	}

	for _, damage := range damages {
		value := damage
		children = append(children, layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), strconv.Itoa(value)).Layout(gtx)
		}))
	}

	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx, children...)
	})
}

func (s *Scoreboard) Headers(theme *material.Theme) layout.FlexChild {
	return layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Horizontal,
		}.Layout(
			gtx,
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Score").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Player").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Mech").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Health").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Match Score").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Damage").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Kills").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Assists").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "KMDD").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Components").Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), "Team Damage").Layout(gtx)
			}),
		)
	})
}

func (s *Scoreboard) UserRow(theme *material.Theme, details api.UserDetails) layout.StackChild {
	return layout.Stacked(func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Horizontal,
		}.Layout(
			gtx,
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), details.Username).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strings.ToUpper(details.MechName)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), s.HealthText(details.HealthPercentage)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.MatchScore)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.Damage)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.Kills)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.Assists)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.KillsMostDamage)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.ComponentsDestroyed)).Layout(gtx)
			}),
			layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), strconv.Itoa(details.TeamDamage)).Layout(gtx)
			}),
		)
	})
}

func (s *Scoreboard) HealthText(value int) string {
	if value == 0 {
		return "DEAD"
	}

	return strconv.Itoa(value) + "%"
}
