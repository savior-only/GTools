// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {internal} from '../models';
import {util} from '../models';

export function AddDirPath(arg1:string):Promise<string>;

export function AddTodoItem(arg1:internal.TodoItem):Promise<util.Resp>;

export function DelMdDir(arg1:string):Promise<string>;

export function DelTodoItem(arg1:internal.TodoItem):Promise<util.Resp>;

export function GetDirList():Promise<util.Resp>;

export function GetMdContent(arg1:string):Promise<util.Resp>;

export function GetTodoList():Promise<util.Resp>;

export function NewMd(arg1:string,arg2:string):Promise<util.Resp>;

export function OpenMdSaveFileWindow():Promise<util.Resp>;

export function OpenSelectFileWindow():Promise<util.Resp>;

export function SaveMdContent(arg1:string,arg2:string):Promise<util.Resp>;

export function StartTomcat():Promise<string>;

export function UploadScreenshot():Promise<util.Resp>;
