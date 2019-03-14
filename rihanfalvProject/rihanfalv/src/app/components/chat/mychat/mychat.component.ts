import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-mychat',
  templateUrl: './mychat.component.html',
  styleUrls: ['./mychat.component.css']
})
export class MychatComponent implements OnInit {

  // 个人信息
  public UserId:number;
  public UserName:string='qq_sasx';
  public Integral:number=0;
  public MyQueCount:number=0;
  public MyAnsCount:number=0;
  public LoginStatus:boolean=false;   //是否登录

  // 个人信息、提问、回答 列表
  public Userinfo:any;
  public QueList:any;
  public AnsList:any;

  //测试例子
  public firstlist:any;
  public secondlist:any;
  public thirdlist:any;

  constructor(public router:Router,public http:HttpClient) { }

  ngOnInit() {

    this.UserId = JSON.parse(localStorage.getItem("id"));
    if(this.UserId) this.LoginStatus=true;
    else this.LoginStatus=false;
    //console.log(this.UserId);
    //console.log(number(id));

    this.loadUserInfo();
    this.loadQueList();
    this.loadAnsList();

  }

  // 加载列表
  //加载个人信息
  loadUserInfo(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://127.0.0.1:7080/showuserinfo";    
    var postdate = {userid:this.UserId,username:"",password:"",email:"",integral:0};
    this.http.post(api,postdate,httpOptions).subscribe((response:any)=>{
      console.log(response);
        this.Userinfo=response;
        
        this.UserName=this.Userinfo.username;
        this.Integral=this.Userinfo.integral;
    })
  }
  // 加载提问列表，取数量
  loadQueList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://127.0.0.1:7080/showuserquelist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      console.log(response);
        this.QueList=response;
        
        this.MyQueCount=response.length;
    })
  }
  // 加载回答列表，取数量
  loadAnsList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://127.0.0.1:7080/showuseranslist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      console.log(response);
        this.AnsList=response;
        
        this.MyAnsCount=response.length;
    })
  }

  //跳转到我的提问列表
  myQue(){

    this.router.navigate(['/myquestion']);
  }

  //跳转到我的回答列表
  myAns(){

    this.router.navigate(['/myanswer']);
  }

  //我要提问
  toQue(){
    this.router.navigate(['/toquestion']);
  }

}
