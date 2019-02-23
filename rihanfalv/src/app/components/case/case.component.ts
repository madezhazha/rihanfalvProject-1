import { Component, OnInit } from '@angular/core';
import{HttpClient,HttpHeaders}from '@angular/common/http'
import {Router} from "@angular/router"
@Component({
  selector: 'app-case',
  templateUrl: './case.component.html',
  styleUrls: ['./case.component.css']
})
export class CaseComponent implements OnInit {

  constructor(
    public http:HttpClient,
    public router:Router

  ) 
  {
    this.DisplayContent();
   }

  ngOnInit() {
  }

    //下面的这些操作都是为了不同页面的显示的
    case:boolean=true;
    reason:boolean;
    time:boolean;
    level:boolean;
    //使用数组类型制作不同按钮显示不同的内容
    caseContent=[false,false,false,false,true];
    searchContent:string="全部";
    public list :any[]=[]  //用来接收数据的
    data:any

  //这里用来分页显示的
  pageNo = 1; //当前页码
  preShow = false; //上一页
  nextShow = true; //下一页
  pageSize = 10; //单页显示数
  // totalCount = 0; //总页数
  pageSizes = [10,0,5,10, 15]; 
  curPage = 1; //当前页
  PageList = [] //分页后前台显示数据


  //为什么我一开始是空呢?这是因为tablePageList这个一开始是空的因此需要点击一次才能调用这个函数
  getPageList() {
    if (this.list.length >= 1) {
      if (this.list.length % this.pageSize === 0) {
        this.pageNo = Math.floor(this.list.length / this.pageSize);   //用来做四舍五入的
      } else {
        this.pageNo = Math.floor(this.list.length / this.pageSize) + 1;
      }
      if (this.pageNo < this.curPage) {
        this.curPage = this.curPage - 1;
      }
      if (this.pageNo === 1 || this.curPage === this.pageNo) {
        this.preShow = this.curPage !== 1;
        this.nextShow = false;
      } else {
        this.preShow = this.curPage !== 1;
        this.nextShow = true;
      }
    } else {
      this.list.length = 0;
      this.pageNo = 1;
      this.curPage = 1;
    }
    this.PageList = this.list.slice((this.curPage - 1) * this.pageSize, (this.curPage) * this.pageSize);   //切片
    // console.log(this.list)

  }
  //点击上一页方法
  showPrePage() {
    this.curPage--;
    if (this.curPage >= 1) {
      this.getPageList();
    } else {
      this.curPage = 1;
    }
  }
//点击下一页方法
  showNextPage() {
    this.curPage++;
    if (this.curPage <= this.pageNo) {
      this.getPageList();
    } else {
      this.curPage = this.pageNo;
    }
  }
//自定义跳页方法
  onChangePage(value) {
    if (value > this.pageNo) {
      confirm('超出最大页数');
    } else if (value <= 0) {
      this.curPage = 1;
      this.getPageList();
    } else {
      this.curPage = value;
      this.getPageList();
    }
  }
  //改变每页显示方法
  onChangePageSize(value) {
    this.pageSize = value;
    this.curPage = 1;
    this.getPageList();
  }


  
    displayCase(){
      this.case=true;
      this.reason=false;
      this.time=false;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
    }
    displayReason(){
      this.case=false;
      this.reason=true;
      this.time=false;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
    displayTime(){
      this.case=false;
      this.reason=false;
      this.time=true;
      this.level=false;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
    displayLevel(){
      this.case=false;
      this.reason=false;
      this.time=false;
      this.level=true;
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
  
    displayCrim(){
      this.caseContent[0]=true;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    displayCivil(){
      this.caseContent[1]=true;
      this.caseContent[0]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    dispalyAdmin(){
      this.caseContent[2]=true;
      this.caseContent[1]=false;
      this.caseContent[0]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    displayEcon(){
      this.caseContent[3]=true;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[0]=false;
      this.caseContent[4]=false;
    }
  
    //这个是全部的按钮的
    hideContent(){
      this.caseContent[0]=false;
      this.caseContent[1]=false;
      this.caseContent[2]=false;
      this.caseContent[3]=false;
      this.caseContent[4]=false;
    }
  
    //上面设置的都是第一，第二层的页面的设置，下面是从数据库接受到数据之后的显示
    // num:number=this.list.length;

    changeSearchContent(content:string){
      this.searchContent=content;
      this.list=[]

      //这里使用get请求就行了，发送的数据再说

      const httpOptions={
        headers:new HttpHeaders({'Content-Type':'application/json'})
      }

      var api = "http://localhost:8009/alldata"

      // 注意这里的content是要搞事情的

      this.http.post(api,{"content":content},httpOptions).subscribe((response:any)=>{
        // console.log(response)
        //遍历对象，并且将数据放在一个数组中
        for(const key of Object.keys(response)){
          if(response.hasOwnProperty(key)){
            this.data=response[key]
            this.list.push(this.data)
          }
        }

        // console.log(this.list)  用来测试，能否成功的将数据接收到
        this.onChangePageSize("10")
      })

    }


  //这个搞定跳转页面，但是不改变地址
  goTo(location){
    window.location.hash = "";
    window.location.hash = location
  }


  //一开始就显示数据
  DisplayContent(){

    //使用post请求
    const httpOptions={
      headers:new HttpHeaders({'Content-Type':'application/json'})
    }

    var api = "http://localhost:8009/alldata"

    // 注意这里的content是要搞事情的

    this.http.post(api,{"content":"全部"},httpOptions).subscribe((response:any)=>{
      // console.log(response)
      //遍历对象，并且将数据放在一个数组中
      for(const key of Object.keys(response)){
        if(response.hasOwnProperty(key)){
          this.data=response[key]
          this.list.push(this.data)
        }
      }

      // console.log(this.list)

      //为什么会在这里显示呢？这是因为执行一次：getPageList()不起任何作用，list的值不会传进去
      this.onChangePageSize("10")
    })
  }


  //页面传值，通过url传值
  sendData(title:string){
    this.router.navigate(["/display-data"],{queryParams:{"title":title}})
  }


}
