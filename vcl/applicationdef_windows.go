
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl

import (
	"github.com/ying32/govcl/vcl/win"
)

// SetIconResId 从资源中设置图标的id
func (a *TApplication) SetIconResId(id int) {
	hicon := win.LoadIcon(win.GetSelfModuleHandle(), id)
	if hicon != 0 {
		a.Icon().SetHandle(hicon)
	}
}
