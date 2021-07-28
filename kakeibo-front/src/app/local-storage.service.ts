import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LocalStorageService {

  constructor() { }

  Get(key:string){
    return localStorage.getItem(key)
  }

  Set(key:string, value:string){
    localStorage.setItem(key, value)
  }
  
  Remove(key:string){
    localStorage.removeItem(key)
  }
}
