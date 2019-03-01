import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  apiUrl1 = "http://localhost:7080/thread"
  apiUrl2 = "http://localhost:7080/thread/list"
  apiUrl3 = "http://localhost:7080/thread/post"
  apiUrl4 = "http://localhost:7080/thread/reply"
  apiUrl5 = "http://localhost:7080/thread/search"
  apiUrl6 = "http://localhost:7080/thread/collect"
  apiUrl7 = "http://localhost:7080/thread/cancel"

  constructor(
    private http: HttpClient
  ) { }

  imitateLogin(): Observable<any> {
    return this.http.get<any>(this.apiUrl1);
  }

  getThreadList(): Observable<any> {
    return this.http.get<any>(this.apiUrl2)
  }

  getPostList(userID, topicID: string): Observable<any> {
    return this.http.get<any>(this.apiUrl3, { params: { userID: userID, topicID: topicID } })
  }

  submitReply(userID: string, topicID: string, text: string, floor: string, ): Observable<any> {
    return this.http.get<any>(this.apiUrl4, { params: { userID: userID, topicID: topicID, text: text, floor: floor, } })
  }

  search(condition: string): Observable<any> {
    return this.http.get<any>(this.apiUrl5, { params: { condition: condition } })
  }

  collect(userID: string, topicID: string): Observable<any> {
    return this.http.get<any>(this.apiUrl6, { params: { userID: userID, topicID: topicID } })
  }

  cancel(userID: string, topicID: string): Observable<any> {
    return this.http.get<any>(this.apiUrl7, { params: { userID: userID, topicID: topicID } })
  }

}
