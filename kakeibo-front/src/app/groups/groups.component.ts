import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Group } from '../models/group.model';
import { GroupServiceService } from './services/group-service.service';

@Component({
  selector: 'app-groups',
  templateUrl: './groups.component.html',
  styleUrls: ['./groups.component.css']
})


export class GroupsComponent implements OnInit {
  @Input() columns = ['id', 'name'];

  dataSource: Group[];

  constructor
  (
    private service: GroupServiceService,
    private router: Router
  ) 
  {

  }

  ngOnInit(): void {

    this.service.getGroups()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    this.router.navigate(['/groups/insert', row.id]);
    
  }

  createNew(){
    this.router.navigate(['/groups/insert']);
  }
}
