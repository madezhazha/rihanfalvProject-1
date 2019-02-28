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

  ID:number; 
  Collectword:string="收藏";  //bug
  Iscollected:boolean=false;  //判断是否收藏文章
  //检查用户收藏有此文章id则显示对应字段
  Paper:Article;    //文章对象
  
  TextHeight:number;     //文章元素高度
  ContentHeight:number;  //内容盒子高度


  //展示剩余内容
  mshow(){
    this.renderer2.setStyle(this.showbutton.nativeElement,"display","none")
    this.renderer2.setStyle(this.contentbox.nativeElement,"height","auto")
    this.renderer2.setStyle(this.more.nativeElement,"display","none")
  }


  //收藏
   collect(){
     this.Iscollected=!this.Iscollected;
     if(this.Iscollected==true){
       this.Collectword="取消收藏";
     }
     else{
      this.Collectword="收藏";
     }
   }

   //服务器获取文章详情
   get(){
    const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
    let api="http://localhost:8000/article";
    this.http.post(api,this.ID,httpOptions).subscribe((response:any)=>
    {
      this.Paper=response
    });
   }


  constructor(private http:HttpClient,private renderer2:Renderer2,private routerIonfo:ActivatedRoute) { 

  }

  ngOnInit() {
    //this.Paper=this.artservice.get()  //获取要展示的文章详情

    //this.ID=this.artservice.getId();  //一.存入本地缓存显示

    //二.地址传值显示
    this.routerIonfo.params
    .subscribe((params:Params)=>{
      this.ID=params['ID']
    })
    this.get();
  }

}
