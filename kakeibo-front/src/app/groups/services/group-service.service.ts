import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Group } from 'src/app/models/group.model';

@Injectable({
  providedIn: 'root'
})
export class GroupServiceService {

  constructor(private http: HttpClient) { }

  getGroups(): any{

   return this.http.get<Group[]>('api/group/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
