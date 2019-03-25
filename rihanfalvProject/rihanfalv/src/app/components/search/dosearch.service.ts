import { Injectable } from '@angular/core';
import {HttpClient,HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class DosearchService {
  constructor(public http:HttpClient) { }
  public page=1;
  public KeyWord:any=""//搜索项
  public Nowcountry:any=null//当前搜索国家
  public getkey:any={}//将传递给后端的值
  public list:any=new Array();//搜索得的表（内容）
public Classify:any="全部"//搜索分类
public Order:any="all"//排列方式
public ifget =0//判断是否获取到搜索内容,0为查找不到，1为正在查询，2为已查询到
  searchtogo(){
    if(this.nullkeyword()==true){
      return
    }
    this.page=1;
    this.ifget=1
    this.list=[]
    const httpOptions={ headers:new HttpHeaders({'Content-Type':'application/json'}) };
    let api='http://blackcardriver.cn:7080/search'; 
    //let api='http://localhost:7080/search'
    this.getkey.KeyWord=this.KeyWord
    this.getkey.Classify=this.Classify
    this.getkey.Nowcountry=this.Nowcountry
    this.getkey.Order=this.Order
    this.http.post(api,this.getkey,httpOptions).subscribe(response=>{ 

      this.list=response
      if(!response){
        this.ifget=0
      this.list=[]}
      if(response)this.ifget=2
    });
  }


  nullkeyword(){  
    const Reg = new RegExp(" ", 'gi');    
    let str=this.KeyWord.replace(Reg,"")
    if(str==""){
  return true
    }
    return false
  }
}
