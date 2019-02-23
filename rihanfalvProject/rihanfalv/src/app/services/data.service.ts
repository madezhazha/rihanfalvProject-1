import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  apiUrl1 = "http://localhost:8888/thread"
  apiUrl2 = "http://localhost:8888/thread/list"
  apiUrl3 = "http://localhost:8888/thread/post"
  apiUrl4 = "http://localhost:8888/thread/reply"
  apiUrl5 = "http://localhost:8888/thread/search"

  constructor(
    private http: HttpClient
  ) { }

  imitateLogin(): Observable<any> {
    return this.http.get<any>(this.apiUrl1);
  }

  getThreadList(): Observable<any> {
    return this.http.get<any>(this.apiUrl2)
  }

  getPostList(topicID: string): Observable<any> {
    return this.http.get<any>(this.apiUrl3, { params: { topicID: topicID } })
  }

  submitReply(userID: string, topicID: string, text: string, floor: string, ): Observable<any> {
    return this.http.get<any>(this.apiUrl4, { params: { userID: userID, topicID: topicID, text: text, floor: floor, } })
  }

  search(condition: string) {
    return this.http.get<any>(this.apiUrl5, { params: { condition: condition } })
  }

}
