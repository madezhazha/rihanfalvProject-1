import { Component, OnInit ,ViewChild} from '@angular/core';
import{HttpClient,HttpHeaders}from '@angular/common/http'
import {Router} from "@angular/router"
import { CasethingComponent } from './casething/casething.component';

@Component({
  selector: 'app-case',
  templateUrl: './case.component.html',
  styleUrls: ['./case.component.css']
})
export class CaseComponent implements OnInit {

  constructor(
    public http:HttpClient,
    public router:Router

  ) 
  { }

  ngOnInit() {
  }

    //下面的这些操作都是为了不同页面的显示的
    case:boolean=true;
    reason:boolean;
    time:boolean;
    level:boolean;
    //使用数组类型制作不同按钮显示不同的内容
    caseContent=[false,false,false,false,true];

    displayCase(){
      this.case=true;
      this.reason=false;
      this.time=false;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
    }
    displayReason(){
      this.case=false;
      this.reason=true;
      this.time=false;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
    displayTime(){
      this.case=false;
      this.reason=false;
      this.time=true;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
    displayLevel(){
      this.case=false;
      this.reason=false;
      this.time=false;
      this.level=true;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
  
    displayCrim(){
      this.caseContent[0]=true;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    displayCivil(){
      this.caseContent[1]=true;
      this.caseContent[0]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    dispalyAdmin(){
      this.caseContent[2]=true;
      this.caseContent[1]=false;
      this.caseContent[0]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    displayEcon(){
      this.caseContent[3]=true;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[0]=false;
      this.caseContent[4]=false;
    }
  
    //这个是全部的按钮的
    hideContent(){
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }


    //调用子组件的函数
    @ViewChild("casething")
    casething:CasethingComponent
    languageType:string="日"
    searchContent:string


    changeSearchContent(search:string){
      this.searchContent = search
      this.casething.getData(search,this.languageType)
    }

    getJapanKorea(wasJapan:boolean){
      if(wasJapan){
        this.languageType = "日"
        this.casething.getData(this.searchContent,this.languageType)
      }else{
        this.languageType = "韩"
        this.casething.getData(this.searchContent,this.languageType)
      }
    }

}
