import { HttpClient } from '@angular/common/http';
import { Injectable, Input } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { User } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
 
  constructor(private http: HttpClient) { }

  getUsers(): Observable<User[]>{
   return this.http.get<User[]>('api/user/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  getUserById(id: number): Observable<User>{
    return this.http.get<User>('api/user/'+ id)
     .pipe( 
       catchError(err =>{
         return EMPTY
       })
     )
   }

  postUser(req: User): Observable<User> {
    return this.http.post<User>('api/user/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  putUser(req: User): Observable<User> {
    return this.http.put<User>('api/user/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  deleteUser(id: number): Observable<User> {
    return this.http.delete<User>('api/user/'+ id)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
