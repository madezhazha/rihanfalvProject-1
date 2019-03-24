import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

import { HttpClient,HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-to-question',
  templateUrl: './to-question.component.html',
  styleUrls: ['./to-question.component.css']
})
export class ToQuestionComponent implements OnInit {

  // 新增主贴信息
  public UserId:number;
  public TopicTitle:string='';
  public TopicContent:string='';
  public TopicLabel:any='';
  public LoginStatus:boolean=false;
  public JapanOrKorea:number;

  constructor(public router:Router,public http:HttpClient) { }

  ngOnInit() {

    this.UserId = JSON.parse(localStorage.getItem("id"));
    if(this.UserId){
      this.LoginStatus=true;
    } 
    else {
      this.LoginStatus=false;
      this.router.navigate(['/mychat']);  //未登录跳转
    } 

    //初始化
    if( localStorage.getItem("JapanOrKorea")=="日" ){
      this.JapanOrKorea=1;
    }
    else this.JapanOrKorea=0;
    //console.log(localStorage.getItem("JapanOrKorea"));
    //console.log(this.JapanOrKorea);
  }

  // 添加主贴
  addTopics(){
    
    if(this.TopicTitle==''){
      alert('标题不能为空！');
      return 
    }
    else if(this.TopicContent==''){
      alert('提问内容不能为空！');
      return 
    }
    else if(this.TopicLabel==''){
      alert('标签内容不能为空！');
      return
    }
    
    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};

    var api='http://blackcardriver.cn:7080/addtopics';
    var postdate = {userid:this.UserId,topictitle:this.TopicTitle,topiccontent:this.TopicContent,japanorkorea:this.JapanOrKorea,topiclabel:this.TopicLabel}
    this.http.post<string>(api,postdate,httpOptions).subscribe((response)=>{
      console.log(response);
      
      alert('提交成功！');
      this.router.navigate(['/mychat']);
      console.log(this.TopicLabel);
    })
    
  }

  //获取日韩转变
  getJapanKorea(e){
    if(e){
      this.JapanOrKorea=1;
    }
    else this.JapanOrKorea=0;
  }

  //返回我的问答
  toQueReturn(){
    this.router.navigate(['/mychat']);
  }

}
