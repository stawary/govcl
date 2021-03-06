
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TXButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewXButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewXButton(owner IComponent) *TXButton {
    x := new(TXButton)
    x.instance = XButton_Create(CheckPtr(owner))
    x.ptr = unsafe.Pointer(x.instance)
    return x
}

// AsXButton
// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsXButton(obj interface{}) *TXButton {
    x := new(TXButton)
    x.instance, x.ptr = getInstance(obj)
    return x
}

// -------------------------- Deprecated begin --------------------------
// XButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsXButton.
func XButtonFromInst(inst uintptr) *TXButton {
    return AsXButton(inst)
}

// XButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsXButton.
func XButtonFromObj(obj IObject) *TXButton {
    return AsXButton(obj)
}

// XButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsXButton.
func XButtonFromUnsafePointer(ptr unsafe.Pointer) *TXButton {
    return AsXButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// Free 
// CN: 释放对象。
// EN: Free object.
func (x *TXButton) Free() {
    if x.instance != 0 {
        XButton_Free(x.instance)
        x.instance = 0
        x.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (x *TXButton) Instance() uintptr {
    return x.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (x *TXButton) UnsafeAddr() unsafe.Pointer {
    return x.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (x *TXButton) IsValid() bool {
    return x.instance != 0
}

// Is 
// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (x *TXButton) Is() TIs {
    return TIs(x.instance)
}

// As 
// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (x *TXButton) As() TAs {
//    return TAs(x.instance)
//}

// TXButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TXButtonClass() TClass {
    return XButton_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (x *TXButton) BringToFront() {
    XButton_BringToFront(x.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (x *TXButton) ClientToScreen(Point TPoint) TPoint {
    return XButton_ClientToScreen(x.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (x *TXButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return XButton_ClientToParent(x.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (x *TXButton) Dragging() bool {
    return XButton_Dragging(x.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (x *TXButton) HasParent() bool {
    return XButton_HasParent(x.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (x *TXButton) Hide() {
    XButton_Hide(x.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (x *TXButton) Invalidate() {
    XButton_Invalidate(x.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (x *TXButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return XButton_Perform(x.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (x *TXButton) Refresh() {
    XButton_Refresh(x.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (x *TXButton) Repaint() {
    XButton_Repaint(x.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (x *TXButton) ScreenToClient(Point TPoint) TPoint {
    return XButton_ScreenToClient(x.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (x *TXButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return XButton_ParentToClient(x.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (x *TXButton) SendToBack() {
    XButton_SendToBack(x.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (x *TXButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    XButton_SetBounds(x.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (x *TXButton) Show() {
    XButton_Show(x.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (x *TXButton) Update() {
    XButton_Update(x.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (x *TXButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return XButton_GetTextBuf(x.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (x *TXButton) GetTextLen() int32 {
    return XButton_GetTextLen(x.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (x *TXButton) SetTextBuf(Buffer string) {
    XButton_SetTextBuf(x.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (x *TXButton) FindComponent(AName string) *TComponent {
    return AsComponent(XButton_FindComponent(x.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (x *TXButton) GetNamePath() string {
    return XButton_GetNamePath(x.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (x *TXButton) Assign(Source IObject) {
    XButton_Assign(x.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (x *TXButton) DisposeOf() {
    XButton_DisposeOf(x.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (x *TXButton) ClassType() TClass {
    return XButton_ClassType(x.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (x *TXButton) ClassName() string {
    return XButton_ClassName(x.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (x *TXButton) InstanceSize() int32 {
    return XButton_InstanceSize(x.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (x *TXButton) InheritsFrom(AClass TClass) bool {
    return XButton_InheritsFrom(x.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (x *TXButton) Equals(Obj IObject) bool {
    return XButton_Equals(x.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (x *TXButton) GetHashCode() int32 {
    return XButton_GetHashCode(x.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (x *TXButton) ToString() string {
    return XButton_ToString(x.instance)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (x *TXButton) Caption() string {
    return XButton_GetCaption(x.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (x *TXButton) SetCaption(value string) {
    XButton_SetCaption(x.instance, value)
}

// ShowCaption
func (x *TXButton) ShowCaption() bool {
    return XButton_GetShowCaption(x.instance)
}

// SetShowCaption
func (x *TXButton) SetShowCaption(value bool) {
    XButton_SetShowCaption(x.instance, value)
}

// BackColor
func (x *TXButton) BackColor() TColor {
    return XButton_GetBackColor(x.instance)
}

// SetBackColor
func (x *TXButton) SetBackColor(value TColor) {
    XButton_SetBackColor(x.instance, value)
}

// HoverColor
func (x *TXButton) HoverColor() TColor {
    return XButton_GetHoverColor(x.instance)
}

// SetHoverColor
func (x *TXButton) SetHoverColor(value TColor) {
    XButton_SetHoverColor(x.instance, value)
}

// DownColor
func (x *TXButton) DownColor() TColor {
    return XButton_GetDownColor(x.instance)
}

// SetDownColor
func (x *TXButton) SetDownColor(value TColor) {
    XButton_SetDownColor(x.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (x *TXButton) BorderWidth() int32 {
    return XButton_GetBorderWidth(x.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (x *TXButton) SetBorderWidth(value int32) {
    XButton_SetBorderWidth(x.instance, value)
}

// BorderColor
func (x *TXButton) BorderColor() TColor {
    return XButton_GetBorderColor(x.instance)
}

// SetBorderColor
func (x *TXButton) SetBorderColor(value TColor) {
    XButton_SetBorderColor(x.instance, value)
}

// Picture
func (x *TXButton) Picture() *TPicture {
    return AsPicture(XButton_GetPicture(x.instance))
}

// SetPicture
func (x *TXButton) SetPicture(value *TPicture) {
    XButton_SetPicture(x.instance, CheckPtr(value))
}

// DrawMode
func (x *TXButton) DrawMode() TDrawImageMode {
    return XButton_GetDrawMode(x.instance)
}

// SetDrawMode
func (x *TXButton) SetDrawMode(value TDrawImageMode) {
    XButton_SetDrawMode(x.instance, value)
}

// NormalFontColor
func (x *TXButton) NormalFontColor() TColor {
    return XButton_GetNormalFontColor(x.instance)
}

// SetNormalFontColor
func (x *TXButton) SetNormalFontColor(value TColor) {
    XButton_SetNormalFontColor(x.instance, value)
}

// DownFontColor
func (x *TXButton) DownFontColor() TColor {
    return XButton_GetDownFontColor(x.instance)
}

// SetDownFontColor
func (x *TXButton) SetDownFontColor(value TColor) {
    XButton_SetDownFontColor(x.instance, value)
}

// HoverFontColor
func (x *TXButton) HoverFontColor() TColor {
    return XButton_GetHoverFontColor(x.instance)
}

// SetHoverFontColor
func (x *TXButton) SetHoverFontColor(value TColor) {
    XButton_SetHoverFontColor(x.instance, value)
}

// Action
func (x *TXButton) Action() *TAction {
    return AsAction(XButton_GetAction(x.instance))
}

// SetAction
func (x *TXButton) SetAction(value IComponent) {
    XButton_SetAction(x.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (x *TXButton) Align() TAlign {
    return XButton_GetAlign(x.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (x *TXButton) SetAlign(value TAlign) {
    XButton_SetAlign(x.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (x *TXButton) Anchors() TAnchors {
    return XButton_GetAnchors(x.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (x *TXButton) SetAnchors(value TAnchors) {
    XButton_SetAnchors(x.instance, value)
}

// BiDiMode
func (x *TXButton) BiDiMode() TBiDiMode {
    return XButton_GetBiDiMode(x.instance)
}

// SetBiDiMode
func (x *TXButton) SetBiDiMode(value TBiDiMode) {
    XButton_SetBiDiMode(x.instance, value)
}

// Constraints
func (x *TXButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(XButton_GetConstraints(x.instance))
}

// SetConstraints
func (x *TXButton) SetConstraints(value *TSizeConstraints) {
    XButton_SetConstraints(x.instance, CheckPtr(value))
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (x *TXButton) Enabled() bool {
    return XButton_GetEnabled(x.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (x *TXButton) SetEnabled(value bool) {
    XButton_SetEnabled(x.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (x *TXButton) Font() *TFont {
    return AsFont(XButton_GetFont(x.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (x *TXButton) SetFont(value *TFont) {
    XButton_SetFont(x.instance, CheckPtr(value))
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (x *TXButton) ParentFont() bool {
    return XButton_GetParentFont(x.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (x *TXButton) SetParentFont(value bool) {
    XButton_SetParentFont(x.instance, value)
}

// ParentShowHint
func (x *TXButton) ParentShowHint() bool {
    return XButton_GetParentShowHint(x.instance)
}

// SetParentShowHint
func (x *TXButton) SetParentShowHint(value bool) {
    XButton_SetParentShowHint(x.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (x *TXButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(XButton_GetPopupMenu(x.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (x *TXButton) SetPopupMenu(value IComponent) {
    XButton_SetPopupMenu(x.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (x *TXButton) ShowHint() bool {
    return XButton_GetShowHint(x.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (x *TXButton) SetShowHint(value bool) {
    XButton_SetShowHint(x.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (x *TXButton) Visible() bool {
    return XButton_GetVisible(x.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (x *TXButton) SetVisible(value bool) {
    XButton_SetVisible(x.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (x *TXButton) SetOnClick(fn TNotifyEvent) {
    XButton_SetOnClick(x.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (x *TXButton) SetOnDblClick(fn TNotifyEvent) {
    XButton_SetOnDblClick(x.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (x *TXButton) SetOnMouseDown(fn TMouseEvent) {
    XButton_SetOnMouseDown(x.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (x *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
    XButton_SetOnMouseEnter(x.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (x *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
    XButton_SetOnMouseLeave(x.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (x *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
    XButton_SetOnMouseMove(x.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (x *TXButton) SetOnMouseUp(fn TMouseEvent) {
    XButton_SetOnMouseUp(x.instance, fn)
}

// BoundsRect
func (x *TXButton) BoundsRect() TRect {
    return XButton_GetBoundsRect(x.instance)
}

// SetBoundsRect
func (x *TXButton) SetBoundsRect(value TRect) {
    XButton_SetBoundsRect(x.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (x *TXButton) ClientHeight() int32 {
    return XButton_GetClientHeight(x.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (x *TXButton) SetClientHeight(value int32) {
    XButton_SetClientHeight(x.instance, value)
}

// ClientOrigin
func (x *TXButton) ClientOrigin() TPoint {
    return XButton_GetClientOrigin(x.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (x *TXButton) ClientRect() TRect {
    return XButton_GetClientRect(x.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (x *TXButton) ClientWidth() int32 {
    return XButton_GetClientWidth(x.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (x *TXButton) SetClientWidth(value int32) {
    XButton_SetClientWidth(x.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (x *TXButton) ControlState() TControlState {
    return XButton_GetControlState(x.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (x *TXButton) SetControlState(value TControlState) {
    XButton_SetControlState(x.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (x *TXButton) ControlStyle() TControlStyle {
    return XButton_GetControlStyle(x.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (x *TXButton) SetControlStyle(value TControlStyle) {
    XButton_SetControlStyle(x.instance, value)
}

// ExplicitLeft
func (x *TXButton) ExplicitLeft() int32 {
    return XButton_GetExplicitLeft(x.instance)
}

// ExplicitTop
func (x *TXButton) ExplicitTop() int32 {
    return XButton_GetExplicitTop(x.instance)
}

// ExplicitWidth
func (x *TXButton) ExplicitWidth() int32 {
    return XButton_GetExplicitWidth(x.instance)
}

// ExplicitHeight
func (x *TXButton) ExplicitHeight() int32 {
    return XButton_GetExplicitHeight(x.instance)
}

// Floating
func (x *TXButton) Floating() bool {
    return XButton_GetFloating(x.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (x *TXButton) Parent() *TWinControl {
    return AsWinControl(XButton_GetParent(x.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (x *TXButton) SetParent(value IWinControl) {
    XButton_SetParent(x.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (x *TXButton) StyleElements() TStyleElements {
    return XButton_GetStyleElements(x.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (x *TXButton) SetStyleElements(value TStyleElements) {
    XButton_SetStyleElements(x.instance, value)
}

// SetOnGesture
func (x *TXButton) SetOnGesture(fn TGestureEvent) {
    XButton_SetOnGesture(x.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (x *TXButton) AlignWithMargins() bool {
    return XButton_GetAlignWithMargins(x.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (x *TXButton) SetAlignWithMargins(value bool) {
    XButton_SetAlignWithMargins(x.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (x *TXButton) Left() int32 {
    return XButton_GetLeft(x.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (x *TXButton) SetLeft(value int32) {
    XButton_SetLeft(x.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (x *TXButton) Top() int32 {
    return XButton_GetTop(x.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (x *TXButton) SetTop(value int32) {
    XButton_SetTop(x.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (x *TXButton) Width() int32 {
    return XButton_GetWidth(x.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (x *TXButton) SetWidth(value int32) {
    XButton_SetWidth(x.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (x *TXButton) Height() int32 {
    return XButton_GetHeight(x.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (x *TXButton) SetHeight(value int32) {
    XButton_SetHeight(x.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (x *TXButton) Cursor() TCursor {
    return XButton_GetCursor(x.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (x *TXButton) SetCursor(value TCursor) {
    XButton_SetCursor(x.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (x *TXButton) Hint() string {
    return XButton_GetHint(x.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (x *TXButton) SetHint(value string) {
    XButton_SetHint(x.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (x *TXButton) Margins() *TMargins {
    return AsMargins(XButton_GetMargins(x.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (x *TXButton) SetMargins(value *TMargins) {
    XButton_SetMargins(x.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (x *TXButton) CustomHint() *TCustomHint {
    return AsCustomHint(XButton_GetCustomHint(x.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (x *TXButton) SetCustomHint(value IComponent) {
    XButton_SetCustomHint(x.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (x *TXButton) ComponentCount() int32 {
    return XButton_GetComponentCount(x.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (x *TXButton) ComponentIndex() int32 {
    return XButton_GetComponentIndex(x.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (x *TXButton) SetComponentIndex(value int32) {
    XButton_SetComponentIndex(x.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (x *TXButton) Owner() *TComponent {
    return AsComponent(XButton_GetOwner(x.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (x *TXButton) Name() string {
    return XButton_GetName(x.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (x *TXButton) SetName(value string) {
    XButton_SetName(x.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (x *TXButton) Tag() int {
    return XButton_GetTag(x.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (x *TXButton) SetTag(value int) {
    XButton_SetTag(x.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (x *TXButton) Components(AIndex int32) *TComponent {
    return AsComponent(XButton_GetComponents(x.instance, AIndex))
}

