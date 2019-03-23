import { HttpClient, HttpHeaders } from '@angular/common/http';
import { APIResponse, APIResponse2, APIResponse3, Nowcountry, Nowpage, Allpage } from './apiresponse';
import { Observable, of } from 'rxjs';
import { Injectable } from '@angular/core';

const httpOptions = {
       headers: new HttpHeaders({'Content-Type': 'application/json'})
};



@Injectable()
export class ApiSerivice {
    constructor(private http: HttpClient) {
    }

    public legalpage(nowpage: number, kind: string): Observable<Nowpage> {
<<<<<<< HEAD
        return this.http.post<Nowpage>('http://localhost:7080/page',
=======
        return this.http.post<Nowpage>('http://blackcardriver.cn:7080/page',
>>>>>>> 4533ee3495af12391c955345f7eebc0240676ed0
        {nowpage, kind}, httpOptions);
    }

    public legalpage2(): Observable<Allpage> {
<<<<<<< HEAD
        return this.http.get<Allpage>('http://localhost:7080/pages');
=======
        return this.http.get<Allpage>('http://blackcardriver.cn:7080/pages');
>>>>>>> 4533ee3495af12391c955345f7eebc0240676ed0
    }

    public country(Country: string): Observable<Nowcountry> {
        return this.http.post<Nowcountry>('http://blackcardriver.cn:7080/country',
        {Country}, httpOptions);
    }

    public legaltype(): Observable<APIResponse[]> {
        return this.http.get<APIResponse[]>('http://blackcardriver.cn:7080/type');
    }

    public legaltype2(legaltype: string): Observable<APIResponse> {
        return this.http.post<APIResponse>('http://blackcardriver.cn:7080/title',
        {legaltype}, httpOptions);
    }

    public legallabel(legallabel: string): Observable<APIResponse> {
        return this.http.post<APIResponse>('http://blackcardriver.cn:7080/label',
        {legallabel}, httpOptions);
    }

    public legallabel2(): Observable<APIResponse[]> {
        return this.http.get<APIResponse[]>('http://blackcardriver.cn:7080/labels');
    }

    public legaltitle(): Observable<APIResponse2[]> {
        return this.http.get<APIResponse2[]>('http://blackcardriver.cn:7080/titles');
    }

    public legaltitle2(legaltitle: string): Observable<APIResponse2> {
        return this.http.post<APIResponse2>('http://blackcardriver.cn:7080/content',
        {legaltitle}, httpOptions);
    }

    public legalcontent(): Observable<APIResponse3> {
        return this.http.get<APIResponse3>('http://blackcardriver.cn:7080/contents');
    }
}
