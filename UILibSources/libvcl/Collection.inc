
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function Collection_Create(AOwner: TCollectionItemClass): TCollection; stdcall;
begin
  Result :=  TCollection.Create(AOwner);
end;

procedure Collection_Free(AObj: TCollection); stdcall;
begin
  AObj.Free;
end;

function Collection_Owner(AObj: TCollection): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Collection_Add(AObj: TCollection): TCollectionItem; stdcall;
begin
  Result :=  AObj.Add;
end;

procedure Collection_Assign(AObj: TCollection; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Collection_BeginUpdate(AObj: TCollection); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure Collection_Clear(AObj: TCollection); stdcall;
begin
  AObj.Clear;
end;

procedure Collection_ClearAndResetID(AObj: TCollection); stdcall;
begin
  AObj.ClearAndResetID;
end;

procedure Collection_Delete(AObj: TCollection; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure Collection_EndUpdate(AObj: TCollection); stdcall;
begin
  AObj.EndUpdate;
end;

function Collection_FindItemID(AObj: TCollection; ID: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.FindItemID(ID);
end;

function Collection_GetNamePath(AObj: TCollection): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Collection_Insert(AObj: TCollection; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

procedure Collection_DisposeOf(AObj: TCollection); stdcall;
begin
  AObj.DisposeOf;
end;

function Collection_ClassType(AObj: TCollection): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Collection_ClassName(AObj: TCollection): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Collection_InstanceSize(AObj: TCollection): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Collection_InheritsFrom(AObj: TCollection; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Collection_Equals(AObj: TCollection; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Collection_GetHashCode(AObj: TCollection): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Collection_ToString(AObj: TCollection): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Collection_GetCapacity(AObj: TCollection): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure Collection_SetCapacity(AObj: TCollection; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function Collection_GetCount(AObj: TCollection): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function Collection_GetItems(AObj: TCollection; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure Collection_SetItems(AObj: TCollection; Index: Integer; AValue: TCollectionItem); stdcall;
begin
  AObj.Items[Index] := AValue;
end;

function Collection_StaticClassType: TClass; stdcall;
begin
  Result := TCollection;
end;

exports
  Collection_Create,
  Collection_Free,
  Collection_Owner,
  Collection_Add,
  Collection_Assign,
  Collection_BeginUpdate,
  Collection_Clear,
  Collection_ClearAndResetID,
  Collection_Delete,
  Collection_EndUpdate,
  Collection_FindItemID,
  Collection_GetNamePath,
  Collection_Insert,
  Collection_DisposeOf,
  Collection_ClassType,
  Collection_ClassName,
  Collection_InstanceSize,
  Collection_InheritsFrom,
  Collection_Equals,
  Collection_GetHashCode,
  Collection_ToString,
  Collection_GetCapacity,
  Collection_SetCapacity,
  Collection_GetCount,
  Collection_GetItems,
  Collection_SetItems,
  Collection_StaticClassType;

