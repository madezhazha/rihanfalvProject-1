import { Injectable } from '@angular/core';
import{InputData} from '../input';
import { HttpClient } from '@angular/common/http';
import {OutputData} from '../output'


@Injectable({
  providedIn: 'root'
})
export class LandService {

  constructor(private http: HttpClient) { }
  getInput(out:OutputData){
    return this.http.post<InputData>("http://localhost:8090/",JSON.stringify(out))
  }
}
