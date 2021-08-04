import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Income } from '../models/income.model';

@Injectable({
  providedIn: 'root'
})
export class IncomeService {

  constructor(private http: HttpClient) { }

  getIncomes(): Observable<Income[]>{
   return this.http.get<Income[]>('api/income/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  getIncomeById(id: number): Observable<Income>{
    return this.http.get<Income>('api/income/'+ id)
     .pipe( 
       catchError(err =>{
         return EMPTY
       })
     )
   }

  postIncome(req: Income): Observable<Income> {
    return this.http.post<Income>('api/income/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  putIncome(req: Income): Observable<Income> {
    return this.http.put<Income>('api/income/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  deleteIncome(id: number): Observable<Income> {
    return this.http.delete<Income>('api/income/'+ id)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
