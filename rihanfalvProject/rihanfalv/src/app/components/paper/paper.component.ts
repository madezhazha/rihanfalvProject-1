import { Component, OnInit} from '@angular/core';
import {Article} from '../paper/article'
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { Router} from '@angular/router';
import {fromEvent} from 'rxjs'
 

@Component({
  selector: 'app-paper',
  templateUrl: './paper.component.html',
  styleUrls: ['./paper.component.css']
})
export class PaperComponent implements OnInit {

  
  Isbottom:boolean=false;  //是否页面到底
  Isover:boolean=false;    //文章是否已无
  CurrentPage:number=1;  //但前加载页数最后文章Id

  Nowcountry:string="Japan";   //当前模块 日/韩 用于进行筛选显示
  Articles:Article[]=[];  //论文列表
  Selectedarticle:Article;  //当前选择的论文


  getArticles(){
   let api="http://localhost:8000/paper";
   const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
   this.http.post(api,{CurrentPage:this.CurrentPage,Country:this.Nowcountry},httpOptions).subscribe((response:any)=>
    {
      if(response!=null)
      {
        this.Articles=this.Articles.concat(response);
        this.CurrentPage++;
        //this.CurrentPage=this.Articles[this.Articles.length-1].ID
      }
      else{
        this.Isover=true
        
      }
    });
  }

  constructor(private http:HttpClient,private router:Router) { 
    this.getArticles();
    //this.scrollCallback = this.getArticles.bind(this);
  }

  ngOnInit() {
    //this.Nowcountry=xxx.get()  //获取当前模块
    fromEvent(window,'scroll')
    .subscribe(
      ()=>{
        const h:any=document.documentElement.clientHeight;
        const H:any=document.body.clientHeight;
        const scrollTop:any=document.documentElement.scrollTop || document.body.scrollTop;
        if(h+scrollTop+20>H){
          if(!this.Isbottom){
            if(!this.Isover){
              setTimeout(()=>{this.getArticles();},1000)    //延时1000ms加载
              this.Isbottom=true
            }
          }
        }
        else
        {
          this.Isbottom=false
        }
      }
    );
  }

}
