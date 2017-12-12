package math

type Rotation struct {
	Pitch, Yaw, HeadYaw float32
}

func NewRotation(pitch, yaw, headYaw float32) Rotation {
	return Rotation{pitch, yaw, headYaw}
}

func (rot *Rotation) GetPitch() float32 {
	return rot.Pitch
}

func (rot *Rotation) SetPitch(v float32)  {
	rot.Pitch = v
}

func (rot *Rotation) GetYaw() float32 {
	return rot.Yaw
}

func (rot *Rotation) SetYaw(v float32)  {
	rot.Yaw = v
}

func (rot *Rotation) GetHeadYaw() float32 {
	return rot.HeadYaw
}

func (rot *Rotation) SetHeadYaw(v float32)  {
	rot.HeadYaw = v
}