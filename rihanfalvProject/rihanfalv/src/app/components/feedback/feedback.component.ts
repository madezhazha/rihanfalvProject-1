import { Component, OnInit } from '@angular/core';
import { Router, NavigationExtras } from '@angular/router';
import { HttpClient, HttpHandler, HttpHeaders} from '@angular/common/http';
import {InputData} from '../head/langing/land/input'
@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.css']
})
export class FeedbackComponent implements OnInit {

  public neixing:any  //反馈类型
  public mark:any    //反馈内容
  //public link:any    
  public id:string   //用户id
  Islogin:boolean=false;
  IfWantLogin:boolean=false;
  constructor(public http:HttpClient,public router:Router) { }

  ngOnInit() {
    //this.id=JSON.parse(localStorage.getItem('id'))    //获取用户id
    //alert(this.id)
  }

  boxClose(){
    this.IfWantLogin=false;
  }

  getLoginData(date:InputData)
  {
    
    if(date.IfLogin==true)
    {
      this.Islogin=true
    }
  }

  submit(){
    this.id=JSON.parse(localStorage.getItem('id'))
   /*if(this.id!=null){
      this.boxClose

    }*/
    if(this.id==null){
      //弹出登录框
      this.IfWantLogin=true;
   }
    else if(this.neixing==null){
          alert('类型不能为空')
    }
    else if(this.mark==null ){
      alert('意见不能为空')
    }
    
    else{

    const httpOptions={headers:new HttpHeaders({'Content-Type' :'application/json'})}
                  let api='http://localhost:7080/addfeedback';
                  this.http.post(api,{"userid":this.id,'feedbacktype':this.neixing,'feedbackcontent':this.mark},httpOptions).subscribe((response:any)=>{
                      
                  })

    this.router.navigate(['/Feedbacksuccess/'])

    }
  }
  bin(){
    this.router.navigate(['/homepage/'])
    
  }

}
