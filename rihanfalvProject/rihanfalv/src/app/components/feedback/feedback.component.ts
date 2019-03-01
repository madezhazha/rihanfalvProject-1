import { Component, OnInit } from '@angular/core';
import { Router, NavigationExtras } from '@angular/router';
import { HttpClient, HttpHandler, HttpHeaders} from '@angular/common/http';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.css']
})
export class FeedbackComponent implements OnInit {

  public neixing:any
  public mark:any
  public link:any
  constructor(public http:HttpClient,public router:Router) { }

  ngOnInit() {
  }
  submit(){
    if(this.neixing==null){
          alert('类型不能为空')
    }
    else if(this.mark==null ){
      alert('意见不能为空')
    }
    
    else{

    const httpOptions={headers:new HttpHeaders({'Content-Type' :'application/json'})}
                  let api='http://localhost:7080/addfeedback';
                  this.http.post(api,{"userid":'9527','feedbacktype':this.neixing,'feedbackcontent':this.mark},httpOptions).subscribe((response:any)=>{
                      
                  })

    this.router.navigate(['/Feedbacksuccess/'])

    }
  }

}
