import { Component, OnInit } from '@angular/core';
import { HttpClient ,HttpHeaders} from '@angular/common/http';
import { Router,NavigationExtras} from '@angular/router';


@Component({
  selector: 'app-thesis',
  templateUrl: './thesis.component.html',
  styleUrls: ['./thesis.component.css']
})
export class ThesisComponent implements OnInit {

  public arr:any[]=[];
  public CollectionMsg:any[]=[];
  public UserID:number=1;
  public Thesis:any[]=[];
  public Flag:number=1;
  


  constructor(private router: Router, public http:HttpClient) { }
  ngOnInit() {
    this.CollectionMsg.length=0
    const httpOptions = {headers: new HttpHeaders({ 'Content-Type':'application/json'})};

       var api ='http://localhost:7080/collection';
       this.http.post(api,{"userid":this.UserID},httpOptions).subscribe((response:any)=>{

         this.arr=response;
         

         var i:number;
         var k:number=0;
        
         for( i=0;i<100;i++){
           if(this.arr[i].CollectionTime.length<3){
             break;
           } 
           if(this.arr[i].CollectionType=="japanthesis"||this.arr[i].CollectionType=="koreathesis"){
             this.CollectionMsg[k]=this.arr[i];
             k++
           }
           
                     
         }

         
         console.log(response);
       })
     
  }
  specificContent(CollectionType,CollectionContentID){

    if(CollectionType=="japanthesis"){

     

      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    }; 
    var JapanOrKorea:string="日";    
    localStorage.setItem("JapanOrKorea",JapanOrKorea)      
    this.router.navigate(['/paperweb'],Contentid) ;
    }
    if(CollectionType=="koreathesis"){
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };   
    var JapanOrKorea:string="韩";      
    localStorage.setItem("JapanOrKorea",JapanOrKorea)
    
    this.router.navigate(['/paperweb'],Contentid) ;
    }
    
  }


}
