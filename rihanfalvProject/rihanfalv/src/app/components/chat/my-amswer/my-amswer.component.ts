import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-my-amswer',
  templateUrl: './my-amswer.component.html',
  styleUrls: ['./my-amswer.component.css']
})
export class MyAmswerComponent implements OnInit {

  // 个人信息
  public UserId:number=1;
  public UserName:string='qq_sasx';
  public MyAnsCount:number=0;

  // 回答 列表
  public Userinfo:any;
  public AnsList:any;

  //显示数量
  public DisplayCount:number=3;

  // 选择状态（补充，此状态不重要）
  public QAStatus:number=1;

  constructor(public router:Router,public http:HttpClient) { }

  ngOnInit() {

    this.loadUserInfo();
    this.loadAnsList();
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
  // 加载回答列表
  loadAnsList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://127.0.0.1:7080/showuseranslist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      console.log(response);
        this.AnsList=response;
        
        this.MyAnsCount=response.length;
    })
  }

  // 加载更多
  loadMore(){
    if(this.DisplayCount<this.MyAnsCount)
    this.DisplayCount+=5;
    console.log(this.DisplayCount);
  }

}
