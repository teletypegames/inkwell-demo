package lib

import p "git.teletypegames.org/games/inkwell"

var CatMorning = p.Dialogue{
	Name:  "cat_morning",
	Start: "start",
	Nodes: []p.DialogueNode{
		{
			Name: "start",
			Lines: []p.DialogueLine{
				{Speaker: "cat", Text: "Mióóóóóóóóu. Késő van."},
			},
			Choices: []p.DialogueChoice{
				{
					Text:    "Mindjárt megy a kávé.",
					Actions: []p.Action{p.SetFlag(FlagPromisedCoffee), p.EndDialogue()},
				},
				{
					Text:    "Te etted meg a babot?",
					Show:    p.Not(p.Flag(FlagCupboardOpen)),
					Actions: []p.Action{p.GotoNode("beans")},
				},
				{
					Text:    "Hagyj békén.",
					Actions: []p.Action{p.EndDialogue()},
				},
			},
		},
		{
			Name: "beans",
			Lines: []p.DialogueLine{
				{Speaker: "cat", Text: "A szekrényben van. Mindig ott szokott lenni."},
			},
			Choices: []p.DialogueChoice{
				{
					Text:    "Köszi.",
					Actions: []p.Action{p.EndDialogue()},
				},
			},
		},
	},
}
