package lib

import p "git.teletypegames.org/engines/inkwell"

var CatMorning = p.Dialog{
	Name:  "cat_morning",
	Start: "start",
	Nodes: []p.DialogNode{
		{
			Name: "start",
			Lines: []p.DialogLine{
				{Speaker: "cat", Text: p.Text{{Lang: "hu", Text: "Mióóóóóóóóu. Késő van."}}},
			},
			Choices: []p.DialogChoice{
				{
					Text:    p.Text{{Lang: "hu", Text: "Mindjárt megy a kávé."}},
					Actions: []p.Action{p.SetFlag(FlagPromisedCoffee), p.EndDialog()},
				},
				{
					Text:    p.Text{{Lang: "hu", Text: "Te etted meg a babot?"}},
					Show:    p.Not(p.Flag(FlagCupboardOpen)),
					Actions: []p.Action{p.GotoNode("beans")},
				},
				{
					Text:    p.Text{{Lang: "hu", Text: "Hagyj békén."}},
					Actions: []p.Action{p.EndDialog()},
				},
			},
		},
		{
			Name: "beans",
			Lines: []p.DialogLine{
				{Speaker: "cat", Text: p.Text{{Lang: "hu", Text: "A szekrényben van. Mindig ott szokott lenni."}}},
			},
			Choices: []p.DialogChoice{
				{
					Text:    p.Text{{Lang: "hu", Text: "Köszi."}},
					Actions: []p.Action{p.EndDialog()},
				},
			},
		},
	},
}
