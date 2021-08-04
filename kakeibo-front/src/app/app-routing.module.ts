import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CategoriesInsertComponent } from './categories/categories-insert/categories-insert.component';
import { CategoriesComponent } from './categories/categories.component';
import { ExpensesInsertComponent } from './expenses/expenses-insert/expenses-insert.component';
import { ExpensesComponent } from './expenses/expenses.component';
import { GroupsInsertComponent } from './groups/groups-insert/groups-insert.component';
import { GroupsComponent } from './groups/groups.component';
import { IncomeInsertComponent } from './incomes/income-insert/income-insert.component';
import { IncomesComponent } from './incomes/incomes.component';
import { LoginComponent } from './login/login.component';
import { UsersInsertComponent } from './users/users-insert/users-insert.component';
import { UsersComponent } from './users/users.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  
  { path: 'groups', component: GroupsComponent },
  { path: 'groups/insert', component: GroupsInsertComponent },
  { path: 'groups/insert/:id', component: GroupsInsertComponent },

  { path: 'categories', component: CategoriesComponent },
  { path: 'categories/insert', component: CategoriesInsertComponent },
  { path: 'categories/insert/:id', component: CategoriesInsertComponent },

  { path: 'expenses', component: ExpensesComponent },
  { path: 'expenses/insert', component: ExpensesInsertComponent },
  { path: 'expenses/insert/:id', component: ExpensesInsertComponent },

  { path: 'incomes', component: IncomesComponent },
  { path: 'incomes/insert', component: IncomeInsertComponent },
  { path: 'incomes/insert/:id', component: IncomeInsertComponent },

  { path: 'users', component: UsersComponent },
  { path: 'users/insert', component: UsersInsertComponent },
  { path: 'users/insert/:id', component: UsersInsertComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { relativeLinkResolution: 'legacy' })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
