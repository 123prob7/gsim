package eula

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/core/player/character/profile"
)

func init() {
	core.RegisterCharFunc(keys.Eula, NewChar)
}

type char struct {
	*tmpl.Character
	burstCounter    int
	burstCounterICD int
	grimheartStacks int
	c1buff          []float64
	particleDone    bool
}

func NewChar(s *core.Core, w *character.CharWrapper, _ profile.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 80
	c.NormalHitNum = normalHitNum
	c.BurstCon = 3
	c.SkillCon = 5

	w.Character = &c

	return nil
}

func (c *char) Init() error {
	if c.Base.Cons >= 1 {
		c.c1buff = make([]float64, attributes.EndStatType)
		c.c1buff[attributes.PhyP] = 0.3
	}
	c.burstStacks()
	c.onExitField()
	if c.Base.Cons >= 4 {
		c.c4()
	}
	return nil
}
