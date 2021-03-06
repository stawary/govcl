
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function MonthCalColors_Create(AOwner: TCommonCalendar): TMonthCalColors; stdcall;
begin
  Result :=  TMonthCalColors.Create(AOwner);
end;

procedure MonthCalColors_Free(AObj: TMonthCalColors); stdcall;
begin
  AObj.Free;
end;

procedure MonthCalColors_Assign(AObj: TMonthCalColors; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function MonthCalColors_GetNamePath(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MonthCalColors_DisposeOf(AObj: TMonthCalColors); stdcall;
begin
  AObj.DisposeOf;
end;

function MonthCalColors_ClassType(AObj: TMonthCalColors): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function MonthCalColors_ClassName(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MonthCalColors_InstanceSize(AObj: TMonthCalColors): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function MonthCalColors_InheritsFrom(AObj: TMonthCalColors; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function MonthCalColors_Equals(AObj: TMonthCalColors; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MonthCalColors_GetHashCode(AObj: TMonthCalColors): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MonthCalColors_ToString(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MonthCalColors_GetBackColor(AObj: TMonthCalColors): TColor; stdcall;
begin
  Result :=  AObj.BackColor;
end;

procedure MonthCalColors_SetBackColor(AObj: TMonthCalColors; AValue: TColor); stdcall;
begin
  AObj.BackColor := AValue;
end;

function MonthCalColors_StaticClassType: TClass; stdcall;
begin
  Result := TMonthCalColors;
end;

exports
  MonthCalColors_Create,
  MonthCalColors_Free,
  MonthCalColors_Assign,
  MonthCalColors_GetNamePath,
  MonthCalColors_DisposeOf,
  MonthCalColors_ClassType,
  MonthCalColors_ClassName,
  MonthCalColors_InstanceSize,
  MonthCalColors_InheritsFrom,
  MonthCalColors_Equals,
  MonthCalColors_GetHashCode,
  MonthCalColors_ToString,
  MonthCalColors_GetBackColor,
  MonthCalColors_SetBackColor,
  MonthCalColors_StaticClassType;

