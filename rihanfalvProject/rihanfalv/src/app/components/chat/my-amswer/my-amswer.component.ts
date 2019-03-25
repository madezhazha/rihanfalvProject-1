import { Component, OnInit } from '@angular/core';

import { Router, NavigationExtras } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';
import { fromEvent } from 'rxjs';

@Component({
  selector: 'app-my-amswer',
  templateUrl: './my-amswer.component.html',
  styleUrls: ['./my-amswer.component.css']
})
export class MyAmswerComponent implements OnInit {

  // 个人信息
  public UserId:number;
  public UserName:string='qq_sasx';
  public MyAnsCount:number=0;
  public LoginStatus:boolean=false;

  // 回答 列表
  public Userinfo:any;
  public AnsList:any;

  //显示数量
  public DisplayCount:number=4;

  // 当前数据的数组是否是最大值
  public isMax: boolean = false;
  public nowData: Array<any> = [];
  public isBottom: boolean = false;

  constructor(public router:Router,public http:HttpClient) { }

  ngOnInit() {

    this.UserId = JSON.parse(localStorage.getItem("id"));
    if(this.UserId){
      this.LoginStatus=true;
    } 
    else {
      this.LoginStatus=false;
      this.router.navigate(['/mychat']);  //未登录跳转
    } 
    
    fromEvent(window, 'scroll').subscribe(() => {
      const h: any = document.documentElement.clientHeight;
      const H: any = document.body.clientHeight;
      const scrollTop: any = document.documentElement.scrollTop || document.body.scrollTop;
      if (h + scrollTop + 20 > H) {
        if (!this.isBottom) {
          setTimeout(() => { this.loadMore()}, 500);
        }
        this.isBottom = true;
      } else {
        this.isBottom = false;
      }
    });

    this.loadUserInfo();
    this.loadAnsList();
  }

  // 加载列表
  //加载个人信息
  loadUserInfo(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://blackcardriver.cn:7080/showuserinfo";    
    var postdate = {userid:this.UserId,username:"",password:"",email:"",integral:0};
    this.http.post(api,postdate,httpOptions).subscribe((response:any)=>{
      //console.log(response);
        this.Userinfo=response;
        
        this.UserName=this.Userinfo.username;
    })
  }
  // 加载回答列表
  loadAnsList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://blackcardriver.cn:7080/showuseranslist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      //console.log(response);
        this.AnsList=response;
        
        this.MyAnsCount=response.length;
        
        this.cutArray(this.AnsList, 8);
        
    })
  }
  
  //点击阅读原文，阅读量加一
  readTopic(item:any){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://blackcardriver.cn:7080/addtopicvisnum";    
    this.http.post(api,{topicid:item.topicid},httpOptions).subscribe((response:any)=>{
      let navigationExtras: NavigationExtras = {
        queryParams: {
          topicID: item.topicid,
        },
      }
      this.router.navigate(['/post'], navigationExtras);
      //console.log(navigationExtras);
      //console.log(response);
        
    })
  }

  // 加载更多
  loadMore(){
    
    this.cutArray(this.AnsList, 2);
  }
  
  // 截取展示数据，并且判断现在长度是否为最大
  cutArray(countryData: Array<any>, length: number) {
    if (countryData) {
      this.nowData = countryData.slice(0, this.nowData.length + length);
      if (this.nowData.length == countryData.length) {
        this.isMax = true;
      } else {
        this.isMax = false;
      }
    } else {
      this.isMax = true;
    }
  }

}
