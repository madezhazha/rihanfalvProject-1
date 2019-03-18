import { Component, OnInit,ElementRef,Renderer2 } from '@angular/core';
import{ActivatedRoute}from "@angular/router"
import{HttpClient,HttpHeaders} from '@angular/common/http'
import {DomSanitizer} from '@angular/platform-browser'
import { InputData } from '../../head/langing/land/input';


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
    // this.initialState()
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
  changeNumber:number=0
  languageType:string      //语言
  titleId:string           //标题的id
  userId:string            //使用者的id
  IsLogin:boolean=false      //判断是否登录


  imageUrl:string='./assets/images/fiveStar1.PNG'

  changeImg(){
    var storage=window.localStorage;
    this.userId=storage["id"]
    if(this.userId!==undefined){
       // console.log("这个是localstorage的数据：",this.userId)
      if(this.changeNumber===0){
        //这个表示收藏信息
        this.imageUrl='./assets/images/fiveStar2.PNG'
        this.changeNumber=this.changeNumber+1
        //给后端发送post请求
        const httpOptions={
          headers:new HttpHeaders({'Content-Type':'application/json'})
        }
        var api="http://localhost:7080/changecollect"
        this.http.post(api,{"title":this.title,"data":"collect","type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
        {
          console.log(response)
        })
      }else{
        //这个表示取消收藏信息
        this.imageUrl='./assets/images/fiveStar1.PNG'
        this.changeNumber=this.changeNumber-1

        const httpOptions={
          headers:new HttpHeaders({'Content-Type':'application/json'})
        }
        var api="http://localhost:7080/changecollect"
        this.http.post(api,{"title":this.title,"data":"cancle","type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
        {
          console.log(response)
        })
      }  
    }else{
      this.IsLogin = true
    }
     
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
    })
    // console.log(this.title)  //成功接收到数据
  }


  getTxt(){
    const httpOptions={
      headers:new HttpHeaders({'content-Type':'application/json'})
    }

    var api = "http://localhost:7080/displaytxt"

    this.http.post(api,{"content":this.title},httpOptions).subscribe((response:any)=>{
      this.viewpoint=response["viewpoint"]
      // this.receiveData=this.viewpoint.substr(0,500)   //限制显示字符串的个数  现在不用字符串限制了
      this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)
      this.header=this.sanitizer.bypassSecurityTrustHtml(response['header'])

      //接收语言类型和title的id
      this.languageType=response["type"]
      this.titleId=response["ID"]   //这个ID是point的id
      this.initialState(this.languageType,this.titleId)   //这个是用来判断时候已经是收藏的状态了

      //判断是和付款
      if(response["ispay"]==="1"){
        this.IsPay=false
        this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")
      }else{
        this.IsPay=true
      }
    })
  }

  displayAlldata(){
    this.IsPay=false
    this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)   //为了要将HTML中的内容与原先的一模一样的显示出来
    this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")
  }


  //一开始的收藏状态
  initialState(languageType,titleId){
    var storage=window.localStorage;
    this.userId=storage["id"]
    if(this.userId!==undefined){
      // console.log(this.userId,this.title,this.languageType,this.titleId)  这个打印的是一开始的收藏状态
      const httpOptions={
        headers:new HttpHeaders({'content-Type':'application/json'})
      }
      var api = "http://localhost:7080/InitialState"
      this.http.post(api,{"title":this.title,"type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
      {
        if (response["data"]==='collect'){
          this.imageUrl='./assets/images/fiveStar2.PNG'
          this.changeNumber=1
        }else{
          this.imageUrl='./assets/images/fiveStar1.PNG'
          this.changeNumber=0
        }
      })
    }
  }

  boxClose(){
    this.IsLogin=false;
  }


  getLoginData(input:InputData){
    this.IsLogin = !input.IfLogin
  }
    
}
