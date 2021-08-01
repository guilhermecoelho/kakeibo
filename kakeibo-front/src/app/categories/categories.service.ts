import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Category } from '../models/category.model';

@Injectable({
  providedIn: 'root'
})
export class CategoriesService {

  constructor(private http: HttpClient) { }

  getCategories(): Observable<Category[]>{
   return this.http.get<Category[]>('api/category/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  getCategoryById(id: number): Observable<Category>{
    return this.http.get<Category>('api/category/'+ id)
     .pipe( 
       catchError(err =>{
         return EMPTY
       })
     )
   }

  postCategory(req: Category): Observable<Category> {
    return this.http.post<Category>('api/category/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  putCategory(req: Category): Observable<Category> {
    return this.http.put<Category>('api/category/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  deleteCategory(id: number): Observable<Category> {
    return this.http.delete<Category>('api/category/'+ id)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
