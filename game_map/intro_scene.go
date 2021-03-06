package game_map

var IntroScene = []interface{}{
	BlackScreen("blackscreen"),
	Wait(1),
	KillState("blackscreen"),
	PlayBGSound("../sound/rain.mp3"),
	TitleCaptionScreen("title", "Chandragupta Maurya", 3),
	SubTitleCaptionScreen("subtitle", "A jRPG game in GO", 2),
	KillState("title"),
	KillState("subtitle"),
	Scene("map_player_house", true, nil),
	RunActionAddNPC("map_player_house", "sleeper", 14, 19, 3),
	RunActionAddNPC("map_player_house", "guard", 19, 23, 0),
	PlaySound("../sound/door_break.mp3", 1),
	MoveNPC("guard", "map_player_house", []string{
		"up", "up", "up", "left", "left", "left",
	}),
	Say("map_player_house", "guard", "You'r coming with me!!", 3),
	StopBGSound(),
	PlaySound("../sound/wagon.mp3", 4),
	BlackScreen("blackscreen"),
	Wait(3),
	KillState("blackscreen"),
	ReplaceScene("map_player_house", "map_jail", 31, 21, false, nil),
	Wait(1),
	Say("map_jail", "hero", "Where am I...", 2),
	Say("map_jail", "hero", "I should keep looking for ways out", 2),
	Wait(1),
	HandOffToMainStack("map_jail"),
}
