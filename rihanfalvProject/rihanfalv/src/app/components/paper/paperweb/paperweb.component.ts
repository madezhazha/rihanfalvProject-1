import { Component, OnInit,ElementRef,ViewChild,Renderer2 } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import {Article} from '../article';
import { ActivatedRoute, Params,Router} from '@angular/router';
import {InputData} from '../../head/langing/land/input'

@Component({
  selector: 'app-paperweb',
  templateUrl: './paperweb.component.html',
  styleUrls: ['./paperweb.component.css']
})
export class PaperwebComponent implements OnInit {

  //获取dom节点
  @ViewChild('contentbox') contentbox:ElementRef;
  @ViewChild('text') text:ElementRef;
  @ViewChild('showbutton') showbutton:ElementRef;
  @ViewChild('more') more:ElementRef;

  ArticleID:string;
  UserID:string=""
  Country:string="Japan";

  Collectword:string  //bug
  //Iscollected:boolean=false;  //判断是否收藏文章
  //检查用户收藏有此文章id则显示对应字段
  Paper:Article;    //文章对象
  
  TextHeight:number;     //文章元素高度
  ContentHeight:number;  //内容盒子高度

  Islogin:boolean=false;

  url:string  //路由


  IfWantLogin:boolean=false;
  //关闭登陆框
  
  boxClose(){
    this.IfWantLogin=false;
  }

  getLoginData(date:InputData)
  {
    
    //this.finishedlogin=date.IfLogin
    if(date.IfLogin==true)
    {
      this.Islogin=true
      this.get()
    }
  }

  //展示剩余内容
  mshow(){
    if(this.Islogin==false)
    {
       //弹出登录框
       this.IfWantLogin=true;
    }
    else{
      this.renderer2.setStyle(this.contentbox.nativeElement,"height","auto")
      this.renderer2.setStyle(this.more.nativeElement,"display","none")
      this.renderer2.setStyle(this.showbutton.nativeElement,"display","none")
     }
   
  }


  //收藏
   collect(){
     if(this.Islogin==false)
     {
       //弹出登录框
       this.IfWantLogin=true;
     }
     else{ 
      const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
      let api="http://localhost:7080/paperweb/collect";
      this.http.post(api,{UserID:this.UserID,ArticleID:this.ArticleID,Country:this.Country,Iscollected:this.Paper.IsCollected},httpOptions).subscribe((response:any)=>
      {
        let result=response
        this.Paper.IsCollected=!this.Paper.IsCollected
        if(this.Paper.IsCollected==true){
          this.Collectword="取消收藏";
        }
        else{
          this.Collectword="收藏";
        }
      });

     }
    
     
   }

   //服务器获取文章详情
   get(){
    let country=localStorage.getItem("JapanOrKorea")
    if(country=="日")
    {
      this.Country="Japan"
    }
    else{
      this.Country="Korea"
    }
    this.UserID=localStorage.getItem("id")
    //console.log(this.Country)
    const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
    let api="http://localhost:7080/paperweb";
    this.http.post(api,{UserID:this.UserID,ArticleID:this.ArticleID,Country:this.Country},httpOptions).subscribe((response:any)=>
    {
      this.Paper=response
      //this.Iscollected=this.Paper.Iscollected
      //console.log(this.Iscollected)
      //console.log(this.Paper.IsCollected)
      if(this.Paper.IsCollected==true){
        this.Collectword="取消收藏";
       }
       else{
        this.Collectword="收藏";
       }
    });
   }


   back(){
     this.rout.navigate([this.url])
   }


  constructor(private http:HttpClient,private renderer2:Renderer2,private routerIonfo:ActivatedRoute,private rout:Router) { 

  }

  ngOnInit() {
    //获取用户信息
    let userid:string=localStorage.getItem("id");
    if(userid!=null)
    {
      this.Islogin=true;
      this.UserID=userid
    }
    else{
      this.Islogin=false
      this.UserID=""
    }
    //console.log(this.Islogin)

    
   //if(this.Islogin==false){
    // this.Collectword=="收藏"
   //}
    //this.Paper=this.artservice.get()  //获取要展示的文章详情

    //this.ID=this.artservice.getId();  //一.存入本地缓存显示

    //二.地址传值显示
    this.routerIonfo.params
    .subscribe((params:Params)=>{
      this.ArticleID=params['ArticleID']
      this.url=params['route']
    })
    this.get();
    //this.Iscollected=this.Paper.Iscollected
  }

}
