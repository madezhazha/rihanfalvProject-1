import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class GetdataService {

  constructor(private http: HttpClient) { }
  get() {
    return this.http.get('http://localhost:8080/get');
  }
  change(user) {
    return this.http.post('http://localhost:8080/put', user);
  }
}
