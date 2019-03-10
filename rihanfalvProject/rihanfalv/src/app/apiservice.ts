import { HttpClient, HttpHeaders } from '@angular/common/http';
import { APIResponse, APIResponse2, APIResponse3 } from './apiresponse';
import { Observable, of } from 'rxjs';
import { Injectable } from '@angular/core';

const httpOptions = {
       headers: new HttpHeaders({'Content-Type': 'application/json'})
};



@Injectable()
export class ApiSerivice {
    constructor(private http: HttpClient) {
    }

    public legaltype(): Observable<APIResponse[]> {
        return this.http.get<APIResponse[]>('http://localhost:7080/type');
    }

    public legaltype2(legaltype: string): Observable<APIResponse> {
        return this.http.post<APIResponse>('http://localhost:7080/title',
        {legaltype: legaltype}, httpOptions);
    }

    public legallabel(legallabel: string): Observable<APIResponse> {
        return this.http.post<APIResponse>('http://localhost:7080/label',
        {legallabel: legallabel}, httpOptions);
    }

    public legallabel2(): Observable<APIResponse[]> {
        return this.http.get<APIResponse[]>('http://localhost:7080/labels');
    }

    public legaltitle(): Observable<APIResponse2[]> {
        return this.http.get<APIResponse2[]>('http://localhost:7080/titles');
    }

    public legaltitle2(legaltitle: string): Observable<APIResponse2> {
        return this.http.post<APIResponse2>('http://localhost:7080/content',
        {legaltitle: legaltitle}, httpOptions);
    }

    public legalcontent(): Observable<APIResponse3> {
        return this.http.get<APIResponse3>('http://localhost:7080/contents');
    }
}
