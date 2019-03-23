import { Injectable } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import { Observable} from 'rxjs';
import {HotnewsBox, ArticalBox} from './struct';

@Injectable({
  providedIn: 'root'
})

export class GolangService {
  constructor(
    private http:HttpClient
  ){}
  //vps
  // private url ='http://192.168.127.128:7080';
  // virtual machin
    private url = 'http://blackcardriver.cn:7080';
  //private url ='http://localhost:7080' ;  
  //请求首页的文章数据，从index 到 index+9 的文章数据，自动添加到文章列表末尾
  GetArtical(index:number):Observable<ArticalBox[]>{  
      var post = {index:index};
      return this.http.post<ArticalBox[]>(
        this.url + "/homepage/articals",post
      ).pipe();
  }
  
  //请求首页自动播放框的数据，长度为5的结构数组
  GetHeadNews():Observable<HotnewsBox[]>{
      var getnews = this.url +"/homepage/hotnews";
      return this.http.get<HotnewsBox[]>(getnews).pipe();
  }

}
// mux.HandleFunc("/homepage/articals", route.GetHomePageArtical)
// mux.HandleFunc("/homepage/hotnews", route.GetHomePageHotnews)