package main

type MoveState struct {
	mCharacter  Character
	mMap        GameMap
	mEntity     *Entity
	mController *StateMachine
	// ^above common with WaitState
	mTileWidth       float64
	mMoveX, mMoveY   float64
	mPixelX, mPixelY float64
	mMoveSpeed       float64
}

func MoveStateCreate(character Character, gMap GameMap) State {
	s := &MoveState{}
	s.mCharacter = character
	s.mMap = gMap
	s.mTileWidth = gMap.mTileWidth
	s.mEntity = character.mEntity
	s.mController = character.mController
	s.mMoveX = 0
	s.mMoveY = 0
	//s.mTween = Tween:Create(0, 0, 1),
	s.mMoveSpeed = 0.3
	return s
}

//The StateMachine class requires each state to have
// four functions: Enter, Exit, Render and Update

func (s *MoveState) Enter(data Direction) {
	//save Move X,Y value to used inside Update call
	s.mMoveX = data.x
	s.mMoveY = data.y
	s.mPixelX = s.mEntity.mTileX
	s.mPixelY = s.mEntity.mTileY
	//s.mTween = Tween:Create(0, self.mTileWidth, self.mMoveSpeed)

	//we update Hero Tile position before,
	// but execute only during exit phase
	s.mEntity.mTileX += data.x
	s.mEntity.mTileY += data.y
}

func (s *MoveState) Exit() {
	//Teleport(s.mEntity, s.mMap)
	s.mEntity.TeleportAndDraw(s.mMap)
}

func (s *MoveState) Render() {
	//pending
}

func (s *MoveState) Update(dt float64) {
	//self.mTween:Update(dt)
	//value := s.mTween:Value()
	x := s.mPixelX + s.mMoveX
	y := s.mPixelY + s.mMoveY
	s.mEntity.mTileX = x
	s.mEntity.mTileY = y
	//if s.mTween:IsFinished() then
	//	s.mController:Change("wait")
	//end
	s.mController.Change("wait", Direction{0, 0})
}