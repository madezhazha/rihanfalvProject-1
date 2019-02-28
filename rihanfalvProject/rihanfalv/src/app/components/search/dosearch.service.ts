import { Injectable } from '@angular/core';
import {HttpClient,HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class DosearchService {
  constructor(public http:HttpClient) { }
  public KeyWord:any
  public getkey:any={}
  public list:any
public Classify:any="全部"

  searchtogo(){
    const httpOptions={ headers:new HttpHeaders({'Content-Type':'application/json'}) };
    let api='http://localhost:9090/legal'; 
    this.getkey.KeyWord=this.KeyWord
    this.getkey.Classify=this.Classify
    this.http.post(api,this.getkey,httpOptions).subscribe(response=>{ 
      this.list=response
    });
  }

}
