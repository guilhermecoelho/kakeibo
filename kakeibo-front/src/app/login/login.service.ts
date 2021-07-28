import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) { }

  postLogin(req: any): Observable<any> {
    return this.http.post<any>('api/login/', req,{
      headers:new HttpHeaders({
        'ignore': 'true'
      })
    })
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
