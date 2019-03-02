import { Component, OnInit,ElementRef,ViewChild,Renderer2 } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import {Article} from '../article';
import { ActivatedRoute, Params} from '@angular/router';

@Component({
  selector: 'app-paperweb',
  templateUrl: './paperweb.component.html',
  styleUrls: ['./paperweb.component.css']
})
export class PaperwebComponent implements OnInit {

  //获取dom节点
  @ViewChild('contentbox') contentbox:ElementRef;
  @ViewChild('text') text:ElementRef;
  @ViewChild('show') showbutton:ElementRef;
  @ViewChild('more') more:ElementRef;

  ArticleID:string;
  UserID:string;
  Country:string="Japan";

  Collectword:string="收藏";  //bug
  Iscollected:boolean=false;  //判断是否收藏文章
  //检查用户收藏有此文章id则显示对应字段
  Paper:Article;    //文章对象
  
  TextHeight:number;     //文章元素高度
  ContentHeight:number;  //内容盒子高度

  Islogin:boolean=false;

  result:any={
    Result:0
  }

  //展示剩余内容
  mshow(){
    if(this.Islogin==false)
     {
       //弹出登录框
     }
     else{
      this.renderer2.setStyle(this.showbutton.nativeElement,"display","none")
      this.renderer2.setStyle(this.contentbox.nativeElement,"height","auto")
      this.renderer2.setStyle(this.more.nativeElement,"display","none")
     }
   
  }


  //收藏
   collect(){
     if(this.Islogin==false)
     {
       //弹出登录框
     }
     else{ 
      const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
      let api="http://localhost:8000/paperweb/collect";
      this.http.post(api,{UserID:this.UserID,ArticleID:this.ArticleID,Country:this.Country,Iscollected:this.Iscollected},httpOptions).subscribe((response:any)=>
      {
        this.result=response
        if(this.result.Result==1){
          this.Iscollected=!this.Iscollected;
          if(this.Iscollected==true){
           this.Collectword="取消收藏";
          }
          else{
           this.Collectword="收藏";
          }
        }
        else{
          console.log("false")
        }
      });

     }
    
     
   }

   //服务器获取文章详情
   get(){
    const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
    let api="http://localhost:8000/paperweb";
    this.http.post(api,{UserID:this.UserID,ArticleID:this.ArticleID,Country:this.Country},httpOptions).subscribe((response:any)=>
    {
      this.Paper=response
    });
   }


  constructor(private http:HttpClient,private renderer2:Renderer2,private routerIonfo:ActivatedRoute) { 

  }

  ngOnInit() {
    //获取用户信息
    let userid:string=localStorage.getItem("id");
    if(userid!="")
    {
      this.Islogin=true;
    }
    
   if(this.Islogin==false){
     this.Collectword=="收藏"
   }
    //this.Paper=this.artservice.get()  //获取要展示的文章详情

    //this.ID=this.artservice.getId();  //一.存入本地缓存显示

    //二.地址传值显示
    this.routerIonfo.params
    .subscribe((params:Params)=>{
      this.ArticleID=params['ArrticleID']
    })
    this.get();
    this.Iscollected=this.Paper.Iscollected
  }

}
