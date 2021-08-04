import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../models/user.model';
import { UsersService } from './users.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  @Input() columns = ['id', 'user'];

  dataSource: User[];

  constructor(
    private service: UsersService,
    private router: Router
  ) { }

  ngOnInit(): void {

    this.service.getUsers()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    this.router.navigate(['/users/insert', row.id]);
    
  }

  createNew(){
    this.router.navigate(['/users/insert']);
  }
}
