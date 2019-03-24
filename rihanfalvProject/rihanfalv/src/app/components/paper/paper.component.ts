import { Component, OnInit,ElementRef,ViewChild,Renderer2 } from '@angular/core';
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

  @ViewChild('tip') tip:ElementRef
  
  Isbottom:boolean=false;  //是否页面到底
  Isover:boolean=false;    //文章是否已无
  CurrentPage:number=1;  //但前加载页数最后文章Id

  Nowcountry:string="Japan";   //当前模块 日/韩 用于进行筛选显示
  Articles:Article[]=[];  //论文列表
  Selectedarticle:Article;  //当前选择的论文


  getArticles(){
    let country=localStorage.getItem("JapanOrKorea")
    if (country=="日")
    {
      this.Nowcountry="Japan"
    }
    else{
      this.Nowcountry="Korea"
    }
   let api="http://blackcardriver.cn:7080/paper";
   const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
   this.http.post(api,{CurrentPage:this.CurrentPage,Country:this.Nowcountry},httpOptions).subscribe((response:any)=>
    {
      if(response!=null)
      {
        this.Articles=this.Articles.concat(response);
        this.CurrentPage++;
        //this.CurrentPage=this.Articles[this.Articles.length-1].ID
        console.log(this.CurrentPage)
      }
      else{
        this.Isover=true
        this.renderer2.setStyle(this.tip.nativeElement,"display","block")
      }
    });
  }

  //日韩切换初始化
  newArticles(){
    let api="http://blackcardriver.cn:7080/paper";
    const httpOptions={headers: new HttpHeaders({ 'Content-Type': 'application/json'})};
    this.http.post(api,{CurrentPage:this.CurrentPage,Country:this.Nowcountry},httpOptions).subscribe((response:any)=>
     {
       if(response!=null)
       {
        //  console.log(response)
         this.Articles=response
         this.CurrentPage++;
        //  this.CurrentPage=this.Articles[this.Articles.length].ID
       }
       else{
         this.Articles=[]
         this.Isover=true
       }
     });
   }

  getJapanKorea(isJapan:boolean){
    if(isJapan){
      this.Nowcountry="Japan"
      this.CurrentPage=1;
      this.newArticles()
    }
    else{
      this.Nowcountry="Korea"
      this.CurrentPage=1;
      this.newArticles()
    }
  }

  constructor(private http:HttpClient,private router:Router,private renderer2:Renderer2) { 
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
              this.getArticles();
              this.Isbottom=true
              console.log(this.Articles)
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
