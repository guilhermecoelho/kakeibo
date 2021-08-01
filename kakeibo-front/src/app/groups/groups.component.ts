import { Component, Input, OnInit } from '@angular/core';
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

  constructor(private service: GroupServiceService) { }

  ngOnInit(): void {

    this.service.getGroups()
    .subscribe(data => this.dataSource = data)
  }

  clickedRows(row: any){
    console.log(row)
  }
}
