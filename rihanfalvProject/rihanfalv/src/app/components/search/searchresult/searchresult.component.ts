import { Component, OnInit} from '@angular/core';
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
public text :string
public page=1
  constructor(public router: Router,public http:HttpClient,public m_search:DosearchService) { 
  }

  ngOnInit() {

  }

  getstatus(){
    if(!this.m_search.list){
         this.text="未获取搜索数据" //未获取数据
         return false     
    }
    if(this.m_search.list.length==0){
      this.text="搜索目标不存在"  
      return false   
}
for(let i =0;i<5;i++){
  if(this.m_search.list[i])
   this.list[i]=this.m_search.list[i]
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
   console.log(this.list)
}


}
