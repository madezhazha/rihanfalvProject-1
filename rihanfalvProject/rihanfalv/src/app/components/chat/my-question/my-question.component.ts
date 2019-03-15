import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-my-question',
  templateUrl: './my-question.component.html',
  styleUrls: ['./my-question.component.css']
})
export class MyQuestionComponent implements OnInit {

  // 个人信息
  public UserId:number;
  public UserName:string='qq_sasx';
  public MyQueCount:number=3;
  public LoginStatus:boolean=false;

  // 回答 列表
  public Userinfo:any;
  public QueList:any;

  //显示数量
  public DisplayCount:number=3;

  // 选择状态（补充，此状态不重要）
  public QAStatus:number=1;

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
    

    this.loadUserInfo();
    this.loadQueList();
  }


  //返回我的问答页面
  detailReturn(){

    this.router.navigate(['/mychat']);
  }
  
  //显示提问列表
  detailMyQue(){

    this.QAStatus=1;
    console.log(this.QAStatus);
    this.router.navigate(['/myquestion']);
  }

  //显示回答列表
  detailMyAns(){

    this.QAStatus=2;
    console.log(this.QAStatus);
    this.router.navigate(['/myanswer']);
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
    })
  }
  // 加载提问列表
  loadQueList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://127.0.0.1:7080/showuserquelist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      console.log(response);
        this.QueList=response;
        
        this.MyQueCount=response.length;
    })
  }

  // 加载更多
  loadMore(){
    if(this.DisplayCount<this.MyQueCount)
    this.DisplayCount+=5;
    console.log(this.DisplayCount);
  }

}
