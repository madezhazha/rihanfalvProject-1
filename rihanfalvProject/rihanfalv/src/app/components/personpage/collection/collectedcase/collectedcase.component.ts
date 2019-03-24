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
  public UserID:any=1;
  public Literature:any[]=[];
  public Cases:any[]=[];
  public Flag:number=1;
  


  constructor(private router: Router, public http:HttpClient) { }
  ngOnInit() {
    let userid=localStorage.getItem("id")

    this.UserID=userid


    this.CollectionMsg.length=0
    const httpOptions = {headers: new HttpHeaders({ 'Content-Type':'application/json'})};

       var api ='http://blackcardriver.cn:7080/collectioncase';
       this.http.post(api,{"userid":this.UserID},httpOptions).subscribe((response:any)=>{

        this.CollectionMsg=response;
         

         
         console.log(response);
       })
     
  }
  //specificContent(CollectionType,CollectionContentID,CollectionTitle){

    specificContent(CollectionTitle){
    /*if(CollectionType=="japancase"){
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };       
    //this.router.navigate(['/case'],Contentid) ;

    this.router.navigate(["/display-data"],{queryParams:{"title":CollectionTitle}})
    }
    if(CollectionType=="koreacase"){
      let Contentid: NavigationExtras = {             
      queryParams: { ContentID:CollectionContentID },                
    };       
    //this.router.navigate(['/case'],Contentid) ;*/

    this.router.navigate(["/display-data"],{queryParams:{"title":CollectionTitle}})
    }
    
  }



