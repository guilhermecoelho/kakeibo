import { Expense } from "./expense.model";
import { Income } from "./income.model";

export interface Report{
    dateStart: Date,
    dateFinish: Date,
    totalIncome: Float32Array,
    totalExpense: Float32Array,
    balance: Float32Array,
    Expenses: Expense[],
    Incomes: Income[]
  }