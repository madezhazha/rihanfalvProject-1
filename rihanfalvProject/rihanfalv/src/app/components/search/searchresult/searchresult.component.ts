import { Component, OnInit,ElementRef,ViewChild,Renderer2} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import{DosearchService}from "../dosearch.service"
import { Router } from '@angular/router';
import { ApiSerivice } from '../../../services/apiservice';
import {fromEvent} from 'rxjs'

@Component({
  selector: 'app-searchresult',
  templateUrl: './searchresult.component.html',
  styleUrls: ['./searchresult.component.css']
})

export class SearchresultComponent implements OnInit {
public list=new Array();
public Isbottom: boolean = false;


@ViewChild('click') changeclass:ElementRef;

  constructor(public router: Router,public http:HttpClient,public m_search:DosearchService,private renderer2:Renderer2,private api: ApiSerivice) { 

  }

  ngOnInit() {

fromEvent(window,'scroll')
    .subscribe(
      ()=>{
        const h:any=document.documentElement.clientHeight;
        const H:any=document.body.clientHeight;
        const scrollTop:any=document.documentElement.scrollTop || document.body.scrollTop;
        if(h+scrollTop+20>H){
          if(!this.Isbottom){
           
              setTimeout(()=>{this.readmore();},700)    //延时700ms加载
            
          }
        }
        else
        {
          this.Isbottom=false
        }
      }
    );

  }

  getstatus(){
    if(this.m_search.ifget==0||this.m_search.ifget==1){//0为查找不到，1为正在查询，2为已查询到
  this.list=[];
  return false
}
     if(this.m_search.ifget==2){
  if(this.m_search.list.length>0 &&this.m_search.list.length<5){
    this.Isbottom=true;
    this.list=this.m_search.list
    return true
  }
  if(this.m_search.list.length>=5){
    for(let i =0;i<5;i++){
   this.list[i]=this.m_search.list[i]}
   return true
  }
  }
    
  }

readmore(){
  let length=this.m_search.list.length
  let addlist=new Array();
  if(length<5){
    return;
  }
  if(this.m_search.page*5>length)
  {
    this.Isbottom=true
    return
  }

  for(let i =0;i<5;i++){
    if(this.m_search.list[this.m_search.page*5+i])
     addlist[i]=this.m_search.list[this.m_search.page*5+i]
  }
    this.m_search.page++
    this.list=this.list.concat(addlist);
}

click(item){
  this.m_search.Classify=item
this.m_search.searchtogo()
}

turntocase(title){//案例跳转
  this.router.navigate(["/display-data"],{queryParams:{"title":title}})
}

turntolegal(Title){//法律跳转
  this.api.legaltitle2(Title).subscribe();
  let that=this
  setTimeout( function() {  that.router.navigate(['/content']) } , 200);   // 延时触发，给服务器留反应时

}
}
