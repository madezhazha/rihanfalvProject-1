import { Component, OnInit } from '@angular/core';
import { HttpClient ,HttpHeaders} from '@angular/common/http';
import { Router,NavigationExtras} from '@angular/router';


@Component({
  selector: 'app-collectedtopic',
  templateUrl: './collectedtopic.component.html',
  styleUrls: ['./collectedtopic.component.css']
})
export class CollectedtopicComponent implements OnInit {

  public arr:any[]=[];
  public CollectionMsg:any[]=[];
  public UserID:number=1;
  public Literature:any[]=[];
  public Topics:any[]=[];
  public Flag:number=1;
  


  constructor(private router: Router, public http:HttpClient) { }
  ngOnInit() {
    this.CollectionMsg.length=0
    const httpOptions = {headers: new HttpHeaders({ 'Content-Type':'application/json'})};

       var api ='http://localhost:7080/collectiontopic';
       this.http.post(api,{"userid":this.UserID},httpOptions).subscribe((response:any)=>{

        this.CollectionMsg=response;
         

         
         console.log(response);
       })
     
  }
  specificContent(CollectionType,CollectionContentID){

    
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };       
    this.router.navigate(['/post'],Contentid) ;
    
    
    
  }


}
