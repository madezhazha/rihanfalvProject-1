import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class GetdataService {

  constructor(private http: HttpClient) { }
  get(id) {
    return this.http.post('http://localhost:7080/get',id);
  }
  change(user) {
    return this.http.post('http://localhost:7080/post', user);
  }
}
