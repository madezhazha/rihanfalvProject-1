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
public list:any[]=[]
public ifreadbutton=true//阅读更多按钮是否显示
public ifreadall=false//是否阅读全部
  constructor(public router: Router,public http:HttpClient,public m_search:DosearchService) { }

  ngOnInit() {
this.ifreadbutton=true
this.ifreadall=false
  }

  getstatus(){
    if(!this.m_search.list){
          return false //未获取数据     
    }
    if(this.m_search.list.length==0){
      return false //未获取数据     
}
    this.list=this.m_search.list
    if(this.m_search.list.length<=2){
      this.ifreadall=true
      this.ifreadbutton=false
        //数据少于2条，显示全部
    }
    else{
      //数据多于2条，显示前2条
    }
     return true

  }
   back() {    // 返回上一页面
    this.router.navigate(['/Search']);
  }
  readall(){
     this.ifreadbutton=false//阅读更多按钮是否显示
 this.ifreadall=true//是否阅读全部
  }



}
