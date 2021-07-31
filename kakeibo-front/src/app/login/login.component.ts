import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import {  Router } from '@angular/router';
import { LocalStorageService } from '../local-storage.service';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup

  constructor(
    private formBuilder: FormBuilder, 
    private loginService: LoginService, 
    private localStorageService: LocalStorageService,
    private router: Router) 
    {
    this.loginForm = formBuilder.group({
      user: [''],
      password: ['']
    })
   }

  ngOnInit(): void {
  }

  onSubmit(): void{
    this.loginService.postLogin(this.loginForm.value)
    .subscribe((resp) =>
    {
      let token = "Bearer "+ resp.token
      this.localStorageService.Set("token", token)
      this.router.navigate(['']);
    })
  }

}
