import { Component, OnInit,ElementRef,ViewChild,Renderer2} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import{DosearchService}from "../dosearch.service"
import { Router } from '@angular/router';
@Component({
  selector: 'app-searchresult',
  templateUrl: './searchresult.component.html',
  styleUrls: ['./searchresult.component.css']

})

export class SearchresultComponent implements OnInit {
public list=new Array();
public text :string="正在获取数据..."
public page=1


@ViewChild('click') changeclass:ElementRef;

  constructor(public router: Router,public http:HttpClient,public m_search:DosearchService,private renderer2:Renderer2) { 
  }

  ngOnInit() {
  }

  getstatus(){
    if(!this.m_search.list){
         this.text="未获取数据..." //未获取数据
         return false     
    }
    if(this.m_search.list=="null"){
      this.text="数据不存在"  
      return false   
}
if(this.m_search.list.length<5){
  this.list=this.m_search.list
  this.text=null;

}
if(this.m_search.list.length>=5)
for(let i =0;i<5;i++){
   this.list[i]=this.m_search.list[i]
   this.text=null;

}

     return true

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
}
