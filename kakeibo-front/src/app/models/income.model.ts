import { Category } from "./category.model";

export interface Income{
    id: number,
    name: string,
    date: Date,
    categoryId: number,
    amount: number,
    note: string,
    category: Category
  }