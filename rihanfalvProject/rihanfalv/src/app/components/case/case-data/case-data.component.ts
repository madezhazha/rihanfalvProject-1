import { Component, OnInit,ElementRef,Renderer2,ViewChild,Output,EventEmitter } from '@angular/core';
import{ActivatedRoute,Router}from "@angular/router"
import{HttpClient,HttpHeaders} from '@angular/common/http'
import {DomSanitizer} from '@angular/platform-browser'
import { InputData } from '../../head/langing/land/input';
import {Location} from '@angular/common'


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
    private renderer2:Renderer2,
    private location : Location,
    private router:Router,
  ) { }

  ngOnInit() {
    this.getData()
    this.getTxt()
  }

  //获取dom节点
  @ViewChild("test") test:ElementRef


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
  integral:string            //积分
  allintegral:string         //改用用户的所有的积分
  username:string            //账号名
  search:string              //查找的东西
  rechargeIntegral:string="50"    //充值的积分


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
        var api="http://blackcardriver.cn:7080/changecollect"
        this.http.post(api,{"title":this.title,"data":"collect","type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
        {
          if(response!==null){
            if(response["data"]==="系统出现错误"){
              alert("系统出现错误")
            }
          }
        })
      }else{
        //这个表示取消收藏信息
        this.imageUrl='./assets/images/fiveStar1.PNG'
        this.changeNumber=this.changeNumber-1

        const httpOptions={
          headers:new HttpHeaders({'Content-Type':'application/json'})
        }
        var api="http://blackcardriver.cn:7080/changecollect"
        this.http.post(api,{"title":this.title,"data":"cancle","type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
        {
          if(response!==null){
            if(response["data"]==="系统出现错误"){
              alert("系统出现错误")
            }
          }
        })
      }  
    }else{
      this.IsLogin = true
    }
     
  }



  money:string="5";


  sumbitFive(){
    this.money="5"
    this.rechargeIntegral = "50"
  }

  sumbitTen(){
    this.money="10"
    this.rechargeIntegral = "120"
  }

  sumbitTwenty(){
    this.money="20"
    this.rechargeIntegral = "300"
  }

  sumbitFifty(){
    this.money="50"
    this.rechargeIntegral = "800"
  }

  //接收页面传过来的值
  getData(){
    this.activatedRoute.queryParams.subscribe(params=>{
      this.title=params["title"]
      this.search = params["type"]
    })
  }


  getTxt(){
    var storage=window.localStorage;
    this.userId=storage["id"]
    const httpOptions={
      headers:new HttpHeaders({'content-Type':'application/json'})
      }

      var api = "http://blackcardriver.cn:7080/displaytxt"

    if(this.userId ===undefined){
      this.http.post(api,{"content":this.title},httpOptions).subscribe((response:any)=>{
        if(response["data"]!=="系统出现错误"){
          this.viewpoint=response["viewpoint"]
          this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)
          this.header=this.sanitizer.bypassSecurityTrustHtml(response['header'])
          //接收语言类型和title的id
          this.languageType=response["type"]
          this.titleId=response["ID"]   //这个ID是point的id
          this.integral = response["integral"]
          // console.log(this.titleId,this.integral)
        }else{
          alert("系统出现错误")
        }   
      })
    }else{ 
        this.http.post(api,{"content":this.title,"userid":this.userId},httpOptions).subscribe((response:any)=>{
          if(response["data"]!=="系统出现错误"){
            this.viewpoint=response["viewpoint"]
            this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)
            this.header=this.sanitizer.bypassSecurityTrustHtml(response['header'])
            //接收语言类型和title的id
            this.languageType=response["type"]
            this.titleId=response["ID"]   //这个ID是point的id
            this.integral = response["integral"]   //这个是观看法官观点的所需要的积分
            this.allintegral = response["allintergral"]
            this.username = response["username"]
            if(response["searchResult"]==="1"){
              this.IsPay = false
              this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")
            }
            // console.log(this.titleId,this.integral,this.userId,response["searchResult"],this.allintegral)
            this.initialState(this.languageType,this.titleId)   //这个是用来判断时候已经是收藏的状态了 
          }else{
            alert("系统出现错误")
          } 
          // console.log(response)       
        })
    }
  }

  displayAlldata(){
    var storage=window.localStorage;
    this.userId=storage["id"]

    if(this.userId!==undefined){
      this.renderer2.setAttribute(this.test.nativeElement,"data-toggle","k")
      const httpOptions={
        headers:new HttpHeaders({'content-Type':'application/json'})
      }
  
      var api = "http://blackcardriver.cn:7080/payment"
  
      this.http.post(api,{"titleid":this.titleId,"userid":this.userId,"integral":this.integral},httpOptions).subscribe((response:any)=>{
        if(response["data"]!=="系统出现错误"){
            if(response["data"]!=="积分不够"){
            this.IsPay = false
            this.displayData=this.sanitizer.bypassSecurityTrustHtml(this.viewpoint)   //为了要将HTML中的内容与原先的一模一样的显示出来
            this.renderer2.setStyle(this.el.nativeElement.querySelector(".judgePoint"),'height',"auto")        
          }else{
            alert("积分不够")
            this.renderer2.setAttribute(this.test.nativeElement,"data-toggle","modal")
          }
        }else{
          alert("系统出现错误")
        }
      })
    }else{
      this.IsLogin = true
    }

  }

  print(){
    var storage=window.localStorage;
    this.userId=storage["id"]
    if(this.userId===undefined){
      this.renderer2.setAttribute(this.test.nativeElement,"data-toggle","k")
      this.IsLogin = true
    }
  }


  //一开始的收藏状态
  initialState(languageType,titleId){
    var storage=window.localStorage;
    this.userId=storage["id"]
    if(this.userId!==undefined){
      const httpOptions={
        headers:new HttpHeaders({'content-Type':'application/json'})
      }
      var api = "http://blackcardriver.cn:7080/InitialState"
      this.http.post(api,{"title":this.title,"type":this.languageType,"titleId":this.titleId,"userid":this.userId},httpOptions).subscribe((response:any)=>
      {
        if(response["data"]!=="系统出现错误"){
          if (response["data"]==='collect'){
            this.imageUrl='./assets/images/fiveStar2.PNG'
            this.changeNumber=1
          }else{
            this.imageUrl='./assets/images/fiveStar1.PNG'
            this.changeNumber=0
          }
        }else{
          alert("系统出现错误")
        }
      })
    }
  }

  boxClose(){
    this.IsLogin=false;
  }


  getLoginData(input:InputData){
    this.IsLogin = !input.IfLogin
    this.initialState(this.languageType,this.titleId)    //这个表示一登录就能判断出是否已经收藏了
    this.getTxt()        //这个要一开始就要判断是否已经付费为了
  }

  goBack(){
    history.go(-1)
  }
    

  //充值
  recharge(){

    var storage=window.localStorage;
    this.userId=storage["id"]

    const httpOptions={
      headers:new HttpHeaders({'content-Type':'application/json'})
    }

    var api = "http://blackcardriver.cn:7080/recharge"

    this.http.post(api,{"integral":this.rechargeIntegral,"userid":this.userId,"allintegral":this.allintegral},httpOptions).subscribe((response:any)=>{
      console.log("充值成功")
      console.log(response)
      if(response){
        alert("系统出现错误")
      }
    })
    
  }
  

}
