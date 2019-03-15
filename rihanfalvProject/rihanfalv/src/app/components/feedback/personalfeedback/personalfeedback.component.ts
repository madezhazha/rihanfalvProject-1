import { Component, OnInit } from '@angular/core';
import { Router, NavigationExtras } from '@angular/router';
import { HttpClient, HttpHandler, HttpHeaders} from '@angular/common/http';
@Component({
  selector: 'app-personalfeedback',
  templateUrl: './personalfeedback.component.html',
  styleUrls: ['./personalfeedback.component.css']
})
export class PersonalfeedbackComponent implements OnInit {

  public list :any[]=[];
  public id:string   //用户id
  constructor(public http:HttpClient,public router:Router) { }

  ngOnInit() {
    this.id=JSON.parse(localStorage.getItem('id'))    //获取用户id
    
    if(this.id==null){
      alert('请登陆')
    }
    else {
    const httpOptions={headers:new HttpHeaders({'Content-Type' :'application/json'})}
    let api='http://localhost:7080/userfeedback';
    this.http.post(api,{"userid":this.id},httpOptions).subscribe((response:any)=>{
            this.list=response;
           // console.log(this.list)
           //console.log(this.id)
           //alert('成功')      
    })
   }
  }
  fuikui(){
    this.router.navigate(['/Feedback/'])
  }
  bin(){
    this.router.navigate(['/homepage/'])
    
  }

}
