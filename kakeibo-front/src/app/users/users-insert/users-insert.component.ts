import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { UsersService } from '../users.service';

@Component({
  selector: 'app-users-insert',
  templateUrl: './users-insert.component.html',
  styleUrls: ['./users-insert.component.css']
})
export class UsersInsertComponent implements OnInit {

  userForm: FormGroup
  user: User

  constructor(
    private formBuilder: FormBuilder, 
    private service: UsersService, 
    private router: Router,
    private route: ActivatedRoute,
    ) {}
   
  ngOnInit(): void {

    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    this.userForm = this.formBuilder.group({
      user: [``],
      password: [``],
    });

    if(id !== null && !isNaN(id)){
      this.service.getUserById(id)
      .subscribe(data => {
        this.user = data
        this.userForm.controls[`user`].setValue(this.user.user)
      });
    }
  }

  onSubmit(): void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

    if(id !== null && !isNaN(id)) {
      this.user.id = id
      this.user.user = this.userForm.value.name

      this.service.putUser(this.user)
      .subscribe((resp) => {
        this.router.navigate(['users']);
      })
      return
    }

    this.service.postUser(this.userForm.value)
    .subscribe((resp) => {
      this.router.navigate(['users']);
    })
  }
  
  onDelete():void{
    const id = parseInt(this.route.snapshot.paramMap.get('id'), 10);

      if(id !== null && !isNaN(id)) {

        this.service.deleteUser(id)
        .subscribe((resp) => {
          this.router.navigate(['users']);
        })
    }
  }
}
