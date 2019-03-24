import { Component, OnInit,ElementRef,ViewChild,Renderer2} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import{DosearchService}from "../dosearch.service"
import { Router } from '@angular/router';
import { ApiSerivice } from '../../../services/apiservice';

@Component({
  selector: 'app-searchresult',
  templateUrl: './searchresult.component.html',
  styleUrls: ['./searchresult.component.css']
})

export class SearchresultComponent implements OnInit {
public list=new Array();
public page=1


@ViewChild('click') changeclass:ElementRef;

  constructor(public router: Router,public http:HttpClient,public m_search:DosearchService,private renderer2:Renderer2,private api: ApiSerivice) { 

  }

  ngOnInit() {
  }

  getstatus(){
    if(this.m_search.ifget==false){
  this.list=[];
  return this.m_search.ifget
}
     if(this.m_search.ifget==true){
  if(this.m_search.list.length>0 &&this.m_search.list.length<5){
    this.list=this.m_search.list
    return this.m_search.ifget
  }
  if(this.m_search.list.length>=5){
    for(let i =0;i<5;i++){
   this.list[i]=this.m_search.list[i]}
   return this.m_search.ifget
  }
  }


    
  }

readmore(){
  let length=this.m_search.list.length
  let addlist=new Array();
  if(length<5)return;
  if(this.page*5>length)
    return

  for(let i =0;i<5;i++){
    if(this.m_search.list[this.page*5+i])
     addlist[i]=this.m_search.list[this.page*5+i]
  }
    this.page++
    this.list=this.list.concat(addlist);
}

click(item){
  this.m_search.Classify=item
this.m_search.searchtogo()
  this.renderer2.setStyle(this.changeclass.nativeElement,"background-color","0000FF")
}

turntocase(title){//案例跳转
  this.router.navigate(["/display-data"],{queryParams:{"title":title}})
}

turntolegal(Title){//法律跳转
  console.log(Title)
  this.api.legaltitle2(Title).subscribe();
  let that=this
  setTimeout( function() {  that.router.navigate(['/content']) } , 200);   // 延时触发，给服务器留反应时

}
}
