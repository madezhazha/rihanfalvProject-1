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

       var api ='http://localhost:7080/collectionthesis';
       this.http.post(api,{"userid":this.UserID},httpOptions).subscribe((response:any)=>{

        this.CollectionMsg=response;
         

         

         
       })
     
  }
  specificContent(CollectionType,CollectionContentID){



    localStorage.setItem("route","/collection") 

    if(CollectionType=="japanthesis"){

    
      var JapanOrKorea:string="日";    
  
      localStorage.setItem("JapanOrKorea",JapanOrKorea)  

     /* let Contentid: NavigationExtras = {             
      queryParams: { ArticleID: CollectionContentID },                
    }; 
        
    this.router.navigate(['/paperweb'],Contentid) ;*/
    }

    if(CollectionType=="koreathesis"){
      
      var JapanOrKorea:string="韩";      
    
      localStorage.setItem("JapanOrKorea",JapanOrKorea)
     /* let Contentid: NavigationExtras = {             
      queryParams: { ArticleID:CollectionContentID },                
    };   
       
    this.router.navigate(['/paperweb'],Contentid) ;*/
    }
    
  }


}
