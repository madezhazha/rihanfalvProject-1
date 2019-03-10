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
  public UserId:number=1;
  public TopicTitle:string='';
  public TopicContent:string='';
  public TopicLabel:any[]=[];

  constructor(public router:Router,public http:HttpClient) { }

  ngOnInit() {
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
    
    const httpOptions={headers:new HttpHeaders({'Content-Type':'application/json'})};

    var api='http://127.0.0.1:7080/addtopics';
    var postdate = {userid:this.UserId,topictitle:this.TopicTitle,topiccontent:this.TopicContent,japanorkorea:1,topiclabel:this.TopicLabel}
    this.http.post<string>(api,postdate,httpOptions).subscribe((response)=>{
      console.log(response);
      
      alert('提交成功！');
      this.router.navigate(['/mychat']);
      console.log(this.TopicLabel);
    })
    
  }

  //返回我的问答
  toQueReturn(){
    this.router.navigate(['/mychat']);
  }

}
