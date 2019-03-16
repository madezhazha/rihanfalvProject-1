import { Component, OnInit,Input } from '@angular/core';
import {HttpClient,HttpHeaders} from "@angular/common/http"
import {ActivatedRoute,Router} from "@angular/router"
import {fromEvent} from 'rxjs'

@Component({
  selector: 'app-casething',
  templateUrl: './casething.component.html',
  styleUrls: ['./casething.component.css']
})
export class CasethingComponent implements OnInit {

  //初始化组件
  constructor(
    public http:HttpClient,
    public activatedRoute:ActivatedRoute,
    public router:Router,
  ) { }
  ngOnInit() {
    this.getData()  //先从后端获取数据
    fromEvent(window,'scroll')
    .subscribe(
      ()=>{
        const h:any=document.documentElement.clientHeight;
        const H:any=document.body.clientHeight;
        const scrollTop:any=document.documentElement.scrollTop || document.body.scrollTop;
        if(h+scrollTop+20>H){
          if(!this.Isbottom){
            if(!this.Isover){
              setTimeout(()=>{this.getData();},1000)   
              this.Isbottom=true
            }
          }
        }
        else
        {
          this.Isbottom=false
        }
      }
    )
  }

  //定义变量
  element:any //用来接收单个元素的
  list:any[]=[]
  search:string="全部"
  languageType:string = "日"
  Isbottom:boolean = false
  Isover:boolean = false
  NumberCasething:number = 0
  CaseList:any[]=[]

   getData(search?:string,languageType?:string){
     if(search){
       this.search = search
       this.list = []
       this.NumberCasething = 0
     }
     if(languageType){
       this.languageType = languageType
       this.list = []
       this.NumberCasething = 0
     }
    const httpOptions={
      headers:new HttpHeaders({'Content-Type':'application/json'})
    }

    var api = "http://localhost:7080/alldata"

    this.http.post(api,{"content":this.search,"languageType":this.languageType,"NumberCasething":this.NumberCasething.toString()},httpOptions).subscribe((response:any)=>{
      if(JSON.stringify(response)!=="{}"){
        for(const key of Object.keys(response)){
          if(response.hasOwnProperty(key)){
            this.element = response[key]
            this.list.push(this.element)
          }  
        }
      this.NumberCasething = this.list.length
      }else{
        this.Isover = true
      }
    })
  }


  sendData(title:string){
    this.router.navigate(["/display-data"],{queryParams:{"title":title}})
  }
}
