import { Component, OnInit } from '@angular/core';
import{ActivatedRoute}from "@angular/router"
import{HttpClient,HttpHeaders} from '@angular/common/http'

@Component({
  selector: 'app-case-data',
  templateUrl: './case-data.component.html',
  styleUrls: ['./case-data.component.css']
})
export class CaseDataComponent implements OnInit {

  constructor(
    public activatedRoute:ActivatedRoute,
    public http:HttpClient
  ) { }

  ngOnInit() {
    this.getData()

    this.getTxt()
  }

  //从后端拿数据所设置的变量
  title:string
  all_data:any  //法官观点
  display_data:any
  firstinstance:any  //一审判决书
  secondtrial:any   //二审判决书
  thirdtrial:any    //再审判决书
  publicoffice:any  //公诉机关
  plaintiff:any     //原告人
  agent:any         //委托人
  defendant:any     //被告人
  counsel:any       //辩护人
  trialgrade:any    //审级
  firstcourt:any    //一审法院
  firstpeople:any   //一审合议庭组成人员
  secondcourt:any   //二审法院
  secondpeople:any  //二审合议庭组成人员
  retrial:any       //再审法院
  retrialpeople:any //再审合议庭组成人员
  firsttime:any     //一审结束时间
  secondtime:any    //二审结束时间
  retrialtime:any   //再审结束时间


  


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
    })
    console.log(this.title)  //成功接收到数据
  }


  getTxt(){
    const httpOptions={
      headers:new HttpHeaders({'content-Type':'application/json'})
    }

    var api = "http://localhost:8009/displaytxt"

    this.http.post(api,{"content":this.title},httpOptions).subscribe((response:any)=>{
      this.all_data=response["judgepoint"]   //法官观点
      this.display_data=this.all_data.substr(0,300)   //法官观点的截取
      this.firstinstance = response["firstinstance"]
      this.secondtrial = response["secondtrial"]
      this.thirdtrial = response["thirdtrial"]
      this.publicoffice = response["publicoffice"]
      this.plaintiff = response["plaintiff"]
      this.agent = response["agent"]
      this.defendant = response["defendant"]
      this.counsel = response["counsel"]
      this.trialgrade = response["trialgrade"]
      this.firstcourt = response["firstcourt"]
      this.firstpeople = response["firstpeople"]
      this.secondcourt = response["secondcourt"]
      this.secondpeople = response["secondpeople"]
      this.retrial = response["retrial"]
      this.retrialpeople = response["retrialpeople"]
      this.firsttime = response["firsttime"]
      this.secondtime = response["secondtime"]
      this.retrialtime = response["retrialtime"]
    })
  }

  displayAlldata(){
    this.display_data=this.all_data
  }

}
