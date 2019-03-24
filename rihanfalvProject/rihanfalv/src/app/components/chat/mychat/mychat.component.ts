import { Component, OnInit, ViewChild } from '@angular/core';

import { Router } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';

import { DomSanitizer } from '@angular/platform-browser';
import { InputData } from '../../head/langing/land/input';

@Component({
  selector: 'app-mychat',
  templateUrl: './mychat.component.html',
  styleUrls: ['./mychat.component.css']
})
export class MychatComponent implements OnInit {

  @ViewChild('getwebhead') getWebHead:any;

  // 个人信息
  public UserId:number;
  public UserName:string='qq_sasx';
  public Integral:number=0;
  public Image:string;
  public MyQueCount:number=0;
  public MyAnsCount:number=0;

  //是否登录
  public LoginStatus:boolean=false;
  //是否点击登录
  public ifWantToLogin:boolean=false;   

  // 个人信息、提问、回答 列表
  public Userinfo:any;
  public QueList:any;
  public AnsList:any;

  //测试例子
  public firstlist:any;
  public secondlist:any;
  public thirdlist:any;

  constructor(public router:Router,public http:HttpClient,public sanitizer: DomSanitizer,) { }

  ngOnInit() {
    //console.log(1);
    this.UserId = JSON.parse(localStorage.getItem("id"));
    if(!this.UserId) this.LoginStatus=false;
    else {
      this.LoginStatus=true;
      this.loadUserInfo();
      this.loadQueList();
      this.loadAnsList();
    } 

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
        this.Integral=this.Userinfo.integral;
        const element = this.Userinfo;
        if (element.image.indexOf("assets") == -1) {
          let temp: any;
          temp = 'data:image/png;base64, ' + element.image; //给base64添加头缀
          element.image = this.sanitizer.bypassSecurityTrustUrl(temp);
        }
        this.Image=this.Userinfo.image;
        //console.log(this.Userinfo.image);
    })
  }
  // 加载提问列表，取数量
  loadQueList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://blackcardriver.cn:7080/showuserquelist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      //console.log(response);
        this.QueList=response;
        
        this.MyQueCount=response.length;
    })
  }
  // 加载回答列表，取数量
  loadAnsList(){

    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};
    let api="http://blackcardriver.cn:7080/showuseranslist";    
    this.http.post(api,{userid:this.UserId},httpOptions).subscribe((response:any)=>{
      //console.log(response);
        this.AnsList=response;
        
        this.MyAnsCount=response.length;
    })
  }

  //跳转到我的页面
  intoUserPage(){
    this.router.navigate(['/userpage']);
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

  //重新加载
  updateAgain(){
    this.ngOnInit();
  }

  //当前页面登录
  mychatLogin(){
    this.ifWantToLogin=true;
  }

  //关闭登陆框
  boxClose() {
    this.ifWantToLogin = false;
  }

  getLoginData(loginData: InputData) {
    if (loginData.IfLogin == true) {
      this.LoginStatus = true
      this.boxClose();
      this.ngOnInit();
      this.getWebHead.ngOnInit();
      // this.router.navigate(['/#']);
      // console.log(2);
    }
  }


}
