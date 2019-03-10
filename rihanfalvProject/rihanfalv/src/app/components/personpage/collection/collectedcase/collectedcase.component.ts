import { Component, OnInit } from '@angular/core';
import { HttpClient ,HttpHeaders} from '@angular/common/http';
import { Router,NavigationExtras} from '@angular/router';

@Component({
  selector: 'app-collectedcase',
  templateUrl: './collectedcase.component.html',
  styleUrls: ['./collectedcase.component.css']
})
export class CollectedcaseComponent implements OnInit {

  public arr:any[]=[];
  public CollectionMsg:any[]=[];
  public UserID:number=1;
  public Literature:any[]=[];
  public Cases:any[]=[];
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
           if(this.arr[i].CollectionType=="japancase"||this.arr[i].CollectionType=="koreacase"){
             this.CollectionMsg[k]=this.arr[i];
             k++
           }
           
                     
         }
     
         console.log(response);
       })
     
  }
  specificContent(CollectionType,CollectionContentID){

    if(CollectionType=="japancase"){
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };       
    this.router.navigate(['/case'],Contentid) ;
    }
    if(CollectionType=="koreacase"){
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };       
    this.router.navigate(['/case'],Contentid) ;
    }
    
  }


}
