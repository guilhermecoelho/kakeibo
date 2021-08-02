import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Expense } from '../models/expense.model';

@Injectable({
  providedIn: 'root'
})
export class ExpensesService {

  constructor(private http: HttpClient) { }

  getExpenses(): Observable<Expense[]>{
   return this.http.get<Expense[]>('api/expense/')
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  getExpenseById(id: number): Observable<Expense>{
    return this.http.get<Expense>('api/expense/'+ id)
     .pipe( 
       catchError(err =>{
         return EMPTY
       })
     )
   }

  postExpense(req: Expense): Observable<Expense> {
    return this.http.post<Expense>('api/expense/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  putExpense(req: Expense): Observable<Expense> {
    return this.http.put<Expense>('api/expense/', req)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }

  deleteExpense(id: number): Observable<Expense> {
    return this.http.delete<Expense>('api/expense/'+ id)
    .pipe( 
      catchError(err =>{
        return EMPTY
      })
    )
  }
}
