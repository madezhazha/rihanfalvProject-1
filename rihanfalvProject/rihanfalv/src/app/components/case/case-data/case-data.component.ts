import { Component, OnInit,ElementRef,Renderer2 } from '@angular/core';
import{ActivatedRoute}from "@angular/router"
import{HttpClient,HttpHeaders} from '@angular/common/http'
import {DomSanitizer} from '@angular/platform-browser'

@Component({
  selector: 'app-case-data',
  templateUrl: './case-data.component.html',
  styleUrls: ['./case-data.component.css']
})
export class CaseDataComponent implements OnInit {

  constructor(
    public activatedRoute:ActivatedRoute,
    public http:HttpClient,
    private sanitizer:DomSanitizer,
    private el:ElementRef,
    private renderer2:Renderer2
  ) { }

  ngOnInit() {
    this.getData()
    this.getTxt()
  }

  //从后端拿数据所设置的变量
  title:string
  viewpoint:any
  header:any
  stringLength:number=500
  displayData:any   //显示数据
  receiveData:any   //接收从后端发来的数据
  IsPay:boolean=true     //是否已经交钱,这个是否付钱的操作决定着：是否出现查看全部
  IsCollection:boolean=false //是否已经收藏
  casecontent:any

  imageUrl:string='./assets/image/fiveStar1.PNG'

  changeImg(){
    this.imageUrl='./assets/image/fiveStar2.PNG'
  }

  money:string="5";


  sumbitFive(){
    this.money="5"
  }

  sumbitTen(){
    this.money="10"
  }

  sumbitTwenty(){
    this.money="20"
  }

  sumbitFifty(){
    this.money="50"
  }

  //接收页面传过来的值
  getData(){
    this.activatedRoute.queryParams.subscribe(params=>{
      this.title=params["title"]
      this.casecontent=params["casecontent"]
    })
    console.log(this.title)  //成功接收到数据
  }


  getTxt(){
    const httpOptions={
      headers:new HttpHeaders({'content-Type':'application/json'})
    }

    var api = "http://localhost:8009/displaytxt"

    this.http.post(api,{"content":this.title,"useid":"悟悔"},httpOptions).subscribe((response:any)=>{
      this.viewpoint=response["viewpoint"]
      // this.receiveData=this.viewpoint.substr(0,500)   //限制显示字符串的个数  现在不用字符串限制了
      this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)
      this.header=this.sanitizer.bypassSecurityTrustHtml(response['header'])

      //判断是否收藏和付款
      if(response["ispay"]==="1"){
        this.IsPay=false
        this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")
      }else{
        this.IsPay=true
      }

      if(response['iscollection']==='1'){
        this.changeImg()
      }

      console.log(this.IsPay)
    })
  }

  displayAlldata(){
    this.IsPay=false
    this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)   //为了要将HTML中的内容与原先的一模一样的显示出来
    this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")
  }

  getest(){
    const httpOptions={
      headers:new HttpHeaders({'Content-Type':"application/json"})
    }

    var api = "http://localhost:8009/alipay"

    this.http.post(api,{"content":"全部"},httpOptions).subscribe((response:any)=>{
      console.log(response)
    })
  }

}
