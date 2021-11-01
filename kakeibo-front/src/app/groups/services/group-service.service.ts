import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Group } from 'src/app/models/group.model';

@Injectable({
  providedIn: 'root'
})
export class GroupServiceService {

  constructor(private http: HttpClient) { }

  getGroups(): Observable<Group[]>{
   return this.http.get<Group[]>('api/group/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  getGroupById(id: number): Observable<Group>{
    return this.http.get<Group>('api/group/'+ id)
     .pipe( 
       catchError(err =>{
         return EMPTY
       })
     )
   }

  postGroups(req: Group): Observable<Group> {
    return this.http.post<Group>('api/group/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  putGroup(req: Group): Observable<Group> {
    return this.http.put<Group>('api/group/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  deleteGroup(id: number): Observable<Group> {
    return this.http.delete<Group>('api/group/'+ id)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
