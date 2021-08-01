import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CategoriesInsertComponent } from './categories/categories-insert/categories-insert.component';
import { CategoriesComponent } from './categories/categories.component';
import { GroupsInsertComponent } from './groups/groups-insert/groups-insert.component';
import { GroupsComponent } from './groups/groups.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  
  { path: 'groups', component: GroupsComponent },
  { path: 'groups/insert', component: GroupsInsertComponent },
  { path: 'groups/insert/:id', component: GroupsInsertComponent },

  { path: 'categories', component: CategoriesComponent },
  { path: 'categories/insert', component: CategoriesInsertComponent },
  { path: 'categories/insert/:id', component: CategoriesInsertComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
