package deepcopymacro

import (
	"github.com/tdakkota/deepcopymacro/internal"
	"github.com/tdakkota/gomacro/derive"
	"go/types"
)

type Macro struct {
	deepCopy *internal.DeepCopy
}

func (m Macro) Protocol() derive.Protocol {
	return m.deepCopy
}

func (m Macro) Name() string {
	return "derive_deepcopy"
}

func (m Macro) Target() *types.Interface {
	return m.deepCopy.Target()
}
